package main

import (
	"fmt"

	"github.com/1ssk/storage/internal/storage"
)

func main() {

	st := storage.NewStorage()

	fmt.Println("Hi", st)
}
