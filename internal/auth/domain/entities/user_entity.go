package authentities
import (
	"time"
)
type UserEntity struct{
	Id string
	Email string
	PhotoUrl *string
	CreatedAt time.Time
	Credits int
}