package audioentities

import audioenums "github.com/purplesvage/moneka-ai/internal/audio/domain/enums"

type SongGeneratorResponse struct {
	GenerationId     string
	StatusGeneration audioenums.Status
}