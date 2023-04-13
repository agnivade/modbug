package modbug

import "modbug/public/v7/model"

func Hi() string {
	return "Hi" + model.Why()
}
