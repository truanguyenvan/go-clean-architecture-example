package validator

import (
	"go-clean-architecture-example/pkg/utils"
	"sync"
)

var lock = &sync.Mutex{}

var validatorInstance utils.StructValidator

func GetValidator() utils.StructValidator {
	if validatorInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if validatorInstance == nil {
			validatorInstance = utils.NewStructValidator()
		}
	}

	return validatorInstance
}
