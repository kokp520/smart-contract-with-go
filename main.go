package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//ethereum env adress
	client, err := ethclient.Dial("http://127.0.0.1:56619")
	// fmt.Println(client)

	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
		panic(err)
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}

	// quick

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(header.Number.String())
	}
}
