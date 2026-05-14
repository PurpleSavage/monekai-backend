package privatemiddlewares


type AuthMiddlewareStrategy interface{
	Auth()
}

type AuthMiddlewareContext struct{
	Strategy AuthMiddlewareStrategy
}
func (a *AuthMiddlewareContext) SetStartegy(authStrategy AuthMiddlewareStrategy){
	a.Strategy= authStrategy
}
func (a *AuthMiddlewareContext) Process(){
	a.Strategy.Auth()
}