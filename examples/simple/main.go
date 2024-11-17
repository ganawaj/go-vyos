package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ganawaj/go-vyos/vyos"
)

func main() {

	// Create a new VyOS API client.
	c := vyos.NewClient(nil).WithToken("AUTH_KEY").WithURL("https://10.1.1.1")

	ctx := context.Background()

	r, _, err := c.ConfigFile.Get(ctx, "interfaces")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Error: %s\n", r.Error)
	fmt.Printf("Success: %v\n", r.Success)
	fmt.Printf("Data: %v\n", r.Data)

	b, err := json.MarshalIndent(r.Data, "", "  ")
	if err != nil {
			fmt.Println("error:", err)
	}
	fmt.Print(string(b))

}