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
func TestQuestion(t *testing.T) {
	g := Goblin(t)

	gin.SetMode(gin.TestMode)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/question", controllers.AddQuestion)
	r.PUT("/question", controllers.UpdateQuestion)
	r.DELETE("/question", controllers.DeleteQuestion)

	g.Describe("AddQuestion", func() {
		// create request
		body := strings.NewReader(`{"text":"Wax iweydi sxb?", "author":{"user_id":1}}`)
		req, err := http.NewRequest(http.MethodPost, "/question", body)
		g.It("Should have no errors when sending request", func() {
			g.Assert(err).Equal(nil)
		})
		req.Header.Add("Authorization", "Bearer "+TOKEN)
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)

		// check status
		g.It("Should have 200 status code", func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		g.It("Should have success in response", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should return the correct message", func() {
			g.Assert(responseJson["success"]).Equal("Saved successfully")
		})
	})

	g.Describe("UpdateQuestion", func() {
		// create request
		body := strings.NewReader(`{"question_id":1,"text":"Su'aasha si kale makuu fahansiiyaa?"}`)
		req, err := http.NewRequest(http.MethodPut, "/question?question_id=1", body)
		g.It("Should not have request problem", func() {
			g.Assert(err).Equal(nil)
		})
		req.Header.Add("Authorization", "Bearer "+TOKEN)
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)

		// check status
		g.It("Should return 200 status code", func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body

		g.It("Should not return warning in response body", func() {
			g.Assert(responseJson["warning"] == nil).Equal(true)
		})
		g.It("Should not return error in response body", func() {
			g.Assert(responseJson["error"] == nil).Equal(true)
		})
		g.It("Should return success in response body", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should mention success", func() {
			g.Assert(responseJson["success"]).Equal("updated successfully")
		})
	})

	g.Describe("DeleteQuestion", func() {
		// create request
		body := strings.NewReader(`{"question_id":1}`)
		req, err := http.NewRequest(http.MethodDelete, "/question?question_id=1", body)
		g.It("Should not have request problem", func() {
			g.Assert(err).Equal(nil)
		})
		req.Header.Add("Authorization", "Bearer "+TOKEN)
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Should return 200 status code", func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		g.It("Should have no error in the response", func() {
			g.Assert(responseJson["error"] == nil).Equal(true)
		})
		g.It("Should have success in the response", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should mention success message", func() {
			g.Assert(responseJson["success"]).Equal("deleted successfully")
		})
	})
}
