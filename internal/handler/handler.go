package handler

import (
	"encoding/json"
	"fmt"
	"github/felipecardosodeoliveira/golang-cloud-run/internal/models"
	"github/felipecardosodeoliveira/golang-cloud-run/internal/services"
	"net/http"
	"regexp"
	"strings"
)

func isValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}

func CepHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cep := strings.TrimPrefix(r.URL.Path, "/cep/")
	if !isValidCEP(cep) {
		http.Error(w, `{"message": "invalid zipcode"}`, http.StatusUnprocessableEntity)
		return
	}
	fmt.Println(cep)
	location, err := services.GetLocationByCEP(cep)
	if err != nil {
		http.Error(w, `{"message": "can not found zipcode"}`, http.StatusNotFound)
		return
	}
	weather, err := services.GetWeatherByCity(location.City)
	if err != nil {
		http.Error(w, `{"message": "could not retrieve weather"}`, http.StatusInternalServerError)
		return
	}
	tempF := weather.Current.TempC*1.8 + 32
	tempK := weather.Current.TempC + 273.15
	response := models.TemperatureResponse{
		TempC: weather.Current.TempC,
		TempF: tempF,
		TempK: tempK,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
