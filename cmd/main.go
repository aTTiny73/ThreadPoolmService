package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/aTTiny73/ThreadPoolmService/internal/ping"
	"github.com/aTTiny73/ThreadPoolmService/pkg/pool"
)

// Function that handles loading and unmarshaling from a config file.
func loadFromfile(path string) pool.Hosts {
	hosts := pool.Hosts{}
	config, _ := os.Open(path)
	bytevalue, _ := ioutil.ReadAll(config)
	err := json.Unmarshal(bytevalue, &hosts)
	if err != nil {
		fmt.Println(err)
	}
	return hosts
}

func main() {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	context, stopCoordinator := context.WithCancel(context.Background())
	pool.CoordinatorInstance.Ctx = context

	var hosts = loadFromfile("config.json")
	dataByte, _ := json.Marshal(hosts)
	fmt.Println(hosts)

	/*
		// Adds workers equal to the number of CPUs
			for i := 0; i < runtime.GOMAXPROCS(runtime.NumCPU()); i++ {
				go pool.CoordinatorInstance.Run()
			}
	*/

	go pool.CoordinatorInstance.Run()

	pool.CoordinatorInstance.Enqueue(ping.Pinger, dataByte)

	<-stop
	stopCoordinator()
	pool.Wg.Wait()
}
