package hmrcinterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// Authenticate takes a usename and pasword and returns an authentication token to use when submitting the returns
// Scope required for VAT retuens is "write:vat"
func Authenticate(Scope string) (authToken string) {
	config := marshalConfig()

	// Build an auth endpoint string
	authEndpoint := "https://test-api.service.hmrc.gov.uk/oauth/authorize?response_type=code&" +
		"client_id=" + config.ClientID +
		"&scope=" + Scope +
		"&redirect_uri=\"\""

	// Hit the service, using the auth endpoint
	resp, _ := http.Get(authEndpoint)

	// convert the body to a string
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	return string(bodyString)
}

// Post the deserialised VATReturn object to the configured service
func (v *VATReturn) Post(authToken string) *http.Response {

	// grab the config to get the service URI
	config := marshalConfig()

	// build an enpoint, using the vrn
	vatReturnsEndpoint := config.ServiceURI + "/organisations/vat/" + strconv.Itoa(v.Vrn) + "/returns"

	// convert the VATReturn object into a JSON string
	vatReturnsString, _ := json.Marshal(v)

	// build the request and hit the service
	client := &http.Client{}
	req, _ := http.NewRequest("POST", vatReturnsEndpoint, bytes.NewBuffer(vatReturnsString))
	req.Header.Set("Accept", "application/vnd.hmrc.1.0+json")
	req.Header.Set("Content-Type", "application/jsonn")
	req.Header.Set("Authorization", "Bearer "+authToken)
	resp, _ := client.Do(req)

	// return the response
	return resp
}

// VATReturn object deserialises to the expected JSON object for "/organisations/vat/{vrn}/returns"
type VATReturn struct {
	Vrn                          int
	PeriodKey                    string  `json:"ServiceURI"`
	VatDueSales                  float32 `json:"vatDueSales"`
	VatDueAcquisitions           float32 `json:"totalVatDue"`
	TotalVatDue                  float32 `json:"vatDueAcquisitions"`
	VatReclaimedCurrPeriod       float32 `json:"vatReclaimedCurrPeriod"`
	NetVatDue                    float32 `json:"netVatDue"`
	TotalValueSalesExVAT         float32 `json:"totalValueSalesExVAT"`
	TotalValuePurchasesExVAT     float32 `json:"totalValuePurchasesExVAT"`
	TotalValueGoodsSuppliedExVAT float32 `json:"totalValueGoodsSuppliedExVAT"`
	TotalAcquisitionsExVAT       float32 `json:"totalAcquisitionsExVAT"`
	Finalised                    bool    `json:"finalised"`
}

// Marshals the config out of the json file into a config object
func marshalConfig() config {

	// Grab the settings.json file
	configString, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Marshal it into a setting object
	var c config
	err2 := json.Unmarshal(configString, &c)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	return c
}

// struct representing the config in RAM
type config struct {
	ServiceURI string `json:"ServiceURI"`
	ClientID   string `json:"ClientID"`
}
