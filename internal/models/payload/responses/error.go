package response

import (
	"fmt"
)

const (
	// Autentication errors
	InvalidToken = "Invalid / expire token"
	AccessDenied = "Access denied"

	// Payload errors
	BADREQUEST = "Invalid request"

	// Error wordings
	INVL string = "Invalid"
	UNKN string = "Unknown"
	MTCh string = "Match"
	MDTR string = "Mandatory"
)

var (
	ErrorMapper = map[string]string{
		INVL: "has invalid format",
		UNKN: "is unknown",
		MTCh: "does not match",
		MDTR: "Cannot be empty",
	}
)

func ErrorBuilder(field, err string) error {
	return fmt.Errorf("%s %s", field, ErrorMapper[err])
}
