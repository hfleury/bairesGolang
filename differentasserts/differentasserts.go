package differentasserts

import "errors"

func AssertEqual(numa, numb int) bool {
	return numa == numb
}

func AssertNil(name string) error {
	if name != "Henrique" {
		return errors.New("Can't work without Henrique")
	}

	return nil
}

func AssertNotNil(user string) bool, error {
	if user != "bairesdev" {
		return nil, errors.New("Username not authorized")
	}
	return true, nil
} 
