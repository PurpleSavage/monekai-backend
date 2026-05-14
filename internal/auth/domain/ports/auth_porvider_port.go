package authports

type ResponseProvider struct{
	ClerkId string
	Email string
	PhotoUrl *string
	PhoneNumber *string
}

type AuthProviderPort interface{
	VerifyAndExtract(token string) (*ResponseProvider,error)
}