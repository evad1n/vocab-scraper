package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/evad1n/vocab-scraper/define"
)

type Result struct {
	Source      string   `json:"source"`
	Definitions []string `json:"definitions"`
}

type InvalidSourceResponse struct {
	Sources []string `json:"acceptedSources"`
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

	var results []Result
	if !sourceSpecified {
		for _, source := range acceptedSources {

		}
	} else {
		defs, err := defineSource(word, source)
		if err != nil {
			return error500(fmt.Errorf("finding definitions for %q: %v", source, err))
		}

		returnJSONData, err := json.Marshal(defs)
		if err != nil {
			return error500(fmt.Errorf("marshalling json for defs: %v", err))
		}

		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: string(returnJSONData),
		}, nil
	}

	response := fmt.Sprintf("Defining %q...", req.QueryStringParameters["word"])

	defs, err := define.DictionaryCom.Define(word)

	// defs

	returnJSONData, err := json.Marshal(r)

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

func invalidSource() {
	s := "accepted sources: \n"
	for _, source := range acceptedSources {
		s += fmt.Sprintf("%s,\n", source)
	}
}

func error500(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       err.Error(),
	}, nil
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
