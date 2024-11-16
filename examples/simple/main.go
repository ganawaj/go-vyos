package main

import (
	"context"
	"fmt"

	"github.com/ganawaj/go-vyos/vyos"
)

func main() {

	// Create a new VyOS API client.
	c := vyos.NewClient(nil).WithToken("AUTH_KEY").WithURL("https://10.1.1.1")

	fmt.Println(c.Token)

	ctx := context.Background()

	r, _, err := c.Show.Do(ctx, "system image")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Error: %s\n", r.Error)
	fmt.Printf("Success: %v\n", r.Success)
	fmt.Printf("Data: %v\n", r.Data)

}