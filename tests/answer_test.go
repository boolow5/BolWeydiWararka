package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/boolow5/BolWeydi/controllers"

	. "github.com/franela/goblin"
	"gopkg.in/gin-gonic/gin.v1"
)

// TestUserAdd tests user creation
func TestAddAnswer(t *testing.T) {
	g := Goblin(t)
	gin.SetMode(gin.TestMode)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/answer", controllers.AddAnswer)
	r.PUT("/answer", controllers.UpdateAnswer)
	r.DELETE("/answer", controllers.DeleteAnswer)

	g.Describe("TestAddAnswer", func() {
		// create request
		body := strings.NewReader(`{"text":"Waa jawaab cusub", "question":{"question_id":1}, "author":{"user_id":1}}`)
		req, err := http.NewRequest(http.MethodPost, "/answer", body)
		g.It("Should not give error while sending request", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Should return 200 success code", func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		expected_data := "Saved successfully"
		g.It("Should not return error", func() {
			g.Assert(responseJson["error"] == nil).Equal(true)
		})
		g.It("Should return success", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should return the expected message", func() {
			g.Assert(responseJson["success"]).Equal(expected_data)
		})
	})

	g.Describe("UpdateAnswer", func() {
		// create request
		body := strings.NewReader(`{"text":"Waa jawaab la cusbooneysiiyay cusub"}`)
		req, err := http.NewRequest(http.MethodPut, "/answer?answer_id=1", body)
		g.It("Should not give error while sending request", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Should return success 200 code", func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		expected_data := "updated successfully"
		g.It("Should not return error", func() {
			g.Assert(responseJson["error"] == nil).Equal(true)
		})
		g.It("Should return success", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should return the expected message", func() {
			g.Assert(responseJson["success"]).Equal(expected_data)
		})
	})

	g.Describe("DeleteAnswer", func() {
		// create request
		req, err := http.NewRequest(http.MethodDelete, "/answer?answer_id=1", nil)
		g.It("Should not give error while sending request", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Should return success 200 code", func() {
			g.Assert(rec.Code).Equal(200)
		})

		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		expected_data := "deleted successfully"
		g.It("Should not return error", func() {
			g.Assert(responseJson["error"] == nil).Equal(true)
		})
		g.It("Should return success", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should return the expected message", func() {
			g.Assert(responseJson["success"]).Equal(expected_data)
		})
	})
}
