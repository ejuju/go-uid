package main

import (
	"fmt"

	"github.com/ejuju/go-uid"
)

func main() {
	id := uid.Must(uid.NewUID(10))
	fmt.Println(id)

	id, err := uid.NewUID(34)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
