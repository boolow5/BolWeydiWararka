package controllers

import (
	"strconv"

	"github.com/boolow5/BolWeydi/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {

}

// AddTopic creates new topic
func AddTopic(context *gin.Context) {
	topic := models.Topic{}
	context.BindJSON(&topic)
	saved, err := topic.Add()
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

func UpdateTopic(context *gin.Context) {
	id, got := context.GetQuery("topic_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing topic id"})
		return
	}
	topicId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if topicId < 1 {
		context.JSON(200, gin.H{"error": "invalid topic id"})
		return
	}

	topic := &models.Topic{TopicId: topicId}
	topic, err = models.GetTopicById(topic.TopicId)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	err = context.BindJSON(topic)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	saved, err := topic.Update()
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

func DeleteTopic(context *gin.Context) {
	id, got := context.GetQuery("topic_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing topic id"})
		return
	}
	topicId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if topicId < 1 {
		context.JSON(200, gin.H{"error": "invalid topic id"})
		return
	}

	topic := &models.Topic{TopicId: topicId}
	topic, err = models.GetTopicById(topic.TopicId)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	deleted, err := topic.Delete()
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
