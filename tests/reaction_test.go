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
func TestReaction(t *testing.T) {
	gin.SetMode(gin.TestMode)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/reaction", controllers.AddReaction)
	r.PUT("/reaction", controllers.UpdateReaction)
	r.DELETE("/reaction", controllers.DeleteReaction)

	g := Goblin(t)

	g.Describe("AddReaction", func() {
		// create request
		body := strings.NewReader(`{"positive":true, "user": {"user_id":1}, "question": {"question_id": 1}}`)
		req, err := http.NewRequest(http.MethodPost, "/reaction", body)

		g.It("Should not have request error", func() {
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
			g.Assert(responseJson["error"]).Equal(nil)
		})
		g.It("Should return success", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should return the expected message", func() {
			g.Assert(responseJson["success"]).Equal(expected_data)
		})
	})

	g.Describe("UpdateReaction", func() {
		// create request
		body := strings.NewReader(`{"positive":true, "user": {"user_id":1}, "question": {"question_id": 1}}`)
		req, err := http.NewRequest(http.MethodPut, "/reaction?reaction_id=1", body)
		g.It("Should not have request error", func() {
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

	g.Describe("DeleteReaction", func() {
		// create request
		body := strings.NewReader(``)
		req, err := http.NewRequest(http.MethodDelete, "/reaction?reaction_id=1", body)
		g.It("Should not have request error", func() {
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
