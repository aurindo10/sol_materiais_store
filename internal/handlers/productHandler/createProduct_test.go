package producthandler_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateProduct(t *testing.T) {
	err := godotenv.Load()
	testToken := os.Getenv("TEST_TOKEN")
	req, _ := http.NewRequest("GET", "http://localhost:8080/all_products", nil)
	req.Header.Set("Content-Type", "application/json")
	bearerToken := fmt.Sprintf("Bearer %s", testToken)
	req.Header.Set("Authorization", bearerToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Falha na requisição: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code esperado %v, obtido %v", http.StatusOK, resp.StatusCode)
	}
	println(resp.Status)
}
