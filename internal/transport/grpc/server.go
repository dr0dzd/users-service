package grpc

import (
  "google.golang.org/grpc"
  "net"
  userpb "github.com/dr0dzd/project-protos/proto_gen/user"
  "github.com/dr0dzd/users-service/internal/user"
)

func RunGRPC(svc *user.Service) error {
  listener, err := net.Listen("tcp", ":50051")
  if err != nil {
    return err
  }

  grpcServer := grpc.NewServer()
  userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

  return grpcServer.Serve(listener)
}
