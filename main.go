package main

import (
	"bufio"
	"flag"
	"go-fs/deamon/client"
	"go-fs/deamon/datanode"
	"go-fs/deamon/namenode"
	"go-fs/pkg/util"
	"log"
	"os"
	"strings"
)

func main() {
	dataNodeCommand := flag.NewFlagSet("datanode", flag.ExitOnError)
	nameNodeCommand := flag.NewFlagSet("namenode", flag.ExitOnError)
	clientCommand := flag.NewFlagSet("client", flag.ExitOnError)

	dataNodePortPtr := dataNodeCommand.Int("port", 7000, "DataNode communication port")
	dataNodeDataLocationPtr := dataNodeCommand.String("data-location", ".", "DataNode data storage location")

	nameNodePortPtr := nameNodeCommand.Int("port", 9000, "NameNode communication port")
	nameNodeListPtr := nameNodeCommand.String("datanodes", "", "Comma-separated list of DataNodes to connect to")
	nameNodeBlockSizePtr := nameNodeCommand.Int("block-size", 32, "Block size to store")
	nameNodeReplicationFactorPtr := nameNodeCommand.Int("replication-factor", 1, "Replication factor of the system")

	clientNameNodePortPtr := clientCommand.String("namenode", "localhost:9000", "NameNode communication port")
	clientOperationPtr := clientCommand.String("operation", "", "Operation to perform")
	clientSourceFilePathPtr := clientCommand.String("source-file-path", "", "Source path of the file\n\t required: GET, PUT")
	clientDestFilePathPtr := clientCommand.String("dest-file-path", "", "Destination path of the file\n\t required: GET, PUT")

	if len(os.Args) < 2 {
		log.Println("sub-command is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "datanode":
		_ = dataNodeCommand.Parse(os.Args[2:])
		datanode.InitializeDataNodeUtil(*dataNodePortPtr, *dataNodeDataLocationPtr)

	case "namenode":
		_ = nameNodeCommand.Parse(os.Args[2:])
		var listOfDataNodes []string
		if len(*nameNodeListPtr) > 1 {
			listOfDataNodes = strings.Split(*nameNodeListPtr, ",")
		} else {
			listOfDataNodes = []string{}
		}
		namenode.InitializeNameNodeUtil(*nameNodePortPtr, *nameNodeBlockSizePtr, *nameNodeReplicationFactorPtr, listOfDataNodes)

	case "client":
		_ = clientCommand.Parse(os.Args[2:])

		if *clientOperationPtr == "put" {
			if *clientNameNodePortPtr == "" || *clientSourceFilePathPtr == "" || *clientDestFilePathPtr == "" {
				log.Println("Need required arguments, use \"go-fs client -help\" for more details")
				return
			}
			status := client.PutHandler(*clientNameNodePortPtr, *clientSourceFilePathPtr, *clientDestFilePathPtr)
			log.Printf("Put status: %t\n", status)

		} else if *clientOperationPtr == "get" {
			if *clientNameNodePortPtr == "" || *clientSourceFilePathPtr == "" || *clientDestFilePathPtr == "" {
				log.Println("Need required arguments, use \"go-fs client -help\" for more details")
				return
			}
			contents, status := client.GetHandler(*clientNameNodePortPtr, *clientSourceFilePathPtr)
			log.Printf("Get status: %t\n", status)
			if !status {
				log.Println("Please check remote file path")
				return
			}
			log.Println(contents)
			fileWriteHandler, err := os.Create(*clientDestFilePathPtr)
			util.Check(err)
			defer fileWriteHandler.Close()

			fileWriter := bufio.NewWriter(fileWriteHandler)
			_, err = fileWriter.WriteString(contents)
			util.Check(err)
			fileWriter.Flush()
		}
	}
}
