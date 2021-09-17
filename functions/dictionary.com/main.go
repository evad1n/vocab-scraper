package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/evad1n/vocab-scraper/define"
)

type Response struct {
	Definitions []string `json:"definitions"`
}

var (
	acceptedSources = []string{
		"dictionaryCom",
		"lexico",
		"cambridge",
	}
)

// Called like <API_ENDPOINT>?word=<WORD>

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	word := req.QueryStringParameters["word"]

	// Only 1 endpoint or gather all definitions?
	source, sourceSpecified := req.QueryStringParameters["source"]

	var defs []string
	if !sourceSpecified {

	} else if {
		defs, err := defineSource(word, source)
	}
	response := fmt.Sprintf("Defining %q...", req.QueryStringParameters["word"])

	defs, err := define.DictionaryCom.Define(word)

	defs

	returnJSONData, err := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: returnJSONData,
	}, nil
}

func defineSource(word string, source string) ([]string, error) {
	switch source {
	case "dictionaryCom":
		return define.DictionaryCom.Define(word)
	case "lexico":
		return define.Lexico.Define(word)
	case "cambridge":
		return define.Cambridge.Define(word)
	default:
		return nil, fmt.Errorf("unrecognized source: %q", source)
	}
}

func returnError(err error) events.APIGatewayProxyResponse {

}

func contains(list []string, val string) bool {
	for _, x := range list {
		if val == x {
			return true
		}
}
	return false
}

func main() {
	lambda.Start(HandleRequest)
}
