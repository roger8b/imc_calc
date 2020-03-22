package model

type ImcResponse struct {
	Imc            float64 `json:"imc"`
	ImcState       ImcState
}
