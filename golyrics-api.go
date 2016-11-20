package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mamal72/golyrics"
	"gopkg.in/kataras/iris.v4"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func searchLyrics(query string, ctx *iris.Context) {
	suggestions, err := golyrics.SearchTrack(query)
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	if len(suggestions) == 0 {
		ctx.EmitError(iris.StatusNotFound)
		return
	}

	track := suggestions[0]
	fetchErr := track.FetchLyrics()
	if err != fetchErr {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	ctx.JSON(iris.StatusOK, track)
}

func searchLyricsByQuery(ctx *iris.Context) {
	query := ctx.Param("query")
	searchLyrics(query, ctx)
}

func main() {
	godotenv.Load()
	iris.Get("/search/:query", searchLyricsByQuery)
	iris.Listen(fmt.Sprintf("%s:%s", getenv("HOST", "0.0.0.0"), getenv("PORT", "8080")))
}
