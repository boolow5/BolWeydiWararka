package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/boolow5/BolWeydi/controllers"

	"gopkg.in/gin-gonic/gin.v1"
)

// TestUserAdd tests user creation
func TestAddTopic(t *testing.T) {
	t.Log("TestAddTopic")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/topic", controllers.AddUser)

	// create request
	body := strings.NewReader(`{"text":"Mowduuca Koowaad"}`)
	req, err := http.NewRequest(http.MethodPost, "/topic", body)
	if err != nil {
		t.Fatalf("Could not create request: %v\n", err)
	}
	// create response recorder
	rec := httptest.NewRecorder()
	// make the call
	r.ServeHTTP(rec, req)
	// check status
	if rec.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, rec.Code)
	}
	// Unmarshal the response body
	var responseJson map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &responseJson)
	// check response body
	expected_data := `{"success": "saved successfully"}`
	if fmt.Sprintf("%v", responseJson) == expected_data {
		t.Fatalf("Expected to get value %v but instead got %v ", expected_data, responseJson)
	}
}

// TestUserAdd tests user creation
func TestUpdateTopic(t *testing.T) {
	t.Log("TestUpdateTopic")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/topic", controllers.AddUser)

	// create request
	body := strings.NewReader(`{"topic_id":1,"text":"Waa mowduuc cusub"}`)
	req, err := http.NewRequest(http.MethodPut, "/topic", body)
	if err != nil {
		t.Fatalf("Could not create request: %v\n", err)
	}
	// create response recorder
	rec := httptest.NewRecorder()
	// make the call
	r.ServeHTTP(rec, req)
	// check status
	if rec.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, rec.Code)
	}
	// Unmarshal the response body
	var responseJson map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &responseJson)
	// check response body
	expected_data := `{"success": "updated successfully"}`
	if fmt.Sprintf("%v", responseJson) == expected_data {
		t.Fatalf("Expected to get value %v but instead got %v ", expected_data, responseJson)
	}
}

// TestUserAdd tests user creation
func TestDeleteTopic(t *testing.T) {
	t.Log("TestDeleteTopic")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/topic", controllers.AddUser)

	// create request
	body := strings.NewReader(`{"topic_id":1}`)
	req, err := http.NewRequest(http.MethodDelete, "/topic", body)
	if err != nil {
		t.Fatalf("Could not create request: %v\n", err)
	}
	// create response recorder
	rec := httptest.NewRecorder()
	// make the call
	r.ServeHTTP(rec, req)
	// check status
	if rec.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, rec.Code)
	}
	// Unmarshal the response body
	var responseJson map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &responseJson)
	// check response body
	expected_data := `{"success": "saved successfully"}`
	if fmt.Sprintf("%v", responseJson) == expected_data {
		t.Fatalf("Expected to get value %v but instead got %v ", expected_data, responseJson)
	}
}
