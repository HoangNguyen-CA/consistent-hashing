package main

import (
	"fmt"

	"github.com/HoangNguyen-CA/consistent-hashing/internal/hash"
	"github.com/HoangNguyen-CA/consistent-hashing/internal/server"
)

func main() {
	hr := hash.NewHashRing(100)
	servers := setUpServers([]string{"s1", "s2", "s3"})

	serverMap := make(map[string]int)

	for _, s := range servers {
		hr.AddServer(s)
		serverMap[s.Id] = 0
	}

	totalRequests := 100000

	for i := 0; i < totalRequests; i++ {
		routedServerId := hr.GetServer([]byte(fmt.Sprint(i))).Id
		serverMap[routedServerId]++
	}

	for key, val := range serverMap {
		fmt.Printf("Server %v got %v requests\n", key, val)
	}
}

func setUpServers(serverIds []string) []*server.Server {
	servers := []*server.Server{}

	for _, serverId := range serverIds {
		server := server.NewServer(serverId)
		servers = append(servers, server)
	}

	return servers
}
