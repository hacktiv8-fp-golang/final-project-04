package helper

import "strings"

func ParseError(err error) Error {
	if strings.Contains(err.Error(), "record not found") {
		return NotFound("Data not found")
	}

	return InternalServerError("Something went wrong")
}
