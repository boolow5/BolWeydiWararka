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

	. "github.com/franela/goblin"
	"gopkg.in/gin-gonic/gin.v1"
)

var TOKEN string

var table []map[string]interface{}

func init() {
	table = []map[string]interface{}{
		{"method": http.MethodPost, "expected_code": http.StatusOK, "url": "/user", "body": strings.NewReader(`{"username":"lajecleey5", "password":"jiijo143"}`), "expected": map[string]interface{}{"message": "Welcome to iWeydi"}},
		{"method": http.MethodPost, "expected_code": http.StatusOK, "url": "/login", "body": strings.NewReader(`{"username":"lajecleey5", "password":"jiijo143"}`), "expected": map[string]interface{}{"expire": "", "token": ""}},
		{"method": http.MethodPut, "expected_code": http.StatusOK, "url": "/user?user_id=1", "body": strings.NewReader(`{"password":"mahdi143"}`), "expected": map[string]interface{}{"message": "Update succeeded"}},
		{"method": http.MethodDelete, "expected_code": http.StatusOK, "url": "/user?user_id=1", "expected": map[string]interface{}{"success": "Update successfully"}},
	}
}

// TestIndex checks the index or / route
func TestIndex(t *testing.T) {
	g := Goblin(t)
	gin.SetMode(gin.TestMode)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", controllers.Index)

	r.POST("/user", controllers.AddUser)
	r.PUT("/user", controllers.UpdateUser)
	r.DELETE("/user", controllers.DeleteUser)

	jwtMiddleWare := middlewares.NewJWTMiddleware()
	r.POST("/login", jwtMiddleWare.LoginHandler)

	g.Describe("Index", func() {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		g.It("Request error should be nil", func() {
			g.Assert(err).Equal(nil)
		})

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		g.It("Response code should 200", func() {
			g.Assert(w.Code).Equal(http.StatusOK)
		})

		var responseJson map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &responseJson)
		g.It("Body unmarshalling error should be nil", func() {
			g.Assert(err).Equal(nil)
		})
		expectedJson := map[string]interface{}{"message": "Welcome to iWeydi"}

		g.It("It should return the expected message", func() {
			g.Assert(responseJson["message"]).Equal(expectedJson["message"])
		})
	})

	g.Describe("AddUser", func() {
		test := table[0]

		// create request
		var body io.Reader
		if test["body"] != nil {
			body = test["body"].(io.Reader)
		}
		req, err := http.NewRequest(test["method"].(string), test["url"].(string), body)

		g.It("Request error should be nil", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Should return success 200 code", func() {
			g.Assert(rec.Code).Equal(test["expected_code"])
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body

		expectedJson := map[string]interface{}{"success": "Saved successfully"}

		g.It("Should return the expected response message", func() {
			g.Assert(responseJson["success"]).Equal(expectedJson["success"])
		})
	})

	g.Describe("Auth", func() {
		test := table[1]

		// create request
		var body io.Reader
		if test["body"] != nil {
			body = test["body"].(io.Reader)
		}
		req, err := http.NewRequest(test["method"].(string), test["url"].(string), body)

		g.It("Should have no request error", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Should return success code 200", func() {
			g.Assert(rec.Code).Equal(test["expected_code"])
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		g.It("Should return the expected response body", func() {
			g.Assert(responseJson["token"] == nil).Equal(false)
		})
		TOKEN = fmt.Sprintf("%v", responseJson["token"])
	})

	g.Describe("UpdateUser", func() {
		test := table[2]

		// create request
		var body io.Reader
		if test["body"] != nil {
			body = test["body"].(io.Reader)
		}
		req, err := http.NewRequest(test["method"].(string), test["url"].(string), body)
		g.It("Should have no request error", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Shoudl return success code 200", func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		g.It("Should return the expected response body", func() {
			g.Assert(responseJson["success"]).Equal("updated successfully")
		})
	})

	g.Describe("DeleteUser", func() {
		test := table[3]

		// create request
		var body io.Reader
		if test["body"] != nil {
			body = test["body"].(io.Reader)
		}
		req, err := http.NewRequest(test["method"].(string), test["url"].(string), body)
		g.It("Should have no request error", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Should return success code 200", func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body

		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		g.It("Should return success in response body", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
	})
}
