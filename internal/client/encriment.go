package client

import (
	"encoding/json"
	"fmt"
	"github.com/Karaulkin/fio-rest-api/internal/models"
	"net/http"
)

// Enrich насыщает данными
func Enrich(name string) (models.EnrichmentData, error) {
	var result models.EnrichmentData

	var err error
	var resp *http.Response

	// Получаем возраст
	if resp, err = http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name)); err == nil {
		defer resp.Body.Close()
		var data models.AgeResponse
		json.NewDecoder(resp.Body).Decode(&data)
		result.Age = data.Age
	}

	if err != nil {
		return models.EnrichmentData{}, err
	}

	// Получаем пол
	if resp, err = http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name)); err == nil {
		defer resp.Body.Close()
		var data models.GenderResponse
		json.NewDecoder(resp.Body).Decode(&data)
		result.Gender = data.Gender
	}

	if err != nil {
		return models.EnrichmentData{}, err
	}

	// Получаем национальность
	if resp, err = http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name)); err == nil {
		defer resp.Body.Close()
		var data models.NationalityResponse
		json.NewDecoder(resp.Body).Decode(&data)
		if len(data.Country) > 0 {
			result.Nationality = data.Country[0].CountryID
		}
	}

	if err != nil {
		return models.EnrichmentData{}, err
	}

	return result, nil
}
