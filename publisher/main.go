package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/rodrwan/gridcube-challenge/publisher/instagram"
	"github.com/rodrwan/gridcube-challenge/publisher/service"
	image "github.com/rodrwan/gridcube-challenge/random-image"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int64("port", 8091, "listening port")

	flag.Parse()

	srv := grpc.NewServer()
	service.RegisterServiceServer(srv, &ImageService{})
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

type response struct {
	Data interface{} `json:"data"`
}

func getImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "This server does not support that HTTP method", http.StatusBadRequest)
			return
		}

		img, err := image.Get()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "could not get image", http.StatusInternalServerError)
			return
		}

		resp := &response{
			Data: b64.StdEncoding.EncodeToString(img),
		}

		w.Header().Set("Accept-Encoding", "gzip")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

// ImageService is the implementation of the service.ServiceServer interface
type ImageService struct{}

// UploadPicture implementation for the same method of service.ServiceServer interface
func (is *ImageService) UploadPicture(ctx context.Context, in *service.GetRequest) (*service.GetResponse, error) {
	ss, err := instagram.NewSession(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	// get a random image
	img, err := image.Get()
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
