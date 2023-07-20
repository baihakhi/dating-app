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
	ULGN string = "Login"
)

var (
	ErrorMapper = map[string]string{
		INVL: "has invalid format",
		UNKN: "is unknown",
		MTCh: "does not match",
		MDTR: "cannot be empty",
		ULGN: "need to login",
	}
)

func ErrorBuilder(field, err string) error {
	return fmt.Errorf("%s %s", field, ErrorMapper[err])
}
