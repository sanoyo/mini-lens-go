package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/sanoyo/mini-lens-go/proto"
	"google.golang.org/grpc"
)

var (
	endpoint   = "localhost:8090"
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

	err = pb.RegisterPodServiceHandlerFromEndpoint(ctx, mux, endpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))

	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))

	glog.Infof("preflight request for %s", r.URL.Path)
}

// ref: https://github.com/rephus/grpc-gateway-example/blob/f9caf363dd6c438968379c548651748de6254bd7/main.go#L68
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// Run starts a HTTP server and blocks forever if successful.
func Run(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	gw, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	return http.ListenAndServe(address, allowCORS(mux))
}

func main() {
	fmt.Println("Listen Address:", listenAddr)
	if err := Run(listenAddr); err != nil {
		panic(err)
	}
}
