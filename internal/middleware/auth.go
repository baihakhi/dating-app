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

func CreateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{}

	claims["username"] = user.Username
	claims["full_name"] = user.Fullname
	claims["is_verified"] = user.IsVerified
	claims["last_login"] = user.LastLogin

	if claims["is_verified"] == false {
		claims["reset"] = time.Now().Add(time.Hour * 24).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SIGNATURE_KEY")))
}

func TokenValid(r *http.Request) (*models.User, error) {
	user := new(models.User)
	tokenString := ExtractToken(r)
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
