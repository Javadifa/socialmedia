package handle

import "fmt"

func iaHandleValid(handle string) (bool, error) {
	//TODO - technical debt - is handle valid
	return false, fmt.Errorf("handle is not valid")
}

func isHandleUnique(handle string) (bool, error) {
	//TODO - technical debt - is handle UNIQUE
	return false, fmt.Errorf("handle is not unique")
}
