package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/MihaiBlebea/go-watcher/builder"
	"github.com/MihaiBlebea/go-watcher/snapshot"
)

func main() {

	rootPath := flag.String("root", ".", "Root folder of your application")
	buildCmd := flag.String("build", "go build .", "Command to build your application")
	interval := flag.Int("interval", 5, "Interval in sec when the check for updated files should run")
	flag.Parse()

	fmt.Printf("Watching folder: %s | Build command: %s\n", *rootPath, *buildCmd)

	service := snapshot.New(*rootPath)

	err := service.Watch(time.Duration(*interval) * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case res := <-service.C:
			fmt.Println("Updating...")

			for _, f := range res.Files {
				fmt.Printf("File changed: %s | mode %v\n", f.Path(), f.Method())
			}

			if res.Err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Running command: %s\n", *buildCmd)

			cmd, args := builder.ParseCmd(*buildCmd)

			out, err := builder.RunCmd(cmd, args...)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Building output: %s\n", out)
		}
	}
}
