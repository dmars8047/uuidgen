package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

func main() {
	v4 := flag.Bool("v4", false, "Generate a version 4 UUID")
	upper := flag.Bool("u", false, "Generate a UUID with uppercase letters")

	flag.Parse()

	if !(*v4) {
		val := uuid.New().String()

		if *upper {
			val = strings.ToUpper(val)
		}

		fmt.Println(val)

		return
	} else {
		val, err := uuid.NewRandom()

		if err != nil {
			log.Fatal(err)
		}

		var valstr string = val.String()

		if *upper {
			valstr = strings.ToUpper(valstr)
		}

		fmt.Println(valstr)

		return
	}
}
