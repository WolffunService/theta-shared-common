package rcfg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByUser(t *testing.T) {
	env := Staging
	name := "abtest.rivals.tutorial"
	user := UserContext{
		UserID: "123",
		Attributes: map[string]interface{}{
			"status": 1,
		},
	}
	option := Option{
		DisablePushEvent: true,
		Country:          "VN",
	}
	request := GetByUserRequest{
		User:   user,
		Option: option,
	}

	// Define the expected result
	expectedResult := map[string]interface{}{
		"hasFlagTest":          true,
		"tr_newbie_battle":     15,
		"tr_real_battle":       51,
		"tr_tutorial_battle":   5,
		"tr_tutorial_with_bot": true,
	}

	// Marshal the mock response to JSON

	// Call the function being tested
	result, err := GetByUser[map[string]any](env, name, request)
	// Check the result
	assert.NoError(t, err)
	//assert.Equal(t, expectedResult, *result)
	assert.EqualValues(t, fmt.Sprint(expectedResult), fmt.Sprint(*result))
}
