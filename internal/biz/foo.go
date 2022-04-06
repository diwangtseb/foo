package biz

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/log"
)

type Foo struct {
	Msg string
}

type FooUseCase struct {
	log *log.Helper
}

func NewFooUseCase(logger log.Logger) *FooUseCase {
	return &FooUseCase{
		log: logger,
	}
}

func (f *FooUseCase) CreateFoo(ctx context.Context, foo *Foo) (*Foo, error) {
	f.log.WithContext(ctx).Infof("CreateFoo: %v", foo.Msg)
	return foo, nil
}
