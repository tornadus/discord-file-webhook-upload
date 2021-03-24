package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Print("Enter a webhook URL: ")
	fmt.Scan(&webhookURL)

	fmt.Print("Enter a file name: ")
	fmt.Scan(&fileName)

	bodyWriter, bodyBuffer := createFile(fileName)

	request, err := http.NewRequest("POST", webhookURL, bodyBuffer)
	handleError(err, true)

	request.Header.Add("Content-Type", bodyWriter.FormDataContentType())

	response, err := httpClient.Do(request)
	handleError(err, true)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		fmt.Println("Sent the request properly.")
		return
	}

	fmt.Println("Failed sending to the webhook")
}

func createFile(fileName string) (*multipart.Writer, *bytes.Buffer) {
	file, err := os.Open(fileName)
	handleError(err, true)
	defer file.Close()

	body := &bytes.Buffer{}
	bodywriter := multipart.NewWriter(body)

	part, err := bodywriter.CreateFormFile("multipart/form-data", filepath.Base(file.Name()))
	handleError(err, true)

	io.Copy(part, file)
	bodywriter.Close()

	return bodywriter, body
}

func handleError(err error, serious bool) {
	if err == nil {
		return
	}

	if serious {
		panic(err.Error())
	}

	fmt.Printf("[Error]: %s\n", err.Error())
}

var (
	fileName   string
	httpClient = &http.Client{}
	webhookURL string
)
