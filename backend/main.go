package main

import (
	"log"
	"net/http"

	"cs-be/cypher"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("encode/caesar", encodeCaesar)
	r.POST("decode/caesar", decodeCaesar)
	r.POST("encode/monoaplphabetic", encodeMono)
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
