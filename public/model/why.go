package model

import (
	"github.com/agnivade/modbug/v7/inner"
)

func Why() string {
	return "why" + inner.Inner()
}
