package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mamal72/golyrics"
	"gopkg.in/kataras/iris.v4"
)

func searchLyrics(query string, ctx *iris.Context) {
	lyrics, err := golyrics.SearchAndGetLyrics(query)
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	ctx.JSON(iris.StatusOK, iris.Map{"lyrics": lyrics})
}

func searchLyricsByQuery(ctx *iris.Context) {
	query := ctx.Param("query")
	searchLyrics(query, ctx)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	iris.Get("/search/:query", searchLyricsByQuery)

	iris.Listen(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}
