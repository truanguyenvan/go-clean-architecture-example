package structure

import (
	"github.com/jinzhu/copier"
)

// Copy
func Copy(s, ts interface{}) error {
	return copier.Copy(ts, s)
}
