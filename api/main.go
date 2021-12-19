package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"

	pb "github.com/sanoyo/mini-lens-go/proto"
	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const (
	port = ":8080"
)

type healthServiceClient struct {
	pb.UnimplementedHealthServiceServer
}

type podServiceClient struct {
	pb.UnimplementedPodServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterHealthServiceServer(s, &healthServiceClient{})
	pb.RegisterPodServiceServer(s, &podServiceClient{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *healthServiceClient) GetStatus(ctx context.Context, in *pb.Empty) (*pb.AliveResponse, error) {
	return &pb.AliveResponse{Status: true}, nil
}

func (s *podServiceClient) GetPodStatus(ctx context.Context, in *pb.PodEmpty) (*pb.PodResponse, error) {
	var kubeconfig *string

	namespace := "default"
	status := false

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pod := "sample"
	_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			pod, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		status = true
	}

	return &pb.PodResponse{
		Name:   pod,
		Status: status,
	}, nil
}
