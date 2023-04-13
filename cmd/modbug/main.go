package main

import (
	"fmt"
	modbug "modbug/v7"
	"modbug/v7/inner"
)

func main() {
	fmt.Println(modbug.Hi())
	fmt.Println(inner.Inner())
}
