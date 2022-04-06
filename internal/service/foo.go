package service

import (
	"context"

	pb "foo/api/foo"
)

type FooService struct {
	pb.UnimplementedFooServer
}

func NewFooService() *FooService {
	return &FooService{}
}

func (s *FooService) CreateFoo(ctx context.Context, req *pb.CreateFooRequest) (*pb.CreateFooReply, error) {
	return &pb.CreateFooReply{}, nil
}
func (s *FooService) UpdateFoo(ctx context.Context, req *pb.UpdateFooRequest) (*pb.UpdateFooReply, error) {
	return &pb.UpdateFooReply{}, nil
}
func (s *FooService) DeleteFoo(ctx context.Context, req *pb.DeleteFooRequest) (*pb.DeleteFooReply, error) {
	return &pb.DeleteFooReply{}, nil
}
func (s *FooService) GetFoo(ctx context.Context, req *pb.GetFooRequest) (*pb.GetFooReply, error) {
	return &pb.GetFooReply{}, nil
}
func (s *FooService) ListFoo(ctx context.Context, req *pb.ListFooRequest) (*pb.ListFooReply, error) {
	return &pb.ListFooReply{}, nil
}
