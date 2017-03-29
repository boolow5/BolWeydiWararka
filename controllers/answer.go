package controllers

import (
	"strconv"

	"github.com/boolow5/BolWeydi/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {

}

// AddAnswer creates new answer
func AddAnswer(context *gin.Context) {
	answer := models.Answer{}
	context.BindJSON(&answer)
	saved, err := answer.Add()
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

func UpdateAnswer(context *gin.Context) {
	id, got := context.GetQuery("answer_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing answer id"})
		return
	}
	answerId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if answerId < 1 {
		context.JSON(200, gin.H{"error": "invalid answer id"})
		return
	}

	answer := &models.Answer{AnswerId: answerId}
	answer, err = models.GetAnswerById(answer.AnswerId)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	err = context.BindJSON(answer)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	saved, err := answer.Update()
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

func DeleteAnswer(context *gin.Context) {
	id, got := context.GetQuery("answer_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing answer id"})
		return
	}
	answerId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if answerId < 1 {
		context.JSON(200, gin.H{"error": "invalid answer id"})
		return
	}

	answer := &models.Answer{AnswerId: answerId}
	answer, err = models.GetAnswerById(answer.AnswerId)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	deleted, err := answer.Delete()
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
