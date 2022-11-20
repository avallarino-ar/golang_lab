package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN es el objeto de conexion a la BD */
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://abasharino:Alfa.001@cluster0.wrdd8va.mongodb.net/?retryWrites=true&w=majority")

/* ConnectDB es la func que permite conectar a la BD */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("connection successful")
	return client
}

/* CheckConnection es el ping a la BD */
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
