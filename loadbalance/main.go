package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Client struct {
	Name string
}

type LoadBalance struct {
	ClientArr []Client
	Size      int
}

func getLoadBalance(size int) *LoadBalance {
	var clients []Client
	for i := 1; i <= size; i++ {
		clients = append(clients, Client{Name: "load" + strconv.Itoa(i)})
	}
	return &LoadBalance{ClientArr: clients, Size: size}
}

func (l *LoadBalance) getClients() *Client {
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(100)
	return &l.ClientArr[intn%l.Size]
}

func (c *Client) Do() {
	fmt.Println(c.Name)
}

func main() {
	//随机负载均衡算法
	balance := getLoadBalance(5)
	balance.getClients().Do()
}
