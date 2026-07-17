package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/yogendraatluri990/go-fem_crypto_masters/data"
)

const apiURL = "https://trade.cex.io/api/spot/rest-public/get_ticker"

func GetRateViaPost(currency string) (*data.Rate, error) {
	postBody := map[string][]string{"pairs": {currency}}

	jsonData, err := json.Marshal(postBody)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("jsonData %s\n", string(jsonData))
	// httpRequestHandler, error := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	// if error != nil {
	// 	return nil, error
	// }

	// httpRequestHandler.Header.Set("Content-Type", "application/json")

	// client := &http.Client{
	// 	Timeout: time.Second * 15,
	// }

	// response, error := client.Do(httpRequestHandler)

	response, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		body := string(bodyBytes)
		fmt.Printf("The response body is %s\n", body)
	} else {
		fmt.Printf("response %v", response)
		return nil, fmt.Errorf("status received was: %v", response.StatusCode)
	}
	rate := &data.Rate{Currency: currency, Price: 20.0}
	return rate, nil
}

// curl -X POST  \
//   -H "Content-Type: application/json" \
//   -d '{"pairs":["BTC-USD"]}'
