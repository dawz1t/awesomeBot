package chatGPT

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Responce struct {
	Id      string
	Object  string
	Created string
	Model   string
	Choices []Choices
}

type Choices struct {
	Text          string
	Index         int
	Logprobs      string
	Finish_reason string
}

func SendQuestion(question string) ([]byte, error) {
	client := &http.Client{}
	string := `{"model": "text-davinci-003", "prompt": "` + strings.Replace(question, "\n", "", -1) + `", "temperature": 1, "max_tokens": 1000}`
	log.Printf(string)
	var data = strings.NewReader(string)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", data)
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
	}
	var errorCode error

	if resp.StatusCode == 400 {
		errorCode = errors.New("Bad Request")
	}
	return bodyText, errorCode
}
