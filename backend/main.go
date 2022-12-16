package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cs-be/cypher"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("encode/caesar", encodeCaesar)
	r.POST("decode/caesar", decodeCaesar)
	r.POST("encode/monoaplphabetic", encodeMono)
	r.POST("decode/monoaplphabetic", decodeMono)
	r.POST("encode/polyalphabetic", encodeMono)
	r.POST("decode/polyalphabetic", decodeMono)
	r.POST("encode/polyalphabetic/photo", encodePolyPhoto)
	r.POST("decode/polyalphabetic/photo", decodePolyPhoto)
	r.Run()
}

func encodeCaesar(ginctx *gin.Context) {
	var message cypher.Text
	err := ginctx.BindJSON(&message)
	if err != nil {
		log.Println("Error on binding, err:", err)
		ginctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	ginctx.JSON(http.StatusOK, gin.H{"text": message.Caesar(cypher.ENCODE)})
}

func decodeCaesar(ginctx *gin.Context) {
	var message cypher.Text
	err := ginctx.BindJSON(&message)
	if err != nil {
		log.Println("Error on binding, err:", err)
		ginctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	ginctx.JSON(http.StatusOK, gin.H{"text": message.Caesar(cypher.DECODE)})
}

func encodeMono(ginctx *gin.Context) {
	var message cypher.Text
	err := ginctx.BindJSON(&message)
	if err != nil {
		log.Println("Error on binding, err:", err)
		ginctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	ginctx.JSON(http.StatusOK, gin.H{"text": message.Monoalphabetic(cypher.ENCODE)})
}

func decodeMono(ginctx *gin.Context) {
	var message cypher.Text
	err := ginctx.BindJSON(&message)
	if err != nil {
		log.Println("Error on binding, err:", err)
		ginctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	ginctx.JSON(http.StatusOK, gin.H{"text": message.Monoalphabetic(cypher.DECODE)})
}

func encodePolyPhoto(ginctx *gin.Context) {
	inputFile, _, err := ginctx.Request.FormFile("file")
	if err != nil {
		log.Println("Error on binding, err:", err)
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	defer inputFile.Close()

	b, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Println("error reading file", err)
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	b64Image := base64.StdEncoding.EncodeToString(b)
	request := cypher.Text{
		Key:     ginctx.PostForm("key"),
		Message: b64Image,
	}

	b64Encrypted := request.Monoalphabetic(cypher.ENCODE)
	output, _ := os.Create("./output/fileEncryptedPolyPhoto.jpg")
	defer output.Close()
	_, err = output.Write([]byte(b64Encrypted))
	ginctx.JSON(http.StatusOK, gin.H{"text": b64Encrypted})
}

func decodePolyPhoto(ginctx *gin.Context) {
	inputFile, _, err := ginctx.Request.FormFile("file")
	if err != nil {
		log.Println("Error on binding, err:", err)
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	defer inputFile.Close()

	encryptedFile, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Println("error reading file", err)
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	request := cypher.Text{
		Key:     ginctx.PostForm("key"),
		Message: string(encryptedFile),
	}
	b64Decrypted := request.Monoalphabetic(cypher.DECODE)
	decode, _ := base64.StdEncoding.DecodeString(b64Decrypted)
	output, err := os.Create("./output/fileDecodedPolyPhoto.jpg")
	fmt.Println(err)
	defer output.Close()
	_, err = output.Write(decode)
	fmt.Println(err)
	ginctx.JSON(http.StatusOK, gin.H{"text": "File is ready :)"})
}
