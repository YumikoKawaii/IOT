package auth

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/xerrors"
	"time"
	"yumikokawaii.iot.com/config"
)

type JWTClaim struct {
	jwt.StandardClaims
	Username string
}

type JWTResolver struct {
	config *config.JWTConfig
}

func NewJWTResolver(appConfig *config.AppConfig) JWTResolver {
	return JWTResolver{
		config: &appConfig.JwtCfg,
	}
}

func (j *JWTResolver) GenerateJWTToken(username string) (string, error) {

	claims := JWTClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.config.TokenDuration).Unix(),
		},
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.config.SecretKey))
}

func (j *JWTResolver) VerifyJWTAccessToken(accessToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, xerrors.Errorf("unexpected token signing method")
			}
			return []byte(j.config.SecretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaim); !ok {
		return nil, xerrors.Errorf("invalid token")
	} else {
		return claims, nil
	}
}
