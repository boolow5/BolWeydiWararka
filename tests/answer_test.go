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
func TestAddAnswer(t *testing.T) {
	t.Log("TestAddAnswer")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/answer", controllers.AddUser)

	// create request
	body := strings.NewReader(`{"text":"Waa jawaab cusub", "question":{"question_id":"1"}}`)
	req, err := http.NewRequest(http.MethodPost, "/answer", body)
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
func TestUpdateAnswer(t *testing.T) {
	t.Log("TestUpdateAnswer")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/answer", controllers.AddUser)

	// create request
	body := strings.NewReader(`{"answer_id":1,"text":"Waa jawaab la cusbooneysiiyay cusub"}`)
	req, err := http.NewRequest(http.MethodPut, "/answer", body)
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
func TestDeleteAnswer(t *testing.T) {
	t.Log("TestDeleteAnswer")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/answer", controllers.AddUser)

	// create request
	body := strings.NewReader(`{"answer_id":1}`)
	req, err := http.NewRequest(http.MethodDelete, "/answer", body)
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
