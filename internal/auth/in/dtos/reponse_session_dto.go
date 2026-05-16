package authdtos


type ResponseSessionDto struct {
	UserData    UserResponseDto `json:"userData"`
	AccessToken string          `json:"accessToken"`
}