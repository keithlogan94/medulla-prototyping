package main

import "fmt"

type MedullaDeployer struct {
	MinioClientWrapper MinioWrapper
}

var wrapper MinioWrapper
var deployer = MedullaDeployer{}

func init() {
	fmt.Println("initializing Minio Client Wrapper")
	wrapper.init()
	deployer.MinioClientWrapper = wrapper
}

func (medulla *MedullaDeployer) LogVersion () {
	fmt.Println(`
	
MedullaDeployer Deploy Version 1.0.0



	`)
}



