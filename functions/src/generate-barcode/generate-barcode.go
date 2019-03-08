package main

import (
	"encoding/base64"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/qr"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	codeType := request.QueryStringParameters["codeType"]
	value := request.QueryStringParameters["value"]

	log.Printf("%+v\n", request)
	log.Printf("%+v\n", request.QueryStringParameters)

	log.Printf("codeType: %+v\n", codeType)
	log.Printf("value: %+v\n", value)

	var imageBytes []byte
	if codeType == "ean" {
		imageBytes = generateEan(value)
	} else if codeType == "code128" {
		imageBytes = generateCode128(value)
	} else if codeType == "qr" {
		imageBytes = generateQr(value)
	} else {
		// TODO 未対応のコード種別のレスポンスはこれでいいか？
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Unsupported Code Type",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Body:            base64.StdEncoding.EncodeToString(imageBytes),
		IsBase64Encoded: true,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}

func generateEan(code string) []byte {
	// Create the barcode
	eanCode, _ := ean.Encode(code)

	// Scale the barcode to 200(width)x200(height) pixels
	// eanCode, _ = barcode.Scale(eanCode., 200, 40)

	// create the output file
	file, _ := os.Create("code.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, eanCode)

	bytes, _ := ioutil.ReadAll(file)
	return ([]byte)(bytes)
}

func generateCode128(content string) []byte {
	// Create the barcode
	code128, _ := code128.Encode(content)

	// Scale the barcode to 200(width)x200(height) pixels

	// create the output file
	file, _ := os.Create("code.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, code128)

	bytes, _ := ioutil.ReadAll(file)
	return bytes
}

func generateQr(content string) []byte {
	// Create the barcode
	qrCode, _ := qr.Encode(content, qr.M, qr.Auto)

	// Scale the barcode to 200(width)x200(height) pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 40)

	// create the output file
	file, _ := os.Create("code.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)

	bytes, _ := ioutil.ReadAll(file)
	return bytes
}
