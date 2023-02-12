package dallE

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type GeneratedImage struct {
	Created string
	Data    []ImageData
}

type ImageData struct {
	Url string
}

func GenerateImage(question string) ([]byte, error) {
	client := &http.Client{}
	var data = strings.NewReader(`{
  	"prompt": "` + strings.Replace(question, "\n", "", -1) + `",
  	"n": 2,
  	"size": "1024x1024"
	}`)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", data)
	if err != nil {
		log.Fatal(err)
	}
	openAIkey, _ := os.LookupEnv("OPENAI_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIkey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		log.Printf("err")
	}
	var errorCode error

	if resp.StatusCode == 400 {
		errorCode = errors.New("Bad Request")
	}

	return bodyText, errorCode
}
