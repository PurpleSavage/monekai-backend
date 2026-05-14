package sharedports

type JwtPort interface{
	GenerateToken(email string, durationStr string ) (string, error)
	VerifyToken(tokenString string) (string, error)
}