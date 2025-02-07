package main

import (
	"context"
	"log"
	"net"

	"github.com/gatsu420/grpc-example/common/config"
	"github.com/gatsu420/grpc-example/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.UserList

func init() {
	localStorage = &model.UserList{}
	localStorage.List = []*model.User{}
}

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (us UsersServer) Register(_ context.Context, param *model.User) (*emptypb.Empty, error) {
	user := param
	localStorage.List = append(localStorage.List, user)
	log.Println("registering user ", user.String())

	return &emptypb.Empty{}, nil
}

func (us UsersServer) List(context.Context, *emptypb.Empty) (*model.UserList, error) {
	log.Println(localStorage)
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	var userSrv UsersServer
	model.RegisterUsersServer(srv, userSrv)

	log.Println("starting RPC server at ", config.ServiceUserPort)

	l, _ := net.Listen("tcp", config.ServiceUserPort)
	log.Fatal(srv.Serve(l))
}
