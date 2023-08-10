package client

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	pb "github.com/wzhanjun/log-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	logsChan = make(chan *pb.LogRequest, 100000)
	workNum  = 5
	connLock sync.RWMutex
)

type GrpcHandler struct {
	handler.NopFlushClose
	slog.LevelWithFormatter

	client pb.LogClient
}

func NewGprcHandler() *GrpcHandler {
	s := &GrpcHandler{}
	// s.client = s.connect()
	s.Level = slog.InfoLevel

	go s.push()

	return s
}

func (s *GrpcHandler) connect() pb.LogClient {
	connLock.Lock()
	defer connLock.Unlock()

	if s.client == nil {
		conn, err := NewGrpcConn(Cfg.LogServiceAddress)
		if err != nil {
			log.Printf("init grpc client failed :%+v \n", err)
			return nil
		}
		s.client = pb.NewLogClient(conn)
	}
	return s.client
}

func (s *GrpcHandler) Handle(r *slog.Record) error {
	logsChan <- &pb.LogRequest{
		AppId:         Cfg.AppId,
		Label:         StrLabel(r),
		Level:         r.Level.String(),
		Content:       r.Message,
		Caller:        StrCaller(r),
		Datatime:      r.Time.Format("2006/01/02T15:04:05.000"),
		EsIndexPrefix: Cfg.LogServiceEsIndex,
	}
	return nil
}

func NewGrpcConn(address string) (*grpc.ClientConn, error) {
	ctx, cf := context.WithTimeout(context.Background(), time.Second*3)
	defer cf()
	if conn, err := grpc.DialContext(
		ctx,
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	); err != nil {
		log.Println(err)
		return nil, err
	} else {
		return conn, nil
	}
}

func (s *GrpcHandler) push() {

	for i := 0; i < workNum; i++ {
		go func() {
			for m := range logsChan {
				log.Println(m)
				if Cfg.AppDeBug {
					continue
				}
				if client := s.connect(); client != nil {
					ctx, cf := context.WithTimeout(context.Background(), time.Second*5)
					defer cf()
					switch m.Level {
					case slog.TraceLevel.Name(), slog.DebugLevel.Name():
						client.Debug(ctx, m)
					case slog.InfoLevel.Name():
						client.Info(ctx, m)
					case slog.WarnLevel.Name(), slog.NoticeLevel.Name():
						client.Warn(ctx, m)
					case slog.ErrorLevel.Name():
						client.Error(ctx, m)
					case slog.FatalLevel.Name(), slog.PanicLevel.Name():
						client.Fatal(ctx, m)
					}
				}
			}
		}()
	}
}
