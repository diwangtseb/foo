package service

import (
	"context"
	"foo/internal/biz"

	pb "foo/api/foo/v1"
)

type FooService struct {
	pb.UnimplementedFooServer
	fuc *biz.FooUseCase
}

func NewFooService(fuc *biz.FooUseCase) *FooService {
	return &FooService{fuc: fuc}
}

func (s *FooService) CreateFoo(ctx context.Context, req *pb.CreateFooRequest) (*pb.CreateFooReply, error) {
	r, err := s.fuc.CreateFoo(ctx, &biz.Foo{Msg: req.Name})
	if err != nil {
		return nil, pb.ErrorUnknown("create foo failed")
	}
	return &pb.CreateFooReply{
		Msg: r.Msg,
	}, nil
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
