package dtos

import (
	errors "Payment-Gateway/pkg/error"
	"encoding/xml"
)

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    SOAPBody `xml:"Body"`
}

type SOAPBody struct {
	DepositRequest    *SOAPDepositRequest    `xml:"DepositRequest,omitempty"`
	WithdrawalRequest *SOAPWithdrawalRequest `xml:"WithdrawalRequest,omitempty"`
}

type SOAPDepositRequest struct {
	XMLName xml.Name `xml:"DepositRequest"`
	Account string   `xml:"Account"`
	Amount  float64  `xml:"Amount"`
}

func (r *SOAPDepositRequest) Validate() error {
	if r.Account == "" {
		return errors.ErrAccountRequired
	}
	if r.Amount <= 0 {
		return errors.ErrAmountMustBePositive
	}
	return nil
}

type SOAPWithdrawalRequest struct {
	XMLName xml.Name `xml:"WithdrawalRequest"`
	Account string   `xml:"Account"`
	Amount  float64  `xml:"Amount"`
}

func (r *SOAPWithdrawalRequest) Validate() error {
	if r.Account == "" {
		return errors.ErrAccountRequired
	}
	if r.Amount <= 0 {
		return errors.ErrAmountMustBePositive
	}
	return nil
}
