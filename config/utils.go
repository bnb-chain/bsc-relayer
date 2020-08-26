package util

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func GetSecret(secretName, region string) (string, error) {
	//Create a Secrets Manager client
	sess, err := session.NewSession(&aws.Config{
		Region: &region,
	})
	if err != nil {
		return "", err
	}

	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return "", err
	}

	var secretString, decodedBinarySecret string
	if result.SecretString != nil {
		secretString = *result.SecretString
		return secretString, nil
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		length, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			return "", err
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:length])
		return decodedBinarySecret, nil
	}
}

func SendTelegramMessage(botId string, chatId string, msg string) {
	if botId == "" || chatId == "" || msg == "" {
		return
	}

	endPoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botId)
	formData := url.Values{
		"chat_id":    {chatId},
		"parse_mode": {"html"},
		"text":       {msg},
	}
	_, err := http.PostForm(endPoint, formData)
	if err != nil {
		fmt.Println(fmt.Sprintf("send telegram message error, bot_id=%s, chat_id=%s, msg=%s, err=%s", botId, chatId, msg, err.Error()))
		return
	}
}
