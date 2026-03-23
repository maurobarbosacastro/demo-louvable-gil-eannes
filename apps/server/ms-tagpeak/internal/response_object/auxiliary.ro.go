package response_object

type AuxLogo struct {
	Id       string `json:"id"`
	Logo     string `json:"logo"`
	Original string `json:"original"`
}
type AuxProfilePicture struct {
	Id                      string `json:"id"`
	Original                string `json:"original"`
	ProfilePicture          string `json:"profilePicture"`
	ProfilePictureThumbnail string `json:"profilePictureThumbnail"`
	ProfilePictureSmall     string `json:"profilePictureSmall"`
}

type EuVatApiResponse struct {
	Success        bool   `json:"success"`
	Valid          bool   `json:"valid"`
	FormatValid    bool   `json:"formatValid"`
	Database       string `json:"database"`
	Query          string `json:"query"`
	CountryCode    string `json:"countryCode"`
	VatNumber      string `json:"vatNumber"`
	CompanyName    string `json:"companyName"`
	CompanyAddress string `json:"companyAddress"`
}

type VatValidityResponse struct {
	Success bool `json:"success"`
}
