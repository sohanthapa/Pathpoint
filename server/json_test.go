package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestIsJSON tests the functionality for IsJSON function
func TestIsJSON(t *testing.T) {
	t.Run("error - invalid json format", func(t *testing.T) {
		isJSON := IsJSON("invalid")
		assert.Equal(t, false, isJSON)
	})

	t.Run("sucess - valid json format", func(t *testing.T) {
		validJSON := `{"name":"John", "age":30, "car":null}`
		isJSON := IsJSON(validJSON)
		assert.Equal(t, true, isJSON)
	})
}

//TestIsJSON tests the functionality for UnMarshallJSON function
func TestUnMarshallJSON(t *testing.T) {
	t.Run("error - unmarshalling string", func(t *testing.T) {
		jsonString := `{"name": 1234, "game":"12324444"}`
		_, err := UnMarshallJSON(jsonString)
		assert.NotNil(t, err)

	})
	t.Run("success - unmarshalling string", func(t *testing.T) {
		jsonString := `{"name":"John", "id":"12324444"}`
		expectedBody := DataFileInput{
			ID:   "12324444",
			Name: "John",
		}
		dataFileInput, err := UnMarshallJSON(jsonString)
		assert.Nil(t, err)
		assert.Equal(t, expectedBody, dataFileInput)

	})
}

//TestWriteJSON tests the functionality for WriteJSON function
func TestWriteJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_, err := WriteJSON(`{"name": "Sohan"}`)
		assert.Nil(t, err)
	})
}
