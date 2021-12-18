package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/sanoyo/mini-lens-go/proto"
	"google.golang.org/grpc"
)

var (
	endpoint   = "localhost:8080"
	listenAddr = ":3000"
)

func init() {
	ep := os.Getenv("ENDPOINT")
	if ep != "" {
		endpoint = ep
	}
}

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, endpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

// Run starts a HTTP server and blocks forever if successful.
func Run(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}

	return http.ListenAndServe(address, gw)
}

func main() {
	fmt.Println("Listen Address:", listenAddr)
	if err := Run(listenAddr); err != nil {
		panic(err)
	}
}
