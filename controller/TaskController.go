package controller

import (
	"net/http"
	"time"
	"todo/database"
	"todo/models"

	"github.com/gin-gonic/gin"
)

const (
	PENDING_STATUS = "pending"
	INPROGRESS_STATUS = "in-progess"
	COMPLETED_STATUS = "completed"
)

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	task.Status = PENDING_STATUS
	task.Duration = 0
	task.CreatedAt = time.Now().UTC();
	task.UpdatedAt = time.Now().UTC();

	result := database.DB.Create(&task)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task added successfully"})
}

func CreateTasksBulk(c *gin.Context) {
	var tasks []models.Task

	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&tasks)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tasks created successfully"})
	
}

func GetTaskByID(c *gin.Context) {
	var task models.Task

	if err := database.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusFound, gin.H{"data": task})
}

func GetAllTasks(c *gin.Context) {
	var tasks []models.Task

	if err := database.DB.Find(&tasks).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func GetTasksByStatus(c *gin.Context) {
	var tasks []models.Task

	if err := database.DB.Find(&tasks, "status = ?",c.Param("status")).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusFound, gin.H{"data": tasks})
}

func GetTasksByCategory(c *gin.Context) {
	var tasks []models.Task

	if err := database.DB.Find(&tasks, "category = ?",c.Param("category")).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusFound, gin.H{"data": tasks})
}

func UpdateTask(c *gin.Context) {
	var updateBody struct {
		Description string `json:"description"`
		Category string `json:"category"`
	}	

	if err := c.ShouldBindJSON(&updateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	var task models.Task

	if err := database.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return 
	}

	if len(updateBody.Category) != 0 {
		task.Category = updateBody.Category
	}

	if len(updateBody.Description) != 0 {
		task.Description = updateBody.Description
	}

	task.UpdatedAt = time.Now().UTC()

	database.DB.Save(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {
	var task models.Task

	if err := database.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

func CompleteTask(c *gin.Context) {
	var requestBody struct {
		Duration int `json:"duration"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	var task models.Task

	if err := database.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return 
	}

	task.Status = INPROGRESS_STATUS
	task.Duration += requestBody.Duration

	task.UpdatedAt = time.Now().UTC()

	database.DB.Save(&task)

	c.JSON(http.StatusOK, gin.H{"messaage": "progress updated successfully"})
}

func MarkTaskCompleted(c *gin.Context) {
	var task models.Task

	if err := database.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return 
	}

	task.Status = COMPLETED_STATUS
	task.UpdatedAt = time.Now().UTC()

	database.DB.Save(&task)

	c.JSON(http.StatusOK, gin.H{"message": "task completed successfully"})
}