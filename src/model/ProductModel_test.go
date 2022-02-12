package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateProductInputRequiredFields(t *testing.T) {
	jsondata := `{"Name":"Prodct Name","Description":"Test product description","Picture":"image/foto.png"}`
	product := &Product{}
	if err := json.Unmarshal([]byte(jsondata), &product); err != nil {
		t.Errorf("failed to unmarshal product data %v", err.Error())
	}
	// fmt.Println("------------------", user)
	expected := ""
	if err := product.Validate(); err != nil {
		fmt.Println("------------------", err.Message())
		expected = "Invalid Product Name"
		if err.Message() == expected {
			assert.EqualValues(t, "", err.Message(), "Error validating Product name")
		}
		expected = "Invalid Description"
		if err.Message() == expected {
			assert.EqualValues(t, "", err.Message(), "Error validating Description")
		}
		expected = "Invalid Picture"
		if err.Message() == expected {
			assert.EqualValues(t, "", err.Message(), "Error validating Picture")
		}
	}

}
