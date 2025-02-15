package main

import (
	"fmt"
	"log"

	"github.com/alexKudryavtsev-web/default-service/config"
)

func main() {
	cnf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("can't init config: %s", err)
	}

	fmt.Println(cnf)
}
