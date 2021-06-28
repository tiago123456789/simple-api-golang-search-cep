package repositories

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"search-cep/src/models"
)

func GetCep(cep string) (models.ResponseViaCep, error) {
	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return models.ResponseViaCep{}, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return models.ResponseViaCep{}, err
	}

	dataInJson := models.ResponseViaCep{}
	json.Unmarshal(data, &dataInJson)
	return dataInJson, err
}
