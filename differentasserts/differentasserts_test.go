package differentasserts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NotEqual | Nil | NotNil

func TestDifferentAsserts(t *testing.T) {

	equalbool := AssertEqual(1, 1)
	assert.Equal(t, true, equalbool, "The number are equal!")
	assert.NotEqual(t, 908, 987, "They shouldn't be equal")

	nilerror := AssertNil("Henrique")
	assert.Nil(t, nilerror)

	_, error := AssertNotNil("bairesdev")
	if assert.NotNil(t, error) {
		fmt.Println(error)
	}

}
