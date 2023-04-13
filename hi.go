package modbug

import "github.com/agnivade/modbug/public/v7/model"

func Hi() string {
	return "Hi" + model.Why()
}
