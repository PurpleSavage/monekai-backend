package authports

// arreglar esto porque el dominio no debería depender de una librería externa
import (
	"github.com/golang-jwt/jwt/v5"
)
type JwtPort interface{
	GenerateToken(email string, durationStr string ) (string, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
}