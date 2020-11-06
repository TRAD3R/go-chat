package helpers

import (
	"encoding/json"
	"net/http"
)

type Avatar struct {
	Url    string `json:"message"`
	Status string `json:"status"`
}

func GetRandomAvatar() (string, error) {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	var avatar *Avatar
	if err := json.NewDecoder(response.Body).Decode(&avatar); err != nil {
		return "", err
	}

	return avatar.Url, nil
}
