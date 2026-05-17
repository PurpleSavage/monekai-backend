package audioenums

type Status string

const (
	Starting   Status = "starting"
	Processing Status = "processing"
	Succeeded  Status = "succeeded"
	Failed     Status = "failed"
	Canceled   Status = "canceled"
)