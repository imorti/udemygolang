package main

import (

	"github.com/imorti/udemygolang/04_scope/01_package_scope/02_visibility/vis"
	"fmt"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}