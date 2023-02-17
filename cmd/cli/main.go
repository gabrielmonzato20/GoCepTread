package main

import (
	"fmt"
	"time"

	"os"

	"github.com/gabrielmonzato20/GoCepTread/config"
	"github.com/gabrielmonzato20/GoCepTread/internal/entity"
	"github.com/gabrielmonzato20/GoCepTread/internal/infra/webserver"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	handler := webserver.NewHandler(config.EndPointServer1, config.EndPointServer2)
	c1 := make(chan *entity.ResponseEntity)
	c2 := make(chan *entity.ResponseEntity)

	go func() {
		c1 <- handler.CallFistServer(os.Args[1])
	}()
	go func() {
		c2 <- handler.CallSecondServer(os.Args[1])
	}()


		select {
		case data := <-c1:
			fmt.Printf("Received from first web server: url: %s - %s\n", data.ApiResponse, data.Response)

		case data := <-c2:
			fmt.Printf("Received from second web server: url: %s - %s\n", data.ApiResponse, data.Response)

		case <-time.After(time.Second * 1):
			println("timeout")
		}
	
}
