package main

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc/server/pb"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

var (
	router   *mux.Router
	grpcPort = "9090"
)

func main() {

	grpcServer := grpc.NewServer()

	// attach the todo service to the server
	pb.RegisterGreetServiceServer(grpcServer, &Server{})

	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithAllowedRequestHeaders([]string{"*"}))

	router = mux.NewRouter()

	router.HandleFunc("/sayHelloRest", SayHelloHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", grpcPort), http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		if strings.Contains(request.Header.Get("Content-Type"), "application/json") {
			router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
			})
			router.ServeHTTP(writer, request)
			return
		}
		if strings.Contains(request.Header.Get("Content-Type"), "application/grpc-web-text") {
			wrappedServer.ServeHTTP(writer, request)
			return
		}
	})))
}

// gRPC server
type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {

	response := "Hello " + request.User.FirstName + " " + request.User.LastName

	return &pb.UserResponse{Message: response}, nil
}

func SayHelloHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	userFullName := *&pb.UserFullName{}
	json.Unmarshal(reqBody, &userFullName)

	response := "Hello " + userFullName.FirstName + " " + userFullName.LastName

	json.NewEncoder(w).Encode(response)
}
