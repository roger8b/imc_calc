package handlers

import (
	"encoding/json"
	"github.com/roger8b/imc_calc/model"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

func ImcCalc(w http.ResponseWriter, r *http.Request) {
	var response model.ImcResponse
	var user model.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-type", "application/json; charset=UTF=8")
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	calculateImc(&user, &response)
	getNutritionalState(&response)

	w.Header().Set("Content-type", "application/json; charset=UTF=8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {

	}
}

func calculateImc(userData *model.User, response *model.ImcResponse) {
	weight, _ := getFloatOrError(userData.Weight)
	height, _ := getFloatOrError(userData.Height)
	response.Imc = weight / (math.Pow(height, 2))
}

func getFloatOrError(value string) (float64, error) {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return result, nil

}

func getNutritionalState(response *model.ImcResponse, ) {
	for _, state := range states {
		if response.Imc >= state.ImcMin && response.Imc <= state.ImcMax {
			response.ImcState = state

		}
	}
}

var states = model.States{
	model.ImcState{
		Classification: "Abaixo do peso",
		ImcMin:         0,
		ImcMax:         18.4,
		WeightMax:      61.3,
		WeightMin:      0,
	},
	model.ImcState{
		Classification: "Peso saudável",
		ImcMin:         18.5,
		ImcMax:         24.9,
		WeightMax:      76.6,
		WeightMin:      61.3,
	},
	model.ImcState{
		Classification: "Sobrepeso",
		ImcMin:         25.0,
		ImcMax:         29.9,
		WeightMax:      91.9,
		WeightMin:      76.6,
	},
	model.ImcState{
		Classification: "Obeso",
		ImcMin:         30.0,
		ImcMax:         39.9,
		WeightMax:      61.3,
		WeightMin:      91.9,
	},
	model.ImcState{
		Classification: "Obeso mórbido",
		ImcMin:         40.0,
		ImcMax:         9999,
		WeightMax:      122.5,
		WeightMin:      91.9,
	},
}
