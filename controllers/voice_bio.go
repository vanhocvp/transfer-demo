package controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func VoiceBioAuth(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	transID := c.PostForm("transaction_id")
	log.Print(transID)

	// Upload the file to specific dst.
	//c.SaveUploadedFile(file, "file/"+file.Filename)

	fileData, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer fileData.Close()

	// Ghi dữ liệu blob vào file mới
	fileBlob, err := ioutil.ReadAll(fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = ioutil.WriteFile("file/"+file.Filename, fileBlob, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "Success",
	})

}
