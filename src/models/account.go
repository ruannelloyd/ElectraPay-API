package models

type Account struct {
	Id                  int64    `json:"id"`
	Uuid                string   `json:"uuid"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	Type                string   `json:"accounttype"`
	LogoURL             string   `json:"logourl"`
	LogoImg             string   `json:"logoimg"`
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
	ContactTitle        string   `json:"contacttitle"`
	ContactFirstname    string   `json:"contactfirstname"`
	ContactMiddlenames  string   `json:"contactmiddlenames"`
	ContactLastname     string   `json:"contactlastname"`
	ContactEmail        string   `json:"contactemail"`
	ContactPhone        string   `json:"contactphone"`
	ContactMobile       string   `json:"contactmobile"`
	VatNo               string   `json:"vatno"`
	DefaultVAT          int64    `json:"defaultvat"`
	Organisation        string   `json:"orgnisation"`
	PluginType          string   `json:"plugintype"`
	Status              string   `json:"status"`
	ResponseCode        string   `json:"responsecode"`
	ResponseDescription string   `json:"responsedescription"`
	Secret              string   `json:"-"`
}