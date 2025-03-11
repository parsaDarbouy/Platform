package main

import (
	"fmt"
	"os"
	"github.com/parsaDarbouy/Platform/pkg"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var val_back string = "val_back" 
var replicas = 1
var repository = "docker.io/redis"
var tag = "latest"
var pullPolicy = "IfNotPresent"
var containerPort = 6379
var serviceType = "ClusterIP"
var port = 6379
var targetPort = 6379

func main() {
	file, err := os.OpenFile(val_back+".yaml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	check(err)
	defer file.Close()

	_,err = file.Write([]byte(pkg.MakeValues(replicas, repository, tag, pullPolicy, containerPort, serviceType, port, targetPort)))
	check(err)

	fmt.Printf("File with %s name created or overwritten\n", file.Name())
}