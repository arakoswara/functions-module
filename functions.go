package functions_module

import (
	"strings"
)

func Ucwords(str string) string {
	return strings.ToTitle(str)
}
