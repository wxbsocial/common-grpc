package api

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	ip   string
	port string
}

func NewServer(
	ip string, port string,
) *Server {

	return &Server{
		port: port,
		ip:   ip,
	}
}

func (s *Server) Start() error {

	address := fmt.Sprintf("%s:%s", s.ip, s.port)

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}

	server := s.NewServer()
	s.RegisterServer(server)

	log.Printf("start to listen: %v", address)
	if err := server.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return err
	}

	return nil
}

func (s *Server) NewServer() *grpc.Server {
	return grpc.NewServer()
}

func (s *Server) RegisterServer(server *grpc.Server) {

}
