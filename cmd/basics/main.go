package main

import (
	"fmt"
	"github.com/AbdulfatahMohammedSheikh/backend/core/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/home/abdulfatah/devlopment/projects/ethiopiaHr/ethiopia_hr/.env")

	if nil != err {
		panic(err.Error())
	}

	fmt.Println(config.GetConigVirable("ADDRESS"))
	fmt.Println(config.GetConigVirable("USER"))
	fmt.Println(config.GetConigVirable("PASS"))
	fmt.Println(config.GetConigVirable("NAMESPACE"))
	fmt.Println(config.GetConigVirable("DATABASE"))
}
