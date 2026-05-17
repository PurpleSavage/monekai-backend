package audioincommands


import audioenums "github.com/purplesvage/moneka-ai/internal/audio/domain/enums"

type GenerateSongDTO struct {

	// User prompt
	Prompt string `json:"prompt" validate:"required,min=5,max=500"`

	// AI model version
	ModelVersion audioenums.MelodyVersion `json:"model_version" validate:"required"`

	// Duration in seconds
	Duration int `json:"duration" validate:"required,min=1,max=30"`

	// Final audio format
	OutputFormat audioenums.OutputFormat `json:"output_format" validate:"required"`
}