package controllers

import (
	"fmt"
	"log"
	"net/http"
	"bytes"
	"encoding/json"
	"os"
	"github.com/joho/godotenv"
)

// Endpoint API OpenAI untuk mendapatkan fakta sepeda
const openAIEndpoint = "https://api.openai.com/v1/engines/gpt-3.5/completions"

type OpenAIRequest struct {
	Prompt string `json:"prompt"`
	MaxTokens int `json:"max_tokens"`
}

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func getBikeFactHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		http.Error(w, "Failed to read the .env file", http.StatusInternalServerError)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	bikeType := r.URL.Query().Get("bikeType")
	if bikeType == "" {
		http.Error(w, "Missing 'bikeType' query parameter", http.StatusBadRequest)
		return
	}

	// Validasi jenis sepeda yang dimasukkan
	if bikeType != "BMX" && bikeType != "Sepeda Lipat" && bikeType != "Sepeda Gunung" {
		http.Error(w, "Jenis sepeda tidak valid. Pilih dari BMX, Sepeda Lipat, atau Sepeda Gunung.", http.StatusBadRequest)
		return
	}

	fact, err := chatAI.GetBikeFacts(bikeType, apiKey)
	if err != nil {
		http.Error(w, "Failed to retrieve bike facts", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Fakta tentang sepeda %s:\n%s", bikeType, fact)
}

func GetBikeFacts(bikeType string, apiKey string) (string, error) {
	prompt := fmt.Sprintf("Beri saya fakta unik tentang sepeda %s.", bikeType)
	requestData := OpenAIRequest{
		Prompt:    prompt,
		MaxTokens: 50, // Jumlah token maksimum dalam respons
	}

	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(requestJSON))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var response OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Choices[0].Text, nil
}
