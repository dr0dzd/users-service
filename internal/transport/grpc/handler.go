package grpc

import (
	"context"

	userpb "github.com/dr0dzd/project-protos/proto_gen/user"
	"github.com/dr0dzd/users-service/internal/user"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
  svc *user.Service
  userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
  return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
  user := user.NewUser(req.Email)
  createdUser, err := h.svc.CreateUser(user)
  if err != nil {
    return nil, err
  }

  return &userpb.CreateUserResponse{
    User: &userpb.User{
      Id:    createdUser.ID.String(),
      Email: createdUser.Email,
      CreatedAt: timestamppb.Now(),
      UpdatedAt: timestamppb.Now(),
    },
  }, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
  id, ok := uuid.Parse(req.Id)
  if ok != nil{
    return nil, status.Error(codes.InvalidArgument, "ID format should be UUID")
  }
  user, err := h.svc.GetUser(id)
  if err != nil{
    return nil, err
  }
  return &userpb.GetUserResponse{
    User: &userpb.User{
      Id: user.ID.String(),
      Email: user.Email},
  }, nil
}

func (h *Handler) GetAllUsers(ctx context.Context, req *userpb.GetAllUsersRequest) (*userpb.GetAllUsersResponse, error){
  users, err := h.svc.GetAllUsers()
  if err != nil{
    return nil, err
  }
  return &userpb.GetAllUsersResponse{
    Users: lo.Map(users, func(item user.User, index int) (*userpb.User) {
      return &userpb.User{
        Id: item.ID.String(),
        Email: item.Email,
      }
    }),
  }, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error){
  id, ok := uuid.Parse(req.Id)
  if ok != nil{
    return nil, status.Error(codes.InvalidArgument, "ID format should be UUID")
  }
  updatedUser, err := h.svc.UpdateUser(user.User{ID: id, Email: req.Email})
  if err != nil{
    return &userpb.UpdateUserResponse{}, err
  }
  return &userpb.UpdateUserResponse{
    User: &userpb.User{Id: updatedUser.ID.String(), Email: updatedUser.Email},
  }, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*empty.Empty, error) {
  id, ok := uuid.Parse(req.Id)
  if ok != nil{
    return nil, status.Error(codes.InvalidArgument, "ID format should be UUID")
  }
  err := h.svc.DeleteUser(id)
  if err != nil{
    return nil, err
  }
  return &emptypb.Empty{}, nil
}

