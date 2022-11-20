package bd

import (
	"context"
	"time"

	"github.com/avallarino-ar/golang_lab/src/github.com/avallarino-ar/golang_lab/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckExistUser(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("golang_lab")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.Usuario

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID

	}
	return result, true, ID
}
