package authadapters

import (
	"errors"
	"github.com/google/uuid"
	authentities "github.com/purplesvage/moneka-ai/internal/auth/domain/entities"
	authports "github.com/purplesvage/moneka-ai/internal/auth/domain/ports"
	authvalueobjects "github.com/purplesvage/moneka-ai/internal/auth/domain/valueobjects"
	domainerrors "github.com/purplesvage/moneka-ai/internal/shared/domain/errors"
	"github.com/purplesvage/moneka-ai/persistence"
	"gorm.io/gorm"
)
 
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) authports.UserPersistencePort{
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(data authvalueobjects.CreateUserVO) (*authentities.UserEntity,error){
	userID := uuid.New()
    sessionID := uuid.New()
	newUser := models.User{
        ID:         userID,
        ExternalID: data.ExternalID,
        Email:      data.Email,
        PhotoUrl:   data.PhotoUrl,
        Credits:    0, // Valor inicial por defecto
    }
	newSession := models.Session{
        ID:           sessionID,
        UserID:       userID, // Vinculamos la sesión al usuario recién creado
        RefreshToken: data.RefreshToken,
        UserAgent: data.UserAgent,
    }
	err := r.db.Transaction(func(tx *gorm.DB) error {
        // 1. Crear el usuario
        if err := tx.Create(&newUser).Error; err != nil {
            return err // Si falla, GORM hace rollback automático
        }
        // 2. Crear la sesión
        if err := tx.Create(&newSession).Error; err != nil {
            return err // Si falla, se deshace la creación del usuario
        }

        return nil
    })

    if err != nil {
        return nil, domainerrors.NewAppError(
            500, 
            "Error registering user", 
            "The registration in the database could not be completed", 
            err,
        )
    }
	return &authentities.UserEntity{
        Id:        newUser.ID.String(),
        Email:     newUser.Email,
        PhotoUrl:  newUser.PhotoUrl,
        CreatedAt: newUser.CreatedAt,
        Credits:   newUser.Credits,
    }, nil
}

func (r *UserRepository) FindUserByEmail(email string)(*authentities.UserEntity,error){
	var userModel  models.User 
	err := r.db.Where("email = ?", email).First(&userModel).Error
	if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
			// El usuario no existe
			return nil, domainerrors.NewAppError(404, "User not found", "There is no user with that email", err)
		}
		// Error de la bd (conexión, permisos, etc.)
		return nil, domainerrors.NewAppError(500, "Error de base de datos", "Hubo un fallo al conectar con el servidor", err)
    }
	return &authentities.UserEntity{
        Id:        userModel.ID.String(),
        Email:     userModel.Email,
        CreatedAt: userModel.CreatedAt,
        Credits:   userModel.Credits,
        PhotoUrl:  userModel.PhotoUrl, 
    }, nil
}

func (r *UserRepository) UpdateSession(token string, userId string) error{
    result := r.db.Model(&models.Session{}).
        Where("user_id = ?", userId).
        Update("refresh_token", token)
    if result.Error != nil {
        return domainerrors.NewAppError(500, "Database Error", "Failed to update session in database", result.Error)
    }
    if result.RowsAffected == 0 {
        // Si no hubo filas afectadas, es porque el user_id no tenía una sesión previa
        return domainerrors.NewAppError(404, "Session Not Found", "No active session found for the given user", nil)
    }
    return  nil
}