package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/italocomini/tasks/models"
	"github.com/jinzhu/gorm"
	"net/http"
)

type CreateTaskRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateTaskRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var tasks []models.Task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

const BadRequestMsg = "Record not found!"

func CreateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input CreateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{Title: input.Author, Author: input.Author}
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func FindTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": BadRequestMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": BadRequestMsg})
		return
	}

	var input UpdateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&task).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": BadRequestMsg})
		return
	}

	db.Delete(&task)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
