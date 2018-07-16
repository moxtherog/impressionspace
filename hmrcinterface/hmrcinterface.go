package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var v VATRequest
	PostVATRequest(v)
}

// PostVATRequest takes a VATRequests object, generates a VAT request xml document and posts it to the configured service
func PostVATRequest(v VATRequest) {

	// grab the config to get the service URI
	c := marshalConfig()

	// make a request object
	vatRequest := makeVATRequestXML(v)

	// hit the service #TODO error handling!
	resp, _ := http.Post(c.ServiceURI, "application/xml", bytes.NewBuffer(vatRequest))

	// print the response
	fmt.Println(resp)
}

// converts a VAT request object into an XML doc, using templates/vat-request-template.xml
func makeVATRequestXML(VATRequest) []byte {

	// Marshal the template into a byte array
	vatRequest, err := ioutil.ReadFile("templates/vat-request-template.xml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// #TODO do stuff to it based on the VATReqest object

	return vatRequest
}

// VATRequest object holds the attributes required for generating a vat request xml doc
type VATRequest struct {
	SenderID               string  // xpath: /GovTalkMessage/Header/SenderDetails/IDAuthentication/SenderID
	AuthenticationValue    string  // xpath: /GovTalkMessage/Header/SenderDetails/IDAuthentication/Authentication/Value
	VATRegNo               int     // xpath: /GovTalkMessage/GovTalkDetails/Keys/Key /IRheader/Keys/Key #TODO need to put this in twice...
	ChannelURI             string  // xpath: /GovTalkMessage/GovTalkDetails/ChannelRouting/Channel/URI
	ChannelProduct         string  // xpath: /GovTalkMessage/GovTalkDetails/ChannelRouting/Channel/Product
	CahnnelVersion         string  // xpath: /GovTalkMessage/GovTalkDetails/ChannelRouting/Channel/Version
	PeriodID               string  // xpath: /IRheader/PeriodID
	IRMark                 string  // xpath: /IRheader/IRmark
	Sender                 string  // xpath: /IRheader/Sender
	VATDueOnOutputs        float32 // xpath: /VATDeclarationRequest/VATDueOnOutputs
	VATDueOnECAcquisitions float32 // xpath: /VATDeclarationRequest/VATDueOnECAcquisitions
	TotalVAT               float32 // xpath: /VATDeclarationRequest/TotalVAT
	VATReclaimedOnInputs   float32 // xpath: /VATDeclarationRequest/VATReclaimedOnInputs
	NetVAT                 float32 // xpath: /VATDeclarationRequest/NetVAT
	NetSalesAndOutputs     int     // xpath: /VATDeclarationRequest/NetSalesAndOutputs
	NetPurchasesAndInputs  int     // xpath: /VATDeclarationRequest/NetPurchasesAndInputs
	NetECSupplies          int     // xpath: /VATDeclarationRequest/NetECSupplies
	NetECAcquisitions      int     // xpath: /VATDeclarationRequest/NetECAcquisitions
	AASBalancingPayment    float32 // xpath: /VATDeclarationRequest/AASBalancingPayment
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
}
