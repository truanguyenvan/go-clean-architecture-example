package decorator

import (
	"fmt"
	"strings"
)

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
