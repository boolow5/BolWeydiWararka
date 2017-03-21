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
func TestAddQuestion(t *testing.T) {
	t.Log("TestAddQuestion")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/question", controllers.AddUser)

	// create request
	body := strings.NewReader(`{"text":"Wax iweydi sxb?", "author":{"user_id":"1"}}`)
	req, err := http.NewRequest(http.MethodPost, "/question", body)
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
func TestUpdateQuestion(t *testing.T) {
	t.Log("TestUpdateQuestion")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/question", controllers.AddUser)

	// create request
	body := strings.NewReader(`{"question_id":1,"text":"Su'aasha si kale makuu fahansiiyaa?"}`)
	req, err := http.NewRequest(http.MethodPut, "/question", body)
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
func TestDeleteQuestion(t *testing.T) {
	t.Log("TestDeleteQuestion")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/question", controllers.AddUser)

	// create request
	body := strings.NewReader(`{"question_id":1}`)
	req, err := http.NewRequest(http.MethodDelete, "/question", body)
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
