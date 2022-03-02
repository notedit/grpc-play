package main

import (
	"context"
	"io"
	"log"
	"net"
	"sync"

	"github.com/notedit/grpc-play/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	protos.UnimplementedSignalingServer
	Broadcast     chan *protos.StreamResponse
	Clients       map[string]string
	ClientStreams map[string]chan *protos.StreamResponse

	userMux, streamMux sync.RWMutex
}

func (s *server) broadcast(ctx context.Context) {
	for res := range s.Broadcast {
		s.streamMux.RLock()
		for _, stream := range s.ClientStreams {
			select {
			case stream <- res:

			default:
				log.Println("stream is full ", stream)
			}
		}
		s.streamMux.RUnlock()
	}
}

func (s *server) Join(ctx context.Context, in *protos.JoinRequest) (*protos.JoinResponse, error) {
	s.userMux.Lock()
	defer s.userMux.Unlock()

	s.userMux.Lock()
	s.Clients[in.User] = in.User
	s.userMux.Unlock()

	s.Broadcast <- &protos.StreamResponse{
		Event: &protos.StreamResponse_Login{
			Login: &protos.Login{
				User: in.User,
			},
		},
	}

	return &protos.JoinResponse{
		Code: 0,
	}, nil
}

func (s *server) Leave(ctx context.Context, in *protos.LeaveRequest) (*protos.LeaveResponse, error) {
	s.userMux.Lock()
	defer s.userMux.Unlock()

	delete(s.Clients, in.User)

	s.Broadcast <- &protos.StreamResponse{
		Event: &protos.StreamResponse_Leave{
			Leave: &protos.Leave{
				User: in.User,
			},
		},
	}

	return &protos.LeaveResponse{}, nil
}

func (s *server) Stream(srv protos.Signaling_StreamServer) error {
	user := srv.Context().Value("user").(string)
	log.Println("stream user ", user)

	stream := make(chan *protos.StreamResponse, 100)
	s.streamMux.Lock()
	s.ClientStreams[user] = stream
	s.streamMux.Unlock()

	defer func() {
		s.streamMux.Lock()
		if local, ok := s.ClientStreams[user]; ok {
			delete(s.ClientStreams, user)
			close(local)
		}

		delete(s.ClientStreams, user)
		s.streamMux.Unlock()
	}()

	go s.sendbroadcaster(srv, stream)

	for {
		req, err := srv.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		s.Broadcast <- &protos.StreamResponse{
			Event: &protos.StreamResponse_Message{
				Message: &protos.Message{
					User:    user,
					Message: req.Room,
				},
			},
		}
	}

	<-srv.Context().Done()
	return srv.Context().Err()
}

func (s *server) sendbroadcaster(srv protos.Signaling_StreamServer, stream chan *protos.StreamResponse) {

	for {
		select {
		case <-srv.Context().Done():
			return
		case res := <-stream:
			if s, ok := status.FromError(srv.Send(res)); ok {
				switch s.Code() {
				case codes.OK:
					// noop
				case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
					log.Println("client terminated connection")
					return
				default:
					return
				}
			}
		}
	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := &server{
		Broadcast:     make(chan *protos.StreamResponse, 10000),
		Clients:       make(map[string]string),
		ClientStreams: make(map[string]chan *protos.StreamResponse),
	}
	grpcServer := grpc.NewServer()
	protos.RegisterSignalingServer(grpcServer, s)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	go s.broadcast(ctx)

	go func() {
		if err := grpcServer.Serve(l); err != nil {
			cancel()
		}
	}()

	<-ctx.Done()

	grpcServer.GracefulStop()
}
