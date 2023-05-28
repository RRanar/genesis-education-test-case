package rate

import (
	"encoding/json"
	"errors"
	"net/http"
	"io/ioutil"
)

type GetRateService interface {
	New() *RateService
	GetRate() (error, float64) 
}

type RateResponseBody struct {
	NewAmount float64		`json:"new_amount"`
	NewCurrency string		`json:"new_currency"`
	OldCurrency string		`json:"old_currency"`
	OldAmount float64		`json:"old_amount"`
}

type RateService struct {
	client *http.Client
	apiKey string
}

func New(apiKey string) *RateService {
	return &RateService{
		client: &http.Client{},
		apiKey: apiKey,
	}
}

func (rateService *RateService) GetRate() (float64, error) {
	var rateResBody RateResponseBody

	req, err := http.NewRequest("GET", "https://api.api-ninjas.com/v1/convertcurrency?have=BTC&want=UAH&amount=1", nil)

	if err != nil {
		return 0, err
	}

	req.Header.Set("X-Api-Key", rateService.apiKey)
	res, err := rateService.client.Do(req)

	if err != nil {
		return 0, err
	}

	if res.StatusCode != http.StatusOK {
		return 0, errors.New("cannot process given rate")
	}
	resBuf, err :=  ioutil.ReadAll(res.Body)

	if err != nil {
		return 0, err
	}

	res.Body.Close()

	if err := json.Unmarshal(resBuf, &rateResBody); err != nil {
		return 0, err
	}
	

	return rateResBody.NewAmount, nil
}