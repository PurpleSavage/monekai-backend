package authvalueobjects



type CreateUserVO struct {
	ExternalID   string
	Email        string
	PhotoUrl     *string
	RefreshToken string
	UserAgent string
}