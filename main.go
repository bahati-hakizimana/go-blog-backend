package main

import(
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/bahati-hakizimana/blogue-backend/database"
	"github.com/bahati-hakizimana/blogue-backend/routes"
)

func main(){
database.Connect()
error:=godotenv.Load()
if error !=nil{
	log.Fatal("Error load .env files")
}else{
	log.Println("Connected successfuly")
}
port:=os.Getenv("PORT")
app:=fiber.New()
routes.Setup(app)
app.Listen(":"+port)
}