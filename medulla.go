package main

import (
	"fmt"
)

type MedullaDeployer struct {
}

func (medulla *MedullaDeployer) LogVersion() {
	fmt.Println(`
	
MedullaDeployer Deploy Version 1.0.0



	`)
}
