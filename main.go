package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	router := gin.New()
	router.GET("/test/hello", Hello)
	log.Printf("server is running...")
	log.Fatal(router.Run("test:8080"))

}

func Hello(c *gin.Context) {
	// Firebase konfiguratsiyasini yaratish
	config := &firebase.Config{
		ProjectID: "test-3aa32",
	}

	// Firebase ilovasini ishga tushirish
	opt := option.WithCredentialsFile("test-3aa32-firebase-adminsdk-ot0zs-de5bba82f1.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("Ilovani ishga tushirishda xatolik: %v", err)
	}

	// FCM mijozini olish
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Messaging mijozini olishda xatolik: %v", err)
	}
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	// Xabar yaratish
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Yangi xabar",
			Body:  "Bu test xabari",
		},
		Token: token,
	}

	// Xabarni yuborish
	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalf("Xabarni yuborishda xatolik: %v", err)
	}

	fmt.Printf("Xabar muvaffaqiyatli yuborildi. Response: %s\n", response)
}
