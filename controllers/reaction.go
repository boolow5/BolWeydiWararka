package controllers

import (
	"strconv"

	"github.com/boolow5/BolWeydi/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {

}

// AddReaction creates new reaction
func AddReaction(context *gin.Context) {
	reaction := models.Reaction{}
	context.BindJSON(&reaction)
	saved, err := reaction.Add()
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

func UpdateReaction(context *gin.Context) {
	id, got := context.GetQuery("reaction_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing reaction id"})
		return
	}
	reactionId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if reactionId < 1 {
		context.JSON(200, gin.H{"error": "invalid reaction id"})
		return
	}

	reaction := &models.Reaction{ReactionId: reactionId}
	reaction, err = models.GetReactionById(reaction.ReactionId)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	err = context.BindJSON(reaction)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	saved, err := reaction.Update()
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

func DeleteReaction(context *gin.Context) {
	id, got := context.GetQuery("reaction_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing reaction id"})
		return
	}
	reactionId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if reactionId < 1 {
		context.JSON(200, gin.H{"error": "invalid reaction id"})
		return
	}

	reaction := &models.Reaction{ReactionId: reactionId}
	reaction, err = models.GetReactionById(reaction.ReactionId)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	deleted, err := reaction.Delete()
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
