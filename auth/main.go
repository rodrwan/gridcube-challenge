package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rodrwan/gridcube-challenge/publisher/service"
	"google.golang.org/grpc"
)

const imageSize = 621

func main() {
	host := flag.String("publisher-host", "publisher", "publisher listening host")
	port := flag.Int64("publisher-port", 8091, "publisher listening port")

	flag.Parse()

	// Initialize grpc connection
	IP := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Publisher IP: %s\n", IP)
	conn, err := grpc.Dial(IP, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	sc := service.NewServiceClient(conn)

	// Serve photoHandler
	http.HandleFunc("/photo", photoHandler(sc))

	log.Println("running auth service")
	http.ListenAndServe(":8090", nil)
}

// request hold request information
type request struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Caption  string `json:"caption"`
}

// photoHandler consume request from /photo this endpoint receive a ServiceClient which allow us
// to post image on instagram
func photoHandler(sc service.ServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var req request
		if err := json.Unmarshal(body, &req); err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := sc.UploadPicture(ctx, &service.GetRequest{
			Size:     imageSize,
			Username: req.Username,
			Password: req.Password,
			Caption:  req.Caption,
		})
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Accept-Encoding", "gzip")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp.GetStatus())
	}
}
