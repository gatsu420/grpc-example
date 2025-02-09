package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gatsu420/grpc-example/common/config"
	"github.com/gatsu420/grpc-example/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	zlog "github.com/rs/zerolog/log"
)

var taskList *model.TaskList

func init() {
	taskList = &model.TaskList{}
	taskList.List = []*model.Task{}
}

type TasksServer struct {
	model.UnimplementedTasksServer
}

func (ts *TasksServer) Create(_ context.Context, task *model.Task) (*emptypb.Empty, error) {
	taskList.List = append(taskList.List, task)
	zlog.Info().Msg(fmt.Sprintf("adding user %v", task.String()))

	return &emptypb.Empty{}, nil
}

func (ts *TasksServer) List(_ context.Context, _ *emptypb.Empty) (*model.TaskList, error) {
	zlog.Info().Msg(fmt.Sprintf("users: %v", taskList))

	return taskList, nil
}

func main() {
	srv := grpc.NewServer()
	model.RegisterTasksServer(srv, &TasksServer{})
	zlog.Info().Str(
		"message", fmt.Sprintf("starting RPC server at %v", config.ServiceTasksPort),
	)

	listener, _ := net.Listen("tcp", config.ServiceTasksPort)
	err := srv.Serve(listener)
	if err != nil {
		zlog.Fatal().Msg("server failed to start")
	}
}
