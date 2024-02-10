package addborder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBorder(t *testing.T) {

	picture := []string{"abc", "ded"}
	expected := []string{"*****", "*abc*", "*ded*", "*****"}

	res := addBorder(picture)

	if len(res) != len(expected) {
		assert.Fail(t, "Expected and result slices have different length")
	}

	for i := range res {
		assert.Equal(t, expected[i], res[i])
	}
}
