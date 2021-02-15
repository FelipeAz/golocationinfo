package locationhelper

import (
	"fmt"

	"github.com/FelipeAz/golocationinfo/structs"
)

// GetLocationInfo returns all informations about a location using the Postal Code.
func GetLocationInfo(postcode string) {
	location, err := getAPIData(postcode)

	if err != nil {
		fmt.Println(err)
	} else if (structs.Location{}) == location {
		fmt.Printf("Nenhum resultado encontrado para o CEP %s", postcode)
	} else {
		fmt.Printf("Dados Encontrados para o CEP %s:\n", postcode)
		fmt.Printf("CEP: %s\n", location.Cep)
		fmt.Printf("Rua: %s\n", location.Street)
		fmt.Printf("Bairro: %s\n", location.Neighborhood)
		fmt.Printf("Cidade: %s\n", location.City)
		fmt.Printf("Estado: %s\n", location.State)
	}
}
