package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Account struct {
	Id                  int64    `json:"id"`
	Uuid                string   `json:"uuid"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	LogoImg             string   `json:"logoimg"`
	Type                string   `json:"accounttype"`
	AddressId           int64    `json:"addressid"`
	Address1            string   `json:"address1"`
	Address2            string   `json:"address2"`
	Address3            string   `json:"address3"`
	Suburb              string   `json:"suburb"`
	PostalCode          string   `json:"postalcode"`
	City                string   `json:"city"`
	Country             string   `json:"country"`
	Language            string   `json:"language"`
	Timezone            string   `json:"timezone"`
	CallbackURI         string   `json:"callbackurl"`
	Website             string   `json:"website"`
	Currencies          []string `json:"currencies"`
	WalletAddress       string   `json:"walletaddress"`
	WalletCurrency      string   `json:"walletcurrency"`
	ContactId           int64    `json:"contactid"`
	ContactTitle        string   `json:"contacttitle"`
	ContactFirstname    string   `json:"contactfirstname"`
	ContactMiddlenames  string   `json:"contactmiddlenames"`
	ContactLastname     string   `json:"contactlastname"`
	ContactEmail        string   `json:"contactemail"`
	ContactPhone        string   `json:"contactphone"`
	ContactMobile       string   `json:"contactmobile"`
	VatNo               string   `json:"vatno"`
	DefaultVAT          int64    `json:"defaultvat"`
	Organisation        string   `json:"organisation"`
	PluginType          string   `json:"plugintype"`
	Status              string   `json:"status"`
	ResponseCode        string   `json:"responsecode"`
	ResponseDescription string   `json:"responsedescription"`
}

type AccountNew struct {
	Name                string `json:"name"`
	Description         string `json:"description"`
	Type                string `json:"accounttype"`
	Country             string `json:"country"`
	Language            string `json:"language"`
	Timezone            string `json:"timezone"`
	WalletAddress       string `json:"walletaddress"`
	WalletCurrency      string `json:"walletcurrency"`
	ContactFirstname    string `json:"contactfirstname"`
	ContactLastname     string `json:"contactlastname"`
	ContactEmail        string `json:"contactemail"`
	Status              string `json:"status"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountPersonal struct {
	Name                string `json:"name"`
	Description         string `json:"description"`
	Type                string `json:"accounttype"`
	LogoImg             string `json:"logoimg"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountPayment struct {
	CallbackURI         string   `json:"callbackurl"`
	Website             string   `json:"website"`
	Currencies          []string `json:"currencies"`
	WalletAddress       string   `json:"walletaddress"`
	WalletCurrency      string   `json:"walletcurrency"`
	PluginType          string   `json:"plugintype"`
	ResponseCode        string   `json:"responsecode"`
	ResponseDescription string   `json:"responsedescription"`
}

type AccountOrg struct {
	VatNo               string `json:"vatno"`
	DefaultVAT          int64  `json:"defaultvat"`
	Organisation        string `json:"organisation"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountAPIKey struct {
	Id                  int64  `json:"id"`
	Uuid                string `json:"uuid"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	ApiKey              string `json:"apikey"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountAddress struct {
	AccountId           int64  `json:"accountid"`
	AddressId           int64  `json:"addressid"`
	DefaultAddress      string `json:"defaultaddress"`
	AddressType         string `json:"addresstype"`
	Address1            string `json:"address1"`
	Address2            string `json:"address2"`
	Address3            string `json:"address3"`
	Suburb              string `json:"suburb"`
	PostalCode          string `json:"postalcode"`
	City                string `json:"city"`
	Country             string `json:"country"`
	Language            string `json:"language"`
	Timezone            string `json:"timezone"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountContact struct {
	AccountId           int64  `json:"accountid"`
	ContactId           int64  `json:"contactid"`
	ContactType         string `json:"contacttype"`
	ContactTitle        string `json:"contacttitle"`
	ContactFirstname    string `json:"contactfirstname"`
	ContactMiddlenames  string `json:"contactmiddlenames"`
	ContactLastname     string `json:"contactlastname"`
	ContactEmail        string `json:"contactemail"`
	ContactPhone        string `json:"contactphone"`
	ContactMobile       string `json:"contactmobile"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountWallet struct {
	WalletAddress       string          `json:"walletaddress"`
	WalletCurrency      string          `json:"walletcurrency"`
	WalletBalance       decimal.Decimal `json:"walletbalance"`
	WalletFiat          decimal.Decimal `json:"walletfiat"`
	ECAPrice            decimal.Decimal `json:"ecaprice"`
	BTCPrice            decimal.Decimal `json:"btcprice"`
	USDPrice            decimal.Decimal `json:"usdprice"`
	ResponseCode        string          `json:"responsecode"`
	ResponseDescription string          `json:"responsedescription"`
}

type AccountActivity struct {
	ActivityDate time.Time `json:"activitydate"`
	Category     string    `json:"category"`
	Description  string    `json:"description"`
	UserName     string    `json:"username"`
}

type RuleParameter struct {
	Parameter           string   `json:"parameter"`
	Description         string   `json:"description"`
	Type                string   `json:"type"`
	Validation          string   `json:"validation"`
	Validation_Comment  string   `json:"validation_comment"`
	Options             []string `json:"options"`
	Value               string   `json:"value"`
	ResponseCode        string   `json:"responsecode"`
	ResponseDescription string   `json:"responsedescription"`
}

type RuleCompulsory struct {
	Value     bool   `json:"value"`
	Condition string `json:"condition"`
}

type AccountRule struct {
	Code           string          `json:"code"`
	Display        string          `json:"display"`
	Description    string          `json:"description"`
	Compulsory     RuleCompulsory  `json:"compulsory"`
	PostValidation string          `json:"postvalidation"`
	Parameters     []RuleParameter `json:"parameters"`
}
