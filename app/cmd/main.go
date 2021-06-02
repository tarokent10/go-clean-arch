package main

//go:generate sqlboiler --wipe mysql
import (
	"fmt"
	"log"
	_ "picture-go-app/infrastructure/config"
	"picture-go-app/infrastructure/database"
	"picture-go-app/infrastructure/server"
	"picture-go-app/injector"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	fmt.Println("main -start")
	defer fmt.Println("main -end")
	// initialization
	conn, err := database.Open()
	if err != nil {
		panic(err)
	}
	diContainer, err := injector.InitDIContainer(conn)
	if err != nil {
		log.Fatal(err)
	}
	// ここじゃない気がするけど
	boil.DebugMode = true
	server.Run(diContainer)
}
