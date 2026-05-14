package authentities


type SessionEntity struct {
	UserData UserEntity
	AccessToken string
	RefreshToken string
}