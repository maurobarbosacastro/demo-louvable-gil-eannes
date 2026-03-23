package dto

type CjTransactionDTO struct {
	PublisherId string `json:"publisherId"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	ApiToken    string `json:"apiToken"`
}

type GraphQLRequest struct {
	Query string `json:"query"`
}

type AdIdDTO struct {
	PublisherId  string `json:"publisherId"`
	AdvertiserId string `json:"advertiserId"`
	ApiToken     string `json:"apiToken"`
}
