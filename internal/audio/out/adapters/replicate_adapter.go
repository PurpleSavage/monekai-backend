package audiooutapters

import (
	"context"
	"github.com/purplesvage/moneka-ai/cmd/config"
	audioentities "github.com/purplesvage/moneka-ai/internal/audio/domain/entities"
	audioenums "github.com/purplesvage/moneka-ai/internal/audio/domain/enums"
	audioports "github.com/purplesvage/moneka-ai/internal/audio/domain/ports"
	audiovalueobjects "github.com/purplesvage/moneka-ai/internal/audio/domain/valueobjects"
	"github.com/replicate/replicate-go"
)

const SONG_MODEL string="meta/musicgen:671ac645ce5e552cc63a54a2bbff63fcf798043055d2dac5fc9e36a837eedcfb"

type ReplicateAdapterService struct{
	client *replicate.Client
}

func NewReplicateAdapterService() (audioports.SongGeneratorPort, error) {
	client, err := replicate.NewClient(
		replicate.WithToken(config.Envs.ReplicateKey),
	)

	if err != nil {
		return nil, err
	}

	return &ReplicateAdapterService{
		client: client,
	}, nil
}
func (re *ReplicateAdapterService) GenerateSong(vo *audiovalueobjects.PayloadGenerateSongVO)(*audioentities.SongGeneratorResponse, error){
	ctx := context.TODO()
	input := replicate.PredictionInput{
		"prompt": vo.Prompt,
		"model_version":vo.ModelVersion,
		"duration":vo.Duration,
		"normalization_strategy":vo.NormalizationStrategy,
		"classifier_free_guidance":vo.ClassifierFreeGuidance,
		"output_format":vo.OutputFormat,
		"top_k":vo.TopK,
		"top_p":vo.TopP,
		"Temperature":vo.Temperature,
	}
	webhook := replicate.Webhook{
		URL:    vo.WebhookUrl,
		Events: []replicate.WebhookEventType{"start", "completed"},
	}
	prediction, err := re.client.CreatePrediction(
		ctx,
		SONG_MODEL,
		input,
		&webhook,
		false,
	)
	if err != nil {
		return nil, err
	}
	response := &audioentities.SongGeneratorResponse{
		GenerationId:     prediction.ID,
		StatusGeneration: audioenums.Status(prediction.Status),
	}

	return response, nil

}



