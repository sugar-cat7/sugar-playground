package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()

	// Firestore クライアントの初期化
	projectID := "keen-index-402705"
	databaseID := "test3" // Firestore のデフォルトデータベースIDは `(default)` です。
	client, err := firestore.NewClientWithDatabase(ctx, projectID, databaseID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// データを書き込む
	ref, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"firstName": "John",
		"lastName":  "Doe",
	})
	if err != nil {
		log.Fatalf("Failed adding a new user: %v", err)
	}

	// データを読み取る
	doc, err := client.Collection("users").Doc(ref.ID).Get(ctx)
	if err != nil {
		log.Fatalf("Failed reading a user: %v", err)
	}
	user := doc.Data()
	fmt.Printf("User: %v\n", user)
}
