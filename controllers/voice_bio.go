package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func VoiceBioAuth(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "file/"+file.Filename)
	transID := c.PostForm("transaction_id")
	log.Print(transID)

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "Success",
	})

}
