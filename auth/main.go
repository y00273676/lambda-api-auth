package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type event struct {
	MethodArn          string `json:"methodArn"`
	AuthorizationToken string `json:"authorizationToken"`
}

type context struct {
	StringKey  string `json:"stringKey"`
	NumberKey  int    `json:"numberKey"`
	BooleanKey bool   `json:"booleanKey"`
}
type policyDocument struct {
	Version   string      `json:"Version"`
	Statement []statement `json:"Statement"`
}
type AuthResponse struct {
	Context        context        `json:"context"`
	PrincipalId    string         `json:"principalId"`
	PolicyDocument policyDocument `json:"policyDocument"`
}

type statement struct {
	Action   string `json:"Action"`
	Effect   string `json:"Effect"`
	Resource string `json:"Resource"`
}

func Handler(e event) (AuthResponse, error) {
	log.Printf("events = %+v", e)
	s := statement{
		Action:   "execute-api:Invoke",
		Effect:   "Allow",
		Resource: e.MethodArn,
	}
	slist := make([]statement, 0)
	slist = append(slist, s)
	return AuthResponse{
		PrincipalId: "user",
		Context: context{
			StringKey:  "stringval",
			NumberKey:  123,
			BooleanKey: true,
		},
		PolicyDocument: policyDocument{
			"2012-10-17",
			slist,
		},
	}, nil
}

func main() {
	lambda.Start(Handler)

}
