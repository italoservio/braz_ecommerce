package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils_ParseDocument(t *testing.T) {
	type SampleStruct struct {
		Name  string `bson:"name"`
		Value int    `bson:"value"`
	}

	t.Run("should parse structure when document is valid", func(t *testing.T) {
		structure := SampleStruct{
			Name:  "TestName",
			Value: 42,
		}

		_, err := ParseToDocument(structure)
		if err != nil {
			t.Fatal(err)
		}

		assert.Nil(t, err, "should not return error")
	})

	t.Run("should fail marshaling when inputted structure is invalid", func(t *testing.T) {
		wrongStructure := ""

		_, err := ParseToDocument(wrongStructure)
		if err == nil {
			t.Fatal(err)
		}

		assert.NotNil(t, err, "should return marshal error")
	})
}