package utils

import "encoding/json"

func BindingStruct(src interface{}, desc interface{}) error {
	// convert to byte
	byteSrc, err := json.Marshal(src)
	if err != nil {
		return err
	}
	// binding to desc
	err = json.Unmarshal(byteSrc, &desc)
	if err != nil {
		return err
	}
	return nil
}
