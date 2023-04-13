package main

import (
	"fmt"
	modbug "github.com/agnivade/modbug/v7"
	"github.com/agnivade/modbug/v7/inner"
)

func main() {
	fmt.Println(modbug.Hi())
	fmt.Println(modbug.PGM())
	fmt.Println(inner.Inner())
}
