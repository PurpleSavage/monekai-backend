package audiousecases

import (
	audioentities "github.com/purplesvage/moneka-ai/internal/audio/domain/entities"
	audioports "github.com/purplesvage/moneka-ai/internal/audio/domain/ports"
	audiovalueobjects "github.com/purplesvage/moneka-ai/internal/audio/domain/valueobjects"
	audioincommands "github.com/purplesvage/moneka-ai/internal/audio/in/dtos/commands"
	audioutils "github.com/purplesvage/moneka-ai/internal/audio/utils"
)

type GenerateSongUseCase struct{
	songGeneratorService audioports.SongGeneratorPort
}

func NewGeneratorSongUseCase(
	songGeneratorService audioports.SongGeneratorPort,
)*GenerateSongUseCase{
	return  &GenerateSongUseCase{
		songGeneratorService:songGeneratorService,
	}
}

func (g *GenerateSongUseCase) Execute(
	dto audioincommands.GenerateSongDTO, 
	email string,
)(*audioentities.SongGeneratorResponse, error){
	vo, appErr := audiovalueobjects.CreatePayloadGenerateSongVO(
		audioutils.BuildWebhook("songs"),
		dto.ModelVersion,
		dto.Prompt,
		dto.Duration,
		dto.OutputFormat,
	)
	if appErr != nil{
		return  nil, appErr
	}
	response, err:= g.songGeneratorService.GenerateSong(vo)
	if err != nil {
		return  nil,err
	}
	return  response,nil
}