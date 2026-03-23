package interactive_brokers

import (
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/http_client"
	"net/http"
)

func GetLastPriceBulk(body GetLastPriceBulkDto) (*[]GetLastPriceBulkRO, error) {
	msInteractiveBrokersUrl := dotenv.GetEnv("MS_INTERACTIVE_BROKERS_URL") + "lastPrice"

	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}
	var response []GetLastPriceBulkRO

	_, err := internalHttpClient.PostJson(msInteractiveBrokersUrl, nil, &body, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
