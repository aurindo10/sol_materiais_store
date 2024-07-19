package producthandler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/aurindo10/sol_store/internal/db/entities"
	"github.com/joho/godotenv"
)

func TestGetAllProducts(t *testing.T) {
	err := godotenv.Load()
	testToken := os.Getenv("TEST_TOKEN")
	req, _ := http.NewRequest("GET", "http://localhost:8080/get_all_products", nil)
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
	var allProducts []entities.Products
	bodyBytes, error := io.ReadAll(resp.Body)
	if error != nil {
		t.Errorf("Erro interno ao tentar ler o body")
	}
	error = json.Unmarshal(bodyBytes, &allProducts)
	if error != nil {
		t.Errorf("Erro interno ao tentar ler o body")
	}
	println(string(bodyBytes))
	defer resp.Body.Close()
	println(resp.Status)
}

func TestCreateManyProductsProduct(t *testing.T) {
	err := godotenv.Load()
	testToken := os.Getenv("TEST_TOKEN")

	products := []entities.Products{
		{
			Name:        "Product 1",
			Description: "Description 1",
			Price:       "10.99",
			Brand:       "Brand 1",
			Visible:     true,
			Amount:      100,
			Category:    "Category 1",
			Weight:      1.5,
		},
		{
			Name:        "Product 2",
			Description: "Description 2",
			Price:       "15.99",
			Brand:       "Brand 2",
			Visible:     true,
			Amount:      150,
			Category:    "Category 2",
			Weight:      2.5,
		},
	}
	productJSON, err := json.Marshal(products)
	if err != nil {
		t.Fatalf("Erro ao serializar o produto: %v", err)
	}
	req, _ := http.NewRequest("POST", "http://localhost:8080/create_products", bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")
	bearerToken := fmt.Sprintf("Bearer %s", testToken)
	req.Header.Set("Authorization", bearerToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Falha na requisição: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code esperado %v, obtido %v", http.StatusOK, resp.StatusCode)
	}
	println(resp.Status)
}
