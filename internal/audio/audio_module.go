package audio

import (
	audioinadapters "github.com/purplesvage/moneka-ai/internal/audio/in/adapters"
	audioHttp "github.com/purplesvage/moneka-ai/internal/audio/in/http"
	audiosse "github.com/purplesvage/moneka-ai/internal/audio/in/sse"
	audiooutapters "github.com/purplesvage/moneka-ai/internal/audio/out/adapters"
	audiousecases "github.com/purplesvage/moneka-ai/internal/audio/usecases"
	sharedvalidators "github.com/purplesvage/moneka-ai/internal/shared/in/validators"
	sharedadapters "github.com/purplesvage/moneka-ai/internal/shared/out/adapters"
	"github.com/purplesvage/moneka-ai/internal/shared/privatemiddlewares"
)
func Bootstrap(v *sharedvalidators.DTOValidator) (*audioHttp.AudioHandler, *audiosse.AudioSSEHandler) {

	sseManager := audioinadapters.NewSSEManager()

	songService,_:=audiooutapters.NewReplicateAdapterService()
	jwt := sharedadapters.NewJwtAdapterService()
	generateSongUC:=audiousecases.NewGeneratorSongUseCase(songService)

	// middlewares 
	authMiddleware := privatemiddlewares.NewAuthMiddleware(jwt)

	audioRESTHandler := audioHttp.NewAudioHandler(
		authMiddleware,
		generateSongUC,
		v,
		sseManager, // <-- Tienes que agregar este parámetro en el constructor de tu http.NewAudioHandler
	)

	audioSSEHandler := audiosse.NewAudioSSEHandler(sseManager,authMiddleware)
	return audioRESTHandler, audioSSEHandler
}