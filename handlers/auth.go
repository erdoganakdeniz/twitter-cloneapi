package handlers

import (
	"github.com/erdoganakdeniz/models"
	"github.com/erdoganakdeniz/utils"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type AuthHandlerInterface interface {
	Login(ctx *fiber.Ctx) interface{}
	Signup(ctx *fiber.Ctx) interface{}
}

type AuthHandler struct {
	UsersColl *mongo.Collection
}

func (a AuthHandler) Login(c *fiber.Ctx) {
	u := new(models.LoginInputs)

	if err := c.BodyParser(u); err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}

	user := new(models.User)

	filter := bson.M{"email": u.Email}
	err := a.UsersColl.FindOne(c.Fasthttp, filter).Decode(user)

	if err != nil {
		c.Status(fiber.StatusUnauthorized).Send(fiber.Map{"message": "Invalid Credentials"})
		return
	}



	isMatch := utils.Password{Password: u.Password}.Compare(user.Password)

	if !isMatch {
		c.Status(fiber.StatusUnauthorized).Send(fiber.Map{"message": "Invalid Credentials"})
		return
	}

	accessToken, err := utils.CreateJWTToken(map[string]interface{}{
		"username": user.UserName,
		"email":    user.Email,
		"id":       user.ID,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Successfully",
		"data":    fiber.Map{"token": accessToken},
	})

	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}
}

func (a AuthHandler) Signup(c *fiber.Ctx) {
	inputs := new(models.SignupInputs)


	if err := c.BodyParser(inputs); err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}


	if err := inputs.Validate(); err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}

	query := bson.D{{Key: "email", Value: inputs.Email}}

	existingUser := new(models.User)
	err := a.UsersColl.FindOne(c.Fasthttp, query).Decode(existingUser)

	if err != nil {

		if err.Error() != "mongo: no documents in result" {
			log.Fatal(err)
			return
		}
	}

	if existingUser.ID != "" {
		c.Status(fiber.StatusForbidden).Send(fiber.Map{"message": "User already exists"})
		return
	}

	p := utils.Password{Password: inputs.Password}
	hashPassword := p.Hash()

	user := models.User{
		Email:     inputs.Email,
		Password:  hashPassword,
		UserName:  inputs.UserName,
		Posts:     []primitive.ObjectID{},
		Following: []primitive.ObjectID{},
		Followers: []primitive.ObjectID{},
	}

	// force MongoDB to always set its own generated ObjectIDs
	user.ID = ""
	insertionResult, err := a.UsersColl.InsertOne(c.Fasthttp, user)

	if err != nil {
		log.Fatal(err)
	}


	createdUser := new(models.User)
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

	if err := a.UsersColl.FindOne(c.Fasthttp, filter).Decode(createdUser); err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}

	if err := c.Status(fiber.StatusCreated).JSON(createdUser); err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}
}