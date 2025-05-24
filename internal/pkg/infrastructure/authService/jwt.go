package authservice

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// type Jwt struct {
// 	SecretKey     string
// 	TokenDuration time.Duration
// }

func TokenGenerator(id string, role string) (string, error) {

	claims := jwt.MapClaims{ //Crear el claims con la data que queremos mandar en el token
		"userID": id,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //Crear el token usando método HS256

	signedToken, err := token.SignedString([]byte("Codigo Secreto")) //Firmar el token con la secret key //os.Getenv() Para obtener data del environment
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}
