package model

type ImcState struct {
	Classification string `json:"classification"`
	ImcMin         float64 `json:"imc_min"`
	ImcMax         float64 `json:"imc_max"`
	WeightMin      float64 `json:"weight_min"`
	WeightMax      float64 `json:"weight_max"`
}

type States []ImcState

