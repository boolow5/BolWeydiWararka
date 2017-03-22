package controllers

import (
	"strconv"

	"github.com/boolow5/BolWeydi/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {

}

// AddQuestion creates new question
func AddQuestion(context *gin.Context) {
	question := models.Question{}
	context.BindJSON(&question)
	saved, err := question.Add()
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if !saved {
		context.JSON(200, gin.H{"warning": "failed without warning"})
		return
	}
	context.JSON(200, gin.H{"success": "Saved successfully"})
}

func UpdateQuestion(context *gin.Context) {
	id, got := context.GetQuery("question_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing question id"})
		return
	}
	questionId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if questionId < 1 {
		context.JSON(200, gin.H{"error": "invalid question id"})
		return
	}

	question := &models.Question{QuestionId: questionId}
	question, err = models.GetQuestionById(question)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	err = context.BindJSON(question)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	saved, err := question.Add()
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if !saved {
		context.JSON(200, gin.H{"warning": "update failed without error"})
		return
	}
	context.JSON(200, gin.H{"success": "updated successfully"})
}

func DeleteQuestion(context *gin.Context) {
	id, got := context.GetQuery("question_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing question id"})
		return
	}
	questionId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if questionId < 1 {
		context.JSON(200, gin.H{"error": "invalid question id"})
		return
	}

	question := &models.Question{QuestionId: questionId}
	question, err = models.GetQuestionById(question)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	deleted, err := question.Delete()
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		context.JSON(200, gin.H{"warning": "update failed without error"})
		return
	}
	context.JSON(200, gin.H{"success": "deleted successfully"})
}
