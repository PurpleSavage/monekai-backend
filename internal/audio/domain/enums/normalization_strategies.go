package audioenums

type NormalizationStrategy string

const(
	Peak 		NormalizationStrategy = "peak"
	Rms  		NormalizationStrategy = "rms"
	Clip 		NormalizationStrategy = "clip"
	Loudness 	NormalizationStrategy = "loudness"
)