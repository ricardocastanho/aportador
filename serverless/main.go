package main

import (
	"aportador/principles"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

type Response events.APIGatewayProxyResponse

type CustomResponse struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func getFloatParams(param string, defaultValue float64) (float64, error) {
	if param == "" {
		return defaultValue, nil
	}

	parsed, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return parsed, err
	}

	return parsed, nil
}

func handleError(logger zerolog.Logger, msg string, err error, statusCode int) (Response, error) {
	resp, _ := json.Marshal(CustomResponse{
		Data: nil,
		Error: map[string]interface{}{
			"message": msg,
		},
	})
	logger.Err(err).Msg(msg)
	return Response{Body: string(resp), StatusCode: statusCode}, err
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	mainLogger := logger.
		With().
		Str("requestId", request.RequestContext.RequestID).
		Str("stocks", request.QueryStringParameters["stock"]).
		// .Str("fiis", request.QueryStringParameters["fii"]).
		Str("dividend-yield", request.QueryStringParameters["dividendYield"]).
		Str("dividend-history", request.QueryStringParameters["dividendHistory"]).
		Logger()

	mainLogger.Debug().Msg("New request received")

	stocks := strings.Split(request.QueryStringParameters["stock"], ",")
	// fiis := strings.Split(request.QueryStringParameters["fii"], ",")

	if len(stocks) == 1 && stocks[0] == "" {
		mainLogger.Error().Msg("No tickers provided")
		return handleError(mainLogger, "No tickers provided", fmt.Errorf("no tickers provided"), 200)
	}

	dividendYield, err := getFloatParams(request.QueryStringParameters["dividendYield"], 10.0)
	if err != nil {
		return handleError(mainLogger, err.Error(), err, 500)
	}

	dividendHistory, err := getFloatParams(request.QueryStringParameters["dividendHistory"], 3.0)
	if err != nil {
		return handleError(mainLogger, err.Error(), err, 500)
	}

	mainLogger.
		Debug().
		Float64("dividendYield", dividendYield).
		Float64("dividendHistory", dividendHistory).
		Msg("Search params")

	results, err := principles.GetStocks(stocks, dividendYield, dividendHistory)
	if err != nil {
		return handleError(mainLogger, "Error while searching", err, 500)
	}

	body, err := json.Marshal(CustomResponse{
		Data:  results,
		Error: nil,
	})
	if err != nil {
		return handleError(mainLogger, err.Error(), err, 500)
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	mainLogger.Info().RawJSON("result", body).Msg("Request completed")

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
