package bd

import (
	"context"
	"time"

	"github.com/avallarino-ar/golang_lab/src/github.com/avallarino-ar/golang_lab/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertUser fun para agregar usuario a la BD*/
func InsertUser(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("golang_lab")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
