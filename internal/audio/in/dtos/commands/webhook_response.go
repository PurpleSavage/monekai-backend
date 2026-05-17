package audioincommands


type BaseMetrics struct {
	PredictTime float32 `json:"predict_time"`
	TotalTime   float32 `json:"total_time"`
}

type BaseUrls struct {
	Get    string  `json:"get"`
	Cancel string  `json:"cancel"`
	Stream *string `json:"stream,omitempty"`
}

type WebhookSongGeneratorResponse[T any] struct {
	CompletedAt string  `json:"completed_at"`
	CreatedAt   string  `json:"created_at"`
	DataRemoved bool    `json:"data_removed"`
	Error       *string `json:"error"`

	ID string `json:"id"`

	Input T `json:"input"`

	Logs string `json:"logs"`

	Metrics BaseMetrics `json:"metrics"`

	Output string `json:"output"`

	StartedAt string `json:"started_at"`
	Status    string `json:"status"`
	Version   string `json:"version"`

	Urls BaseUrls `json:"urls"`
}