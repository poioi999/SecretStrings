package main

import (
	"bufio"
	//	"net/rpc"
	"flag"
	"net/rpc"
	"os"

	//	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
	"fmt"

	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	file, err := os.Open("wordlist")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		request := stubs.Request{Message: line}
		response := new(stubs.Response)
		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Responded : " + response.Message)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	//TODO: connect to the RPC server and send the request(s)
}
