package main

func main() {
	deployer.LogVersion()
	deployer.MinioClientWrapper.Deploy()
}
