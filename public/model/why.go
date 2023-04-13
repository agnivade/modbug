package model

import (
	"modbug/v7/inner"
)

func Why() string {
	return "why" + inner.Inner()
}
