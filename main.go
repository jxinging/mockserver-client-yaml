package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jxinging/mockserver-client-yaml/mockyaml"
)

var mockserver = flag.String("s", "http://127.0.0.1:1080", "mock-server address")

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-s mock-server] dir\n", os.Args[0])
		os.Exit(1)
	}

	dir := flag.Args()[0]
	err := mockyaml.WalkYamlFiles(dir, func(apis []mockyaml.MockExpectation) error {
		for _, api := range apis {
			if err := mockyaml.CreateMockExpectation(*mockserver, &api); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}
