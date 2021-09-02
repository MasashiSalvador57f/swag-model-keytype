package main

import (
	"flag"
	"fmt"
)

var (
	swaggerFileOpt = flag.String("f", "./swagger.yaml", "path to swagger file")
)

func main() {
	flag.Parse()
	fmt.Println(*swaggerFileOpt)
}
