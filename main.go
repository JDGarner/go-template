package main

import (
	"fmt"

	// TODO: these packages are just here so that go.sum has some files for the template, remove them on setup
	_ "github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	fmt.Println("Hello World")
}
