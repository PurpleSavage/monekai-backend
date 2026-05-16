package authdtos

type UserResponseDto struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	PhotoURL  *string `json:"photoUrl"`
	CreatedAt string  `json:"createdAt"`
	Credits   int     `json:"credits"`
}