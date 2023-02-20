package grpclibdemo

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gopractice/lib/proto/common"
	"net"
)

type Config struct {
	Host string `json:"host" toml:"host"`
	Port string `json:"port" toml:"port"`
}

func (c *Config) GetAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

type Server struct {
	conf   *Config
	server *grpc.Server
	router Router
}

func (s *Server) Start() (err error) {
	addr := net.JoinHostPort(s.conf.Host, s.conf.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen: %s \n", err)
		return
	}

	s.server = grpc.NewServer()
	s.router.RegGrpc(s.server)

	go func() {
		s.server.Serve(lis)
	}()

	return
}

func NewServer(conf *Config, router Router) *Server {
	return &Server{conf: conf, router: router}
}

type Router interface {
	RegGrpc(server *grpc.Server)
}

type GRouter struct {
}

func (r *GRouter) RegGrpc(server *grpc.Server) {
	common.RegisterCommonServiceServer(server, commonServer)
	return
}

var commonServer = &commonService{}

type commonService struct {
	*common.UnimplementedCommonServiceServer
}

func (c *commonService) CommonTest(ctx context.Context, empty *common.Empty) (rsp *common.Response, err error) {
	rsp = &common.Response{
		Code:    200,
		Message: "success",
	}

	return
}
