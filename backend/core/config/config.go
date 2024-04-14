package config

import (
	"os"

	"github.com/joho/godotenv"
)


func GetConigVirable(key string) string {
	err := godotenv.Load("/home/abdulfatah/devlopment/projects/ethiopiaHr/ethiopia_hr/.env")

	if nil != err {
		panic(err.Error())
	}

	k := os.Getenv(key)
	return k
}
