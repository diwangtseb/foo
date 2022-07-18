package biz

import (
	"context"

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
		log: log.NewHelper(logger),
	}
}

func (f *FooUseCase) CreateFoo(ctx context.Context, foo *Foo) (*Foo, error) {
	f.log.Debugw("____msg", foo.Msg)
	return foo, nil
}
