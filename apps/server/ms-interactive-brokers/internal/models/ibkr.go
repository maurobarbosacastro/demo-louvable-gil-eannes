package models

import (
	"ms-interactive-brokers/pkg/types"
)

// Define the structs to match the JSON structure
type Section struct {
	SecType  string `json:"secType"`
	Months   string `json:"months,omitempty"`
	Exchange string `json:"exchange,omitempty"`
	Conid    string `json:"conid,omitempty"`
}

type Issuer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ContractResponse struct {
	Conid         string    `json:"conid"`
	CompanyHeader string    `json:"companyHeader"`
	CompanyName   string    `json:"companyName"`
	Symbol        string    `json:"symbol"`
	Description   string    `json:"description"`
	Restricted    *bool     `json:"restricted"`
	Fop           *string   `json:"fop"`
	Opt           string    `json:"opt"`
	War           string    `json:"war"`
	Sections      []Section `json:"sections"`
	Issuers       []Issuer  `json:"issuers,omitempty"`
	BondID        int       `json:"bondid,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type MarketDataSnapshot struct {
	LastPrice              *string `json:"31"` // Last Price. The last price at which the contract traded. May contain one of the following prefixes: C - Previous day's closing price. H - Trading has halted.
	MarkerMethodDelivery   string  `json:"6119"`
	MarketDataAvailability string  `json:"6509"`
	Updated                float64 `json:"updated"`
	ConID                  int     `json:"conid"`
	ConIDEx                string  `json:"conidEx"`
	ServerID               string  `json:"server_id"`
	PriorClose             string  `json:"7741"`
}

type BrokerageSession struct {
	Authenticated bool   `json:"authenticated"`
	Competing     bool   `json:"competing"`
	Connected     bool   `json:"connected"`
	MAC           string `json:"MAC"`
	Message       string `json:"message"`
}

type Logout struct {
	Status bool `json:"status"`
}

type TickleResponseSuccess struct {
	Session    string `json:"session"`
	SsoExpires int    `json:"ssoExpires"`
	Collission bool   `json:"collission"`
	UserID     int    `json:"userId"`
	Hmds       struct {
		Error string `json:"error"`
	} `json:"hmds"`
	Iserver struct {
		Tickle     bool `json:"tickle"`
		AuthStatus struct {
			Authenticated bool   `json:"authenticated"`
			Competing     bool   `json:"competing"`
			Connected     bool   `json:"connected"`
			Message       string `json:"message"`
			MAC           string `json:"MAC"`
			ServerInfo    struct {
				ServerName    string `json:"serverName"`
				ServerVersion string `json:"serverVersion"`
			} `json:"serverInfo"`
		} `json:"authStatus"`
	} `json:"iserver"`
}

type TickleResponseError struct {
	Error string `json:"error"`
}

type AccessTokenFromCurrentSessionResponse struct {
	IsPaper          bool   `json:"is_paper"`
	OauthToken       string `json:"oauth_token"`
	OauthTokenSecret string `json:"oauth_token_secret"`
}

type Parent struct {
	AccountID   string        `json:"accountId"`
	IsMChild    bool          `json:"isMChild"`
	IsMParent   bool          `json:"isMParent"`
	IsMultiplex bool          `json:"isMultiplex"`
	MMC         []interface{} `json:"mmc"` // Empty slice, type unclear from data
}

// Account represents the main account structure
type Account struct {
	PrepaidCryptoP          bool    `json:"PrepaidCrypto-P"`
	PrepaidCryptoZ          bool    `json:"PrepaidCrypto-Z"`
	AccountAlias            *string `json:"accountAlias"` // nil in your data
	AccountID               string  `json:"accountId"`
	AccountStatus           int64   `json:"accountStatus"` // Large number, likely timestamp or status code
	AccountTitle            string  `json:"accountTitle"`
	AccountVan              string  `json:"accountVan"`
	AcctCustType            string  `json:"acctCustType"`
	BrokerageAccess         bool    `json:"brokerageAccess"`
	BusinessType            string  `json:"businessType"`
	Category                string  `json:"category"`
	ClearingStatus          string  `json:"clearingStatus"`
	Covestor                bool    `json:"covestor"`
	Currency                string  `json:"currency"`
	Desc                    string  `json:"desc"`
	DisplayName             string  `json:"displayName"`
	Faclient                bool    `json:"faclient"`
	IBEntity                string  `json:"ibEntity"`
	ID                      string  `json:"id"`
	NoClientTrading         bool    `json:"noClientTrading"`
	Parent                  Parent  `json:"parent"`
	TrackVirtualFXPortfolio bool    `json:"trackVirtualFXPortfolio"`
	TradingType             string  `json:"tradingType"`
	Type                    string  `json:"type"`
}

type IncrementRule struct {
	LowerEdge float64 `json:"lowerEdge"`
	Increment float64 `json:"increment"`
}

type DisplayRuleStep struct {
	DecimalDigits int     `json:"decimalDigits"`
	LowerEdge     float64 `json:"lowerEdge"`
	WholeDigits   int     `json:"wholeDigits"`
}

type DisplayRule struct {
	Magnification   int               `json:"magnification"`
	DisplayRuleStep []DisplayRuleStep `json:"displayRuleStep"`
}

type Position struct {
	AcctID          string          `json:"acctId"`
	Conid           int             `json:"conid"`
	ContractDesc    string          `json:"contractDesc"`
	Position        float64         `json:"position"`
	MktPrice        float64         `json:"mktPrice"`
	MktValue        float64         `json:"mktValue"`
	Currency        string          `json:"currency"`
	AvgCost         float64         `json:"avgCost"`
	AvgPrice        float64         `json:"avgPrice"`
	RealizedPnl     float64         `json:"realizedPnl"`
	UnrealizedPnl   float64         `json:"unrealizedPnl"`
	Exchs           interface{}     `json:"exchs"`
	Expiry          *string                `json:"expiry"`
	PutOrCall       *string                `json:"putOrCall"`
	Multiplier      float32                `json:"multiplier"`
	Strike          *types.FlexibleFloat64 `json:"strike"`
	ExerciseStyle   *string                `json:"exerciseStyle"`
	ConExchMap      []interface{}   `json:"conExchMap"`
	AssetClass      string          `json:"assetClass"`
	UndConid        int             `json:"undConid"`
	Model           string          `json:"model"`
	IncrementRules  []IncrementRule `json:"incrementRules"`
	DisplayRule     DisplayRule     `json:"displayRule"`
	CrossCurrency   *bool           `json:"crossCurrency,omitempty"`
	Time            int             `json:"time"`
	ChineseName     string          `json:"chineseName"`
	AllExchanges    string          `json:"allExchanges"`
	ListingExchange string          `json:"listingExchange"`
	CountryCode     string          `json:"countryCode"`
	Name            string          `json:"name"`
	LastTradingDay  *string         `json:"lastTradingDay"`
	Group           *string         `json:"group"`
	Sector          *string         `json:"sector"`
	SectorGroup     *string         `json:"sectorGroup"`
	Ticker          string          `json:"ticker"`
	Type            string          `json:"type"`
	UndComp         *string         `json:"undComp,omitempty"`
	UndSym          *string         `json:"undSym,omitempty"`
	HasOptions      bool            `json:"hasOptions"`
	FullName        string          `json:"fullName"`
	IsUS            *bool           `json:"isUS,omitempty"`
	IsEventContract bool            `json:"isEventContract"`
	PageSize        int             `json:"pageSize"`
}
