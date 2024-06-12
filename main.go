package main

import (
	"fmt"
	"log"
	"net-http-api/coincap/coincap"
	"time"
)

func main() {
	coincapClient, err := coincap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	bitcoin, err := coincapClient.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bitcoin.Info())
}
