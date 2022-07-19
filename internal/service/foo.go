package service

import (
	"context"
	pb "foo/api/foo/v1"
	"foo/internal/biz"

	"github.com/golang/protobuf/ptypes/empty"
)

type FooService struct {
	pb.UnimplementedFooServer
	fuc *biz.FooUseCase
}

func NewFooService(fuc *biz.FooUseCase) *FooService {
	return &FooService{fuc: fuc}
}

func (s *FooService) CreateFoo(ctx context.Context, req *pb.CreateFooRequest) (*empty.Empty, error) {
	_, err := s.fuc.CreateFoo(ctx, &biz.Foo{Msg: req.Name})
	if err != nil {
		return nil, pb.ErrorUnknown("create foo failed")
	}
	return nil, nil
}

func naming(s string) string {
	if s == "name" {
		return "Name"
	}
	return s
}

func (s *FooService) UpdateFoo(ctx context.Context, req *pb.UpdateFooRequest) (*empty.Empty, error) {
	// use mask to update
	return nil, nil
}
func (s *FooService) DeleteFoo(ctx context.Context, req *pb.DeleteFooRequest) (*empty.Empty, error) {
	return nil, nil
}
func (s *FooService) GetFoo(ctx context.Context, req *pb.GetFooRequest) (*pb.GetFooReply, error) {
	return &pb.GetFooReply{}, nil
}
func (s *FooService) ListFoo(ctx context.Context, req *pb.ListFooRequest) (*pb.ListFooReply, error) {
	return &pb.ListFooReply{}, nil
}
