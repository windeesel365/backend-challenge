package fetch

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func FetchMeatText() (string, error) {
	apiurl := os.Getenv("API_URL") //API_URL จาก .env file
	resp, err := http.Get(apiurl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, resp.Body); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
