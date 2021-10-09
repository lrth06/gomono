package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/gomono/config"
	"github.com/lrth06/gomono/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func GetPosts(c *fiber.Ctx) error {
	return c.SendString("Success!")
}

func Ping(c *fiber.Ctx) error {
	return c.SendString("Server is available and healthy!")
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello from Go API 👋!")
}

func UpdateUser(c *fiber.Ctx) error {
	userId, err := primitive.ObjectIDFromHex(
		c.Params("id"),
	)
	if err != nil {
		return c.SendStatus(400)
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: user.Name},
				{Key: "email", Value: user.Email},
			},
		},
	}
	err = config.ConnDB("Users").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	collection := config.ConnDB("Users")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		println(user)
		return c.Status(400).SendString(err.Error())
	}

	insertionResult, err := collection.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)
	createdUser := &models.User{}
	createdRecord.Decode(createdUser)

	return c.Status(201).JSON(filter)
}

func GetUsers(c *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := config.ConnDB("Users").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var users []models.User = make([]models.User, 0)

	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString(err.Error())

	}
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userId, err := primitive.ObjectIDFromHex(idParam)
	fmt.Println(userId)
	if err != nil {
		return c.SendStatus(400)
	}

	var user bson.D
	if err = config.ConnDB("Users").FindOne(c.Context(), bson.D{}).Decode(&user); err != nil {
		fmt.Println("error")
	}

	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	userId, err := primitive.ObjectIDFromHex(
		c.Params("id"),
	)

	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: userId}}
	result, err := config.ConnDB("Users").DeleteOne(c.Context(), &query)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.SendStatus(204)
}
