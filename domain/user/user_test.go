package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashing(t *testing.T) {
	t.Run("it should be able to return user's hash", func(t *testing.T) {
		user := &User{
			ID:   1,
			Name: "Hisma",
			City: "Cimahi",
		}

		other := &User{
			ID:   1,
			Name: "Hisma",
			City: "Cimahi",
		}

		assert.Equal(t, other.Hash(), user.Hash())
	})
}
