package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
	Exported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println("-----printing marshaled person object-----")
	fmt.Println(bs)
	fmt.Println("-----printing marshaled person type-----")
	fmt.Printf("%T \n", bs)
	fmt.Println("-----printing marshaled person string-----")
	fmt.Println(string(bs))
}
