package main

import (
	"context"

	"github.com/gatsu420/grpc-example/common/config"
	"github.com/gatsu420/grpc-example/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func serviceUser() model.UsersClient {
	port := config.ServiceUserPort
	conn, _ := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return model.NewUsersClient(conn)
}

func main() {
	user1 := &model.User{
		Id:       "qweqweqwe",
		Name:     "cayoyoo",
		Password: "asdasd980",
		Gender:   model.UserGender_MALE,
	}

	user := serviceUser()
	user.Register(context.Background(), user1)
	user.List(context.Background(), &emptypb.Empty{})
}
