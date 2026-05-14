package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/datatypes" 
)
//---SESSIONS---
type Session struct {
    ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
    UserID       uuid.UUID `gorm:"type:uuid;not null;index"` // Bien indexado
    RefreshToken string    `gorm:"uniqueIndex;not null"`    // ¡Debe ser único e indexado!
    UserAgent    string    `gorm:"size:255"`                // Opcional: Para saber si es Chrome, iPhone, etc.
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
// --- USUARIO ---
type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	ExternalID    string    `gorm:"uniqueIndex;not null"` 
	Email         string    `gorm:"unique;not null"`
	Credits       int       `gorm:"default:0"`
	CreatedAt     time.Time
	PhotoUrl      *string        `gorm:"type:text"`
	Notifications []Notification `gorm:"foreignKey:UserID"`
}

// --- GENERACIÓN IA (REPLICATE) ---
type Sample struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID          uuid.UUID `gorm:"type:uuid;not null;index"`
	SampleName      string    `gorm:"size:100;not null"`
	Prompt          string    `gorm:"type:text"`
	InitialAudioURL string    
	// PredictionID guarda el ID de Replicate para cuando llegue el webhook
	PredictionID    string    `gorm:"uniqueIndex"` 
	Status          string    `gorm:"default:'processing'"` // processing, succeeded, failed
	CreatedAt       time.Time
}

// --- EDICIÓN Y EFECTOS ---
type SampleVersion struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey"`
	SampleID      uuid.UUID      `gorm:"type:uuid;not null;index"`
	Effects       datatypes.JSON `gorm:"type:jsonb"` 
	FinalAudioURL string    
	CreatedAt     time.Time
}

// --- SOCIAL ---
type SharedSample struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	SampleID        *uuid.UUID `gorm:"type:uuid;index"` 
	SampleVersionID *uuid.UUID `gorm:"type:uuid;index"`
	UserID          uuid.UUID `gorm:"type:uuid;not null;index"`
	Likes           int       `gorm:"default:0"`
	Downloads       int       `gorm:"default:0"`
	CreatedAt       time.Time
}

// --- NOTIFICACIONES Y WEBHOOKS ---
type Notification struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;index"`
	Type        string    `gorm:"size:50"` // 'replicate', 'payment', 'info'
	Title       string    `gorm:"size:255"`
	Message     string    `gorm:"type:text"`
	Status      string    `gorm:"size:20;default:'unread'"` // 'unread', 'read', 'failed'
	ReferenceID string    `gorm:"size:255"` // ID del sample relacionado
	CreatedAt   time.Time
}

// --- PAGOS (PADDLE) ---
type PaymentLog struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey"`
	ExternalRefID string         `gorm:"uniqueIndex"` 
	UserID        uuid.UUID      `gorm:"type:uuid;not null"`
	RawPayload    datatypes.JSON `gorm:"type:jsonb"`
	Status        string         
	CreatedAt     time.Time
}