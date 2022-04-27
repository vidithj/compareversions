package main

import (
	"compareversions/api"
	"fmt"
	"log"
	"os"
)

func main() {

	//check length of args
	if len(os.Args) < 3 {
		log.Fatal("Minimum 2 versions required for comparison")
	}

	//return the compare result. If v1 > v2 result is 1 , if v1 < v2 result is -1 else result is 0
	fmt.Println(api.CompareVersions(os.Args[1], os.Args[2]))

}
