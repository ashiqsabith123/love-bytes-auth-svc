package utils

import (
	"errors"

	"github.com/ashiqsabith123/auth-svc/pkg/config"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var SERVICE_ID string

var client *twilio.RestClient

func InitTwilio(config config.Config) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.Twilio.AccSid,
		Password: config.Twilio.Auth,
	})
	SERVICE_ID = config.Twilio.SerSid
}

func SendOtp(phone string) error {
	params := &openapi.CreateVerificationParams{}
	params.SetTo("+91" + phone)
	params.SetChannel("sms")

	_, err := client.VerifyV2.CreateVerification(SERVICE_ID, params)

	if err != nil {
		return err
	}

	return err

}

func VerifyOtp(phone string, code string) (int, error) {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(SERVICE_ID, params)

	if err != nil {
		return 403, errors.New("Something went wrong")
	} else if *resp.Status == "pending" {
		return 401, errors.New("Incorrect OTP")
	} else if *resp.Status == "canceled" {
		return 403, errors.New("OTP expired")
	}

	return 200, nil
}
