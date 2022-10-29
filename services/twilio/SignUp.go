package twilio

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	Account_Sid string
	Auth_Token  string
	Service_Sid string
)

type Do struct{}

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file loading error - ", err)
		fmt.Println("error here", err.Error())
		return err
	}
	Account_Sid = os.Getenv("TWILIO_ACCOUNT_SID")
	Auth_Token = os.Getenv("TWILIO_AUTH_TOKEN")
	Service_Sid = os.Getenv("SIGNUP_SERVICE_SID")
	return nil
}

func (t Do) SendOtp(To string) error {
	if err := LoadEnv(); err != nil {
		log.Fatal("Env Load Err")
		return err
	}
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: Account_Sid,
		Password: Auth_Token,
	})

	params := &openapi.CreateVerificationParams{}
	params.SetTo(To)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(Service_Sid, params)

	if err != nil {
		return err
	}
	fmt.Printf("Verification has been send, Id :'%s'\n", *resp.Sid)
	return nil
}
func (t Do) CheckOtp(To, code string) (bool, error) {
	if err := LoadEnv(); err != nil {
		log.Fatal("Env Load Err")
		return false, err
	}
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: Account_Sid,
		Password: Auth_Token,
	})
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(To)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(Service_Sid, params)

	if err != nil {
		fmt.Println("error :", err.Error())
		return false, err
	}
	if *resp.Status == "approved" {
		return true, nil
	}
	return false, nil
}
