package controllers

import (
	"strconv"

	"github.com/boolow5/BolWeydi/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {

}

// AddDiscussion creates new discussion
func AddDiscussion(context *gin.Context) {
	discussion := models.Discussion{}
	context.BindJSON(&discussion)
	saved, err := discussion.Add()
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

func UpdateDiscussion(context *gin.Context) {
	id, got := context.GetQuery("discussion_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing discussion id"})
		return
	}
	discussionId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if discussionId < 1 {
		context.JSON(200, gin.H{"error": "invalid discussion id"})
		return
	}

	discussion := &models.Discussion{DiscussionId: discussionId}
	discussion, err = models.GetDiscussionById(discussion.DiscussionId)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	err = context.BindJSON(discussion)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	saved, err := discussion.Update()
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

func DeleteDiscussion(context *gin.Context) {
	id, got := context.GetQuery("discussion_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing discussion id"})
		return
	}
	discussionId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if discussionId < 1 {
		context.JSON(200, gin.H{"error": "invalid discussion id"})
		return
	}

	discussion := &models.Discussion{DiscussionId: discussionId}
	discussion, err = models.GetDiscussionById(discussion.DiscussionId)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	deleted, err := discussion.Delete()
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
