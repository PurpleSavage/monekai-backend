package audioutils

import (
	"fmt"

	"github.com/purplesvage/moneka-ai/cmd/config"
)

func BuildWebhook(path string) string {
	baseUrl := config.Envs.BackendServerBaseUrl
	webhookUrl:=fmt.Sprintf(
		"%s/webhook/%s",
		baseUrl,
		path,
	)
	return  webhookUrl
}