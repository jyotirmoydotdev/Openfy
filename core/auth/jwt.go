package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(username string, isAdmin bool) (string, error) {
	var role string
	if isAdmin {
		role = "admin"
	} else {
		role = "user"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	var secretKey string
	if isAdmin {
		secretKey = adminSecrets[username]
	} else {
		secretKey = userSecrets[username]
	}
	signalToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "Internal Server error", err
	}
	return signalToken, nil
}
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(extractSecretkeyFromToken(token)), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
func extractSecretkeyFromToken(token *jwt.Token) string {
	username := extractUsernameFromToken(token)
	var secretkey string
	var isOk bool
	if isAdmin, err := extractIsAdminFromToken(token); err == nil && isAdmin {
		secretkey, isOk = adminSecrets[username]
	} else {
		secretkey, isOk = userSecrets[username]
	}
	if !isOk {
		return ""
	}
	return secretkey
}
func extractUsernameFromToken(token *jwt.Token) string {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}
	username, ok := claims["username"].(string)
	if !ok {
		return ""
	}
	return username
}
func extractIsAdminFromToken(token *jwt.Token) (bool, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("invalid token claims")
	}
	role, ok := claims["role"].(string)
	if !ok {
		return false, fmt.Errorf("invalid or missing isAdmin field in token claims")
	}
	return role == "admin", nil
}
func generateRandomKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "Something went wrong", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}