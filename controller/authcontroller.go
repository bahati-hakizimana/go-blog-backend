package controller

import (
	"fmt"
	"time"
	"strconv"
	"log"
	"strings"
	"regexp"
	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm/utils"
	"github.com/bahati-hakizimana/blogue-backend/models"
	"github.com/bahati-hakizimana/blogue-backend/util"
	"github.com/bahati-hakizimana/blogue-backend/database"
	"github.com/dgrijalva/jwt-go"
)
func validateEmail(email string) bool{
	Re:=regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}
func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err:=c.BodyParser(&data);err!=nil{
		fmt.Println("Unable to parse body")
	}


	// check if password is less than 6 character

	if len(data["password"].(string))<=6{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"password must be greaterthan 6 characters",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))){
		c.Status(400)
		return c.JSON(fiber.Map{

			"message":"Invalid Email address",

		})
		
	}
	//check if Email is already exist in database
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id!=0{
		c.Status(400)
		return c.JSON(fiber.Map{

			"message":"Email Already Exist",
	})
}
    //save into data base
	user:=models.User{
		FirstName:data["first_name"].(string),
		LastName:data["last_name"].(string),
		Email:strings.TrimSpace(data["email"].(string)),
		Phone:data["phone"].(string),
	}

	user.SetPassword(data["password"].(string))
	err:=database.DB.Create(&user)
	if err != nil{
    log.Println(err)

	}
	    c.Status(200)
         return c.JSON(fiber.Map{
			"user":user,
		"message":"Account created Successfuly",

	})
}
func Login(c *fiber.Ctx) error{
	var data map[string]string
	if err:=c.BodyParser(&data);err!=nil{
		fmt.Println("Unable to parse body")
	}

	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)

	if user.Id ==0{
		c.Status(404)
		return c.JSON(fiber.Map{

			"message": "Email Address doesn't exist, Kindly create an account", 
			
		})
		
	}

	if err:=user.ComparePassword(data["password"]); err !=nil{
		c.Status(4000)
		return c.JSON(fiber.Map{
			"message":"Incorrect password",
		})
	}
	token, err:= util.GenerateJWT(strconv.Itoa(int(user.Id)),)

	if err != nil{
		c.Status(fiber.StatusInternalServerError)

		return nil

	}

	cookie:=fiber.Cookie{
		Name:"jwt",
		Value:token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message":"You have Successful Loged in",
		"user":user,
	})

	
}
type Claims struct{
	jwt.StandardClaims
}

