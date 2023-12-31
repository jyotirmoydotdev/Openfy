package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/jyotirmoydotdev/openfy/db/models"
)

var UserJWT string

// Check it a new user can signup or not
// Expected : 200
func TestUserSignup(t *testing.T) {
	newUser := map[string]string{
		"email":     "testuser@example.com",
		"password":  "testpassword",
		"firstname": "Jyotirmoy",
		"lastname":  "Barman",
	}

	jsonUser, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.Post(server.URL+"/signup", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong staus code: got %v want %v", status, http.StatusOK)
	}
}

// Check if a new user can login or not
// Expected : 200
func TestUserLogin(t *testing.T) {
	loginCredentials := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonCredentials, err := json.Marshal(loginCredentials)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.Post(server.URL+"/login", "application/json", bytes.NewBuffer(jsonCredentials))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var reponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&reponse)
	if err != nil {
		t.Errorf("Error decoding JSON response:%v", err)
		return
	}
	UserJWTfetch, ok := reponse["token"].(string)
	if !ok {
		t.Errorf("Something went wrrong while fetching token from the reponse")
	}
	UserJWT = UserJWTfetch
}
func TestUserLogin2(t *testing.T) {
	loginCredentials := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonCredentials, err := json.Marshal(loginCredentials)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.Post(server.URL+"/login", "application/json", bytes.NewBuffer(jsonCredentials))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var reponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&reponse)
	if err != nil {
		t.Errorf("Error decoding JSON response:%v", err)
		return
	}
	UserJWTfetch, ok := reponse["token"].(string)
	if !ok {
		t.Errorf("Something went wrrong while fetching token from the reponse")
	}
	UserJWT = UserJWTfetch
}

// Check is same username can signup
// Expected: 400
func TestFailSameUsername(t *testing.T) {
	newUser := models.Customer{
		Email:    "testuser@example.com",
		Password: "testpassword2",
	}
	jsonUser, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.Post(server.URL+"/signup", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong staus code: got %v want %v", status, http.StatusBadRequest)
	}
}
func TestNthUserSignup(t *testing.T) {
	testNthUser := 10
	for i := 0; i < testNthUser; i++ {
		email := strconv.Itoa(i) + "testuser@example.com"
		newUser := map[string]string{
			"email":     email,
			"password":  "testpassword",
			"firstname": strconv.Itoa(i) + "Jyotirmoy",
			"lastname":  strconv.Itoa(i) + "Barman",
		}
		jsonUser, err := json.Marshal(newUser)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.Post(server.URL+"/signup", "application/json", bytes.NewBuffer(jsonUser))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if status := resp.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong staus code: got %v want %v", status, http.StatusOK)
		}
	}
}
func TestNthUserLogin(t *testing.T) {
	testNthUser := 10
	for i := 0; i < testNthUser; i++ {
		email := strconv.Itoa(i) + "testuser@example.com"
		newUser := map[string]string{
			"email":    email,
			"password": "testpassword",
		}
		jsonUser, err := json.Marshal(newUser)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.Post(server.URL+"/login", "application/json", bytes.NewBuffer(jsonUser))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if status := resp.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong staus code: got %v want %v", status, http.StatusOK)
		}
	}
}

func TestUserPingPong(t *testing.T) {
	// Create a request with the correct endpoint
	req, err := http.NewRequest("GET", server.URL+"/user/ping", nil)
	if err != nil {
		t.Fatal("Error creating request:", err)
	}

	// Set the Authorization header with the UserJWT
	req.Header.Set("Authorization", "Bearer "+UserJWT)

	// Make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("Error making request:", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Decode the JSON response
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
		return
	}

	// Check if the "message" field exists in the response
	message, ok := response["message"].(string)
	if !ok {
		t.Error("Expected 'message' field in response, but it was not found")
	}
	if message != "pong" {
		t.Error("Message does not match")
	}
}
