package audioports

import (
	audioentities "github.com/purplesvage/moneka-ai/internal/audio/domain/entities"
	audiovalueobjects "github.com/purplesvage/moneka-ai/internal/audio/domain/valueobjects"
)

type SongGeneratorPort interface {
	GenerateSong(vo *audiovalueobjects.PayloadGenerateSongVO) (*audioentities.SongGeneratorResponse, error)
}