package locationhelper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FelipeAz/golocationinfo/structs"
)

const httpURL = "http://viacep.com.br/ws/"

func requestURL(postcode string) string {
	return httpURL + postcode + "/json/"
}

func getAPIData(postcode string) (locationData structs.Location, err error) {
	locationData = structs.Location{}

	if len(postcode) != 8 && len(postcode) != 9 {
		return locationData, fmt.Errorf("Invalid Format for CEP. Use 00000-000 or 00000000")
	}

	res, err := http.Get(requestURL(postcode))
	if err != nil {
		return locationData, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&locationData)

	return
}
