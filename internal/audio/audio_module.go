package audio

import (
	audioHttp "github.com/purplesvage/moneka-ai/internal/audio/in/http"
	audiooutapters "github.com/purplesvage/moneka-ai/internal/audio/out/adapters"
	audiousecases "github.com/purplesvage/moneka-ai/internal/audio/usecases"
	sharedvalidators "github.com/purplesvage/moneka-ai/internal/shared/in/validators"
	sharedadapters "github.com/purplesvage/moneka-ai/internal/shared/out/adapters"
	"github.com/purplesvage/moneka-ai/internal/shared/privatemiddlewares"
)
func Bootstrap(v *sharedvalidators.DTOValidator) *audioHttp.AudioHandler {
	songService,_:=audiooutapters.NewReplicateAdapterService()
	jwt := sharedadapters.NewJwtAdapterService()
	generateSongUC:=audiousecases.NewGeneratorSongUseCase(songService)



	// middlewares 
	authMiddleware := privatemiddlewares.NewAuthMiddleware(
		jwt,
	)

	return  audioHttp.NewAudioHandler(
		authMiddleware ,
		generateSongUC,
		v,
	)
}