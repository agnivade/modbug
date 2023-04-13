package modbug

import (
	"fmt"

	"github.com/agnivade/pgm"
)

func PGM() string {
	index := pgm.NewIndex([]float64{1.0, 2.0}, 2)
	return fmt.Sprintf("%v", index)
}
