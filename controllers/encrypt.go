package controllers

import (
	"github.com/LuanSapelli/golang-gcm-cipher/settings"
	"github.com/LuanSapelli/golang-gcm-cipher/utils"
	"github.com/gin-gonic/gin"
	"log"
)


func EncryptData(context *gin.Context) {

	type requestBody struct {
		DecryptedString string `json:"string"`
	}

	body := new(requestBody)
	err := context.Bind(body)
	if err != nil {
		log.Printf("Error to get body - %v", err)
	}

	decryptedString := body.DecryptedString

	gcm, err := utils.NewGCM(settings.NewRegularKey().Key, settings.NewRegularKey().Iv)
	if err != nil {
		log.Printf("Error in GCM - %v", err)
	}

	encryptedString, err := gcm.Encrypt(decryptedString)
	if err != nil {
		log.Printf("Error to encrypt - %v", err)
	}

	context.JSON(200, gin.H{
		"encrypted_string": encryptedString,
	})

}
