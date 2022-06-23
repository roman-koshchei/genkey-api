package structs

type WeightData struct {
	FSpeed struct {
		SFB       float64
		DSFB      float64
		KeyTravel float64
		KPS       [8]float64
	}
	Dist struct {
		Lateral float64
	}
	Score struct {
		FSpeed       float64
		IndexBalance float64
		LSB          float64

		TrigramPrecision int
		LeftInwardRoll   float64
		LeftOutwardRoll  float64
		RightInwardRoll  float64
		RightOutwardRoll float64
		Alternate        float64
		Redirect         float64
		Onehand          float64
	}
}
