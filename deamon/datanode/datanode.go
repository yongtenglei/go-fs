package datanode

import (
	"errors"
	"fmt"
	"go-fs/datanode"
	datanode_pb "go-fs/proto/datanode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func InitializeDataNodeUtil(serverPort int, dataLocation string) {
	dataNodeInstance := new(datanode.Service)
	dataNodeInstance.DataDirectory = dataLocation
	dataNodeInstance.ServicePort = uint16(serverPort)

	log.Printf("Data storage location is %s\n", dataLocation)

	server := grpc.NewServer()
	datanode_pb.RegisterDataNodeServer(server, dataNodeInstance)

	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 创建监听器, 如果端口被占用, 则端口号+1, 直至找到空闲端口
	var listener net.Listener
	initErr := errors.New("init")

	for initErr != nil {
		listener, initErr = net.Listen("tcp", ":"+strconv.Itoa(serverPort))
		serverPort += 1
	}
	log.Printf("DataNode port is %d\n", serverPort-1)
	defer listener.Close()

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Printf(fmt.Sprintf("Server Serve failed in %s", ":"+strconv.Itoa(serverPort)), "err", err.Error())
			panic(err)
		}
	}()

	log.Println("DataNode daemon started on port: " + strconv.Itoa(serverPort-1))

	// graceful shutdown
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	<-sig

	server.GracefulStop()
}
