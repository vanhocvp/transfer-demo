package util

import (
	"github.com/dambaquyen96/smartivr-backend-go/pkg/setting"
	"github.com/dgrijalva/jwt-go"
)

func newJWTFromClaims(claims jwt.MapClaims) (string, error) {
	// add expired time
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := at.SignedString([]byte(setting.APISetting.JWTRoleTokenSecretKey))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func GenerateNewAPIKey(APIKeyID int) (string, error) {
	mapClaims := make(jwt.MapClaims)
	mapClaims["id"] = APIKeyID
	token, err := newJWTFromClaims(mapClaims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GenerateRoleKey(roleString string, isAdmin bool) (string, error) {
	mapClaims := make(jwt.MapClaims)
	mapClaims["roles_string"] = roleString
	mapClaims["is_admin"] = isAdmin
	mapClaims["create_at"] = GetCurrentTimeByMillisecond()
	token, err := newJWTFromClaims(mapClaims)
	if err != nil {
		return "", err
	}
	return token, nil
}
