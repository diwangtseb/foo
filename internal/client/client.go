package client

import (
	"context"
	v1 "foo/api/foo/v1"
	"log"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
)

func discover() {
	// new consul client
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(consulCli)

	// new grpc client
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///foo"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	gClient := v1.NewFooClient(conn)
	gClient.UpdateFoo(context.Background(), &v1.UpdateFooRequest{
		Foobar: &v1.FooBar{
			Name:     "11",
			Describe: "22",
		},
	})

	// new http client
	hConn, err := http.NewClient(
		context.Background(),
		http.WithMiddleware(
			recovery.Recovery(),
		),
		http.WithEndpoint("discovery:///foo"),
		http.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer hConn.Close()
	hClient := v1.NewFooHTTPClient(hConn)
	hClient.CreateFoo(context.TODO(), &v1.CreateFooRequest{
		Name: "112",
	})
}
