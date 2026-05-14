
package authdtos

type AuthRequestDto struct {
    Token string `json:"token" validate:"required,string"`
}