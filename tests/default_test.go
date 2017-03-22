package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/boolow5/BolWeydi/controllers"
	"github.com/boolow5/BolWeydi/middlewares"

	"gopkg.in/gin-gonic/gin.v1"
)

var table []map[string]interface{}

func init() {
	table = []map[string]interface{}{
		{"method": http.MethodPost, "expected_code": http.StatusOK, "url": "/user", "body": strings.NewReader(`{"username":"lajecleey5", "password":"jiijo143"}`), "expected": map[string]interface{}{"message": "Welcome to iWeydi"}},
		{"method": http.MethodPost, "expected_code": http.StatusOK, "url": "/login", "body": strings.NewReader(`{"username":"lajecleey5", "password":"jiijo143"}`), "expected": map[string]interface{}{"expire": "", "token": ""}},
		{"method": http.MethodPut, "expected_code": http.StatusOK, "url": "/user", "body": strings.NewReader(`{"password":"mahdi143"}`), "expected": map[string]interface{}{"message": "Update succeeded"}},
		{"method": http.MethodDelete, "expected_code": http.StatusOK, "url": "/user", "expected": map[string]interface{}{"success": "Update successfully"}},
	}
}

// TestIndex checks the index or / route
func TestIndex(t *testing.T) {
	t.Log("TestIndex")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/", controllers.Index)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v\n", err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	var responseJson map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &responseJson)
	if err != nil {
		t.Fatalf("Unmarshaling Response Body Failed, GOT ERROR: %v", err)
	}
	expectedJson := map[string]interface{}{"message": "Welcome to iWeydi"}
	if responseJson["message"] != expectedJson["message"] {
		t.Fatalf("Expected to get json data %v but instead got %v\n", expectedJson["message"], responseJson["message"])
	}
}

// TestUserAdd tests user creation
func TestAddUser(t *testing.T) {
	t.Log("TestAddUser")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/user", controllers.AddUser)
	test := table[0]

	// create request
	var body io.Reader
	if test["body"] != nil {
		body = test["body"].(io.Reader)
	}
	req, err := http.NewRequest(test["method"].(string), test["url"].(string), body)
	if err != nil {
		t.Fatalf("Could not create request: %v\n", err)
	}
	// create response recorder
	rec := httptest.NewRecorder()
	// make the call
	r.ServeHTTP(rec, req)
	// check status
	if rec.Code != test["expected_code"] {
		t.Fatalf("Expected to get status %d but instead got %d\n", test["expected_code"], rec.Code)
	}
	// Unmarshal the response body
	var responseJson map[string]interface{}
	fmt.Println(rec.Body.String())
	json.Unmarshal(rec.Body.Bytes(), &responseJson)
	// check response body
	if fmt.Sprintf("%v", responseJson) == fmt.Sprintf("%v", test["expected"]) {
		t.Fatalf("Expected to get value %v but instead got %v ", test["expected"], responseJson)
	}
}

// TestAuth test login endpoint
func TestAuth(t *testing.T) {
	t.Log("TestAuth")
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	jwtMiddleWare := middlewares.NewJWTMiddleware()

	r.POST("/login", jwtMiddleWare.LoginHandler)

	test := table[1]

	// create request
	var body io.Reader
	if test["body"] != nil {
		body = test["body"].(io.Reader)
	}
	req, err := http.NewRequest(test["method"].(string), test["url"].(string), body)
	if err != nil {
		t.Fatalf("Could not create request: %v\n", err)
	}
	// create response recorder
	rec := httptest.NewRecorder()
	// make the call
	r.ServeHTTP(rec, req)
	// check status
	if rec.Code != test["expected_code"] {
		t.Fatalf("Expected to get status %d but instead got %d\n", test["expected_code"], rec.Code)
	}
	// Unmarshal the response body
	fmt.Println(rec.Body.String())
	var responseJson map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &responseJson)
	// check response body
	if fmt.Sprintf("%v", responseJson) == fmt.Sprintf("%v", test["expected"]) {
		t.Fatalf("Expected to get value %v but instead got %v ", test["expected"], responseJson)
	}
}

// TestUserAdd tests user creation // updating, deletion and authentcation
func TestUpdateUser(t *testing.T) {
	t.Log("TestUpdateUser")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/user", controllers.UpdateUser)
	test := table[2]

	// create request
	var body io.Reader
	if test["body"] != nil {
		body = test["body"].(io.Reader)
	}
	req, err := http.NewRequest(test["method"].(string), test["url"].(string), body)
	if err != nil {
		t.Fatalf("Could not create request: %v\n", err)
	}
	// create response recorder
	rec := httptest.NewRecorder()
	// make the call
	r.ServeHTTP(rec, req)
	// check status
	if rec.Code != test["expected_code"] {
		t.Fatalf("Expected to get status %d but instead got %d\n", test["expected_code"], rec.Code)
	}
	// Unmarshal the response body
	fmt.Println(rec.Body.String())
	var responseJson map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &responseJson)
	// check response body
	if fmt.Sprintf("%v", responseJson) == fmt.Sprintf("%v", test["expected"]) {
		t.Fatalf("Expected to get value %v but instead got %v ", test["expected"], responseJson)
	}
}

// TestDeleteUser test user removal
func TestDeleteUser(t *testing.T) {
	t.Log("TestDeleteUser")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/user", controllers.DeleteUser)
	test := table[3]

	// create request
	var body io.Reader
	if test["body"] != nil {
		body = test["body"].(io.Reader)
	}
	req, err := http.NewRequest(test["method"].(string), test["url"].(string), body)
	if err != nil {
		t.Fatalf("Could not create request: %v\n", err)
	}
	// create response recorder
	rec := httptest.NewRecorder()
	// make the call
	r.ServeHTTP(rec, req)
	// check status
	if rec.Code != test["expected_code"] {
		t.Fatalf("Expected to get status %d but instead got %d\n", test["expected_code"], rec.Code)
	}
	// Unmarshal the response body
	fmt.Println(rec.Body.String())
	var responseJson map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &responseJson)
	// check response body
	if fmt.Sprintf("%v", responseJson) == fmt.Sprintf("%v", test["expected"]) {
		t.Fatalf("Expected to get value %v but instead got %v ", test["expected"], responseJson)
	}
}
