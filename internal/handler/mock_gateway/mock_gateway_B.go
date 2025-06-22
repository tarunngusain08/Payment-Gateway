package mockgateway

import (
	"encoding/xml"
	"net/http"
)

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    SOAPBody `xml:"Body"`
}

type SOAPBody struct {
	Response SOAPResponse `xml:"Response"`
}

type SOAPResponse struct {
	Status  string `xml:"Status"`
	Message string `xml:"Message"`
}

func GatewayBMockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	resp := SOAPEnvelope{
		Body: SOAPBody{
			Response: SOAPResponse{
				Status:  "success",
				Message: "Mock Gateway B processed the request successfully",
			},
		},
	}
	xml.NewEncoder(w).Encode(resp)
}

func GatewayBMockDepositHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	resp := SOAPEnvelope{
		Body: SOAPBody{
			Response: SOAPResponse{
				Status:  "success",
				Message: "Mock Gateway B processed the deposit successfully",
			},
		},
	}
	xml.NewEncoder(w).Encode(resp)
}

func GatewayBMockWithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	resp := SOAPEnvelope{
		Body: SOAPBody{
			Response: SOAPResponse{
				Status:  "success",
				Message: "Mock Gateway B processed the withdrawal successfully",
			},
		},
	}
	xml.NewEncoder(w).Encode(resp)
}
