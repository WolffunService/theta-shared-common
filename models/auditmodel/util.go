package auditmodel

import (
	"strconv"
)

func stringToInt(inp string, defaultV int) (int, error) {
	res, err := strconv.Atoi(inp)
	if err != nil {
		return defaultV, err
	}

	return res, nil
}
