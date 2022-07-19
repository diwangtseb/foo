package middleware

import (
	"context"
	"errors"
	"fmt"
	v1 "foo/api/foo/v1"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
)

func naming(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FieldMask() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			fmt.Println("FieldMask middleware")
			//Do something before request
			if v, ok := req.(*v1.UpdateFooRequest); ok {
				//use mask to update
				mask, err := fieldmask_utils.MaskFromProtoFieldMask(v.UpdateMask, naming)
				if err != nil {
					return nil, errors.New("invalid mask")
				}
				// dst := make(map[string]interface{})
				dst := &v1.UpdateFooRequest{
					Foobar: &v1.FooBar{},
				}
				err = fieldmask_utils.StructToStruct(mask, v.Foobar, dst.Foobar)
				if err != nil {
					return nil, errors.New("invalid mask")
				}
				v.Foobar = dst.Foobar
			}
			return handler(ctx, req)
		}
	}
}
