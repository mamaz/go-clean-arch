package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnHash(t *testing.T) {
	user := &User{
		ID:   "1",
		Name: "Hisma",
		City: "Cimahi",
	}

	other := &User{
		ID:   "1",
		Name: "Hisma",
		City: "Cimahi",
	}

	assert.Equal(t, other.Hash(), user.Hash())
}
