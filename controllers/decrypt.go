package controllers

import (
	"github.com/LuanSapelli/golang-gcm-cipher/settings"
	"github.com/LuanSapelli/golang-gcm-cipher/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func DecryptData(context *gin.Context) {

	type requestBody struct {
		EncryptedString string `json:"string"`
	}

	body := new(requestBody)
	err := context.Bind(body)
	if err != nil {
		log.Printf("Error to get body - %v", err)
	}

	encryptedString := body.EncryptedString

	gcm, err := utils.NewGCM(settings.NewRegularKey().Key, settings.NewRegularKey().Iv)
	if err != nil {
		log.Printf("Error in GCM - %v", err)
	}

	decryptedString, err := gcm.Decrypt(encryptedString)
	if err != nil {
		log.Printf("Error to decrypt - %v", err)
	}

	context.JSON(200, gin.H{
		"decrypted_string": decryptedString,
	})

}
