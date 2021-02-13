package locationhelper

import (
	"encoding/json"
	"net/http"
)

const httpURL = "http://viacep.com.br/ws/"

type location struct {
	Cep          string `json:"cep"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
}

func requestURL(postcode string) string {
	return httpURL + postcode + "/json/"
}

func getAPIData(postcode string) (locationData location, err error) {
	locationData = location{}

	res, err := http.Get(requestURL(postcode))
	if err != nil {
		return locationData, err
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&locationData)

	return
}
