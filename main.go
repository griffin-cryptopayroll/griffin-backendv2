package main

import (
	"fmt"
	"griffin-dao/start"
)

func main() {
	dao, err := start.NewDao()
	fmt.Println(err)
	fmt.Println(dao)
}
