package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/rodrwan/gridcube-challenge/publisher/instagram"
	"github.com/rodrwan/gridcube-challenge/publisher/service"
	image "github.com/rodrwan/gridcube-challenge/random-image"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int64("port", 8091, "listening port")

	flag.Parse()

	srv := grpc.NewServer()
	service.RegisterServiceServer(srv, &PublisherService{
		ImageService: image.NewService(200),
	})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Starting Publisher service...")
	log.Println(fmt.Sprintf("Publisher service, Listening on: %d", *port))
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// PublisherService is the implementation of the service.ServiceServer interface
type PublisherService struct {
	ImageService *image.Service
}

// UploadPicture implementation for the same method of service.ServiceServer interface
func (ps *PublisherService) UploadPicture(ctx context.Context, in *service.GetRequest) (*service.GetResponse, error) {
	ss, err := instagram.NewSession(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	// get a random image
	img, err := ps.ImageService.Get()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// upload the image
	if err := ss.UploadPhoto(img, in.Caption); err != nil {
		return nil, err
	}

	// close the session
	defer ss.Close()

	return &service.GetResponse{
		Status: "OK",
	}, nil
}
