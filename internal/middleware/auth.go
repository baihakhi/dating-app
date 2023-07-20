package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/baihakhi/dating-app/internal/models"
	"github.com/golang-jwt/jwt"
)

// CreateToken generates a JWT token for the given user.
// and last login time. If the user is not verified, it adds a "reset" claim to the token,
// The token is then signed using the HMAC-SHA256 signing method with the JWT signature key from the environment variables.
func CreateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{}

	claims["username"] = user.Username
	claims["full_name"] = user.Fullname
	claims["is_verified"] = user.IsVerified
	claims["last_login"] = user.LastLogin

	// If the user is not verified, set the "reset" claim to indicate that verification needs to be reset within 24 hours.
	if claims["is_verified"] == false {
		claims["reset"] = time.Now().Add(time.Hour * 24).Unix()
	}

	// Create a new JWT token with the specified claims and sign it using the HMAC-SHA256 signing method
	// with the JWT signature key from the environment variables.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SIGNATURE_KEY")))
}

// TokenValid validates the JWT token in the given HTTP request (r).
// It parses and verifies the token using the HMAC-SHA256 signing method
func TokenValid(r *http.Request) (*models.User, error) {
	user := new(models.User)

	// Extract the JWT token from the HTTP request.
	tokenString := ExtractToken(r)

	// Parse and verify the JWT token using the HMAC-SHA256 signing method and JWT signature key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SIGNATURE_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, err = CreateUserFromMap(claims)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}


// ExtractToken extract body token to get information
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// CreateFromMap function for convert map to user struct
func CreateUserFromMap(m map[string]interface{}) (*models.User, error) {
	data, _ := json.Marshal(m)
	var result = new(models.User)
	err := json.Unmarshal(data, &result)
	return result, err
}
