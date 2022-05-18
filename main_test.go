package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func callAPI(t *testing.T, method, url string, body []byte) (int, []byte) {
	t.Helper()

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if err := res.Body.Close(); err != nil {
		t.Fatal(err)
	}

	return res.StatusCode, resBody
}

// TestMainAPI expects that postgres and the service is already running
func TestMainAPI(t *testing.T) {
	var accountId uuid.UUID

	t.Run("Create Account", func(t *testing.T) {
		t.Run("GET /account/create 405", func(t *testing.T) {
			status, _ := callAPI(t, "GET", "http://localhost:8080/account/create", nil)
			if status != 405 {
				t.Errorf("expected status 405 but got %d", status)
			}
		})

		t.Run("POST /account/create 201", func(t *testing.T) {
			status, body := callAPI(t, "POST", "http://localhost:8080/account/create", nil)
			if status != 201 {
				t.Errorf("expected status 201 but got %d", status)
			}

			account := struct {
				Id uuid.UUID `json:"id"`
			}{}
			if err := json.Unmarshal(body, &account); err != nil {
				t.Fatal(err)
			}
			if account.Id.String() == (uuid.UUID{}).String() {
				t.Fatal(`response body doesn't contain "id" field'`)
			}

			accountId = account.Id
		})
	})

	t.Run("Account Balance", func(t *testing.T) {
		t.Run("GET /account/balance 400", func(t *testing.T) {
			status, _ := callAPI(t, "GET", "http://localhost:8080/account/balance", nil)
			if status != 400 {
				t.Errorf("expected status 400 but got %d", status)
			}
		})

		t.Run("GET /account/balance 404", func(t *testing.T) {
			status, _ := callAPI(t, "GET", "http://localhost:8080/account/balance?id=29e75493-1b5f-4bea-a356-7207d1527645", nil)
			if status != 404 {
				t.Errorf("expected status 404 but got %d", status)
			}
		})

		t.Run("GET /account/balance 200", func(t *testing.T) {
			status, body := callAPI(t, "GET", "http://localhost:8080/account/balance?id="+accountId.String(), nil)
			if status != 200 {
				t.Errorf("expected status 200 but got %d", status)
			}

			if string(body) != `{"balance":0}` {
				t.Errorf("expected empty balance but got: %s", body)
			}
		})
	})

	t.Run("Deposit Amount", func(t *testing.T) {
		t.Run("PUT /account/balance 405", func(t *testing.T) {
			status, _ := callAPI(t, "PUT", "http://localhost:8080/account/deposit", nil)
			if status != 405 {
				t.Errorf("expected status 405 but got %d", status)
			}
		})

		t.Run("POST /account/balance 422", func(t *testing.T) {
			body := []byte(fmt.Sprintf(`{"id":"%s","amount":0}`, accountId))
			status, _ := callAPI(t, "POST", "http://localhost:8080/account/deposit", body)
			if status != 422 {
				t.Errorf("expected status 422 but got %d", status)
			}
		})

		t.Run("POST /account/balance 204", func(t *testing.T) {
			body := []byte(fmt.Sprintf(`{"id":"%s","amount":5}`, accountId))
			status, _ := callAPI(t, "POST", "http://localhost:8080/account/deposit", body)
			if status != 204 {
				t.Errorf("expected status 204 but got %d", status)
			}

			// verify deposit
			status, body = callAPI(t, "GET", "http://localhost:8080/account/balance?id="+accountId.String(), nil)
			if status != 200 {
				t.Errorf("expected status 200 but got %d", status)
			}

			if string(body) != `{"balance":5}` {
				t.Errorf("expected balance=5 but got: %s", body)
			}
		})
	})
}
