package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Karaulkin/fio-rest-api/internal/models"
	"net/http"
)

func Enrich(name string) (models.EnrichmentData, error) {
	if name == "" {
		return models.EnrichmentData{}, errors.New("name is empty")
	}

	var result models.EnrichmentData

	if age, err := fetchAge(name); err == nil {
		result.Age = age
	}

	if gender, err := fetchGender(name); err == nil {
		result.Gender = gender
	}

	if nat, err := fetchNationality(name); err == nil {
		result.Nationality = nat
	}

	return result, nil
}

func fetchAge(name string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("agify API returned status %d", resp.StatusCode)
	}

	var data models.AgeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	return data.Age, nil
}

func fetchGender(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("genderize API returned status %d", resp.StatusCode)
	}

	var data models.GenderResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.Gender, nil
}

func fetchNationality(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("nationalize API returned status %d", resp.StatusCode)
	}

	var data models.NationalityResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if len(data.Country) > 0 {
		return data.Country[0].CountryID, nil
	}

	return "", nil
}
