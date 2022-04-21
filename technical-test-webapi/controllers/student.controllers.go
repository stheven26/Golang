package controllers

import (
	// "fmt"
	"net/http"
	"technical-test-webapi/models"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator"
)

type StudentInput struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

func RegisterStudent(c *gin.Context) {
	//validasi inputan
	var dataInput StudentInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	prose input data
	student := models.Student{
		ID:   dataInput.ID,
		Name: dataInput.Name,
		Age:  dataInput.Age,
	}

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sukses input data Student",
		"Data":    student,
		"time":    time.Now(),
	})
}

//GET Data
func GetStudent(c *gin.Context) {
	var student []StudentInput

	c.JSON(http.StatusOK, gin.H{
		"data": student,
		"time": time.Now(),
	})
}

//UPDATE Data >> Update Data
func UpdateStudent(c *gin.Context) {
	//cek data
	var student []StudentInput
	// if err := db.Where("nim = ?", c.Param("nim")).First(&student).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Data Student tidak di temukan",
	// 	})
	// 	return
	// }

	//validasi inputan
	var dataInput StudentInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sukses mengubah data Student",
		"Data":    student,
		"time":    time.Now(),
	})
}

func DeleteStudent(c *gin.Context) {
	//cek data
	var student models.Student
	// if err := db.Where("nim = ?", c.Query("nim")).First(&mhs).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Data Student tidak di temukan",
	// 	})
	// 	return
	// }

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data": student,
	})
}
