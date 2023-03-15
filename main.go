package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Dapil struct {
	Dapil []int  `json:"dapil"`
	Nama  string `json:"nama"`
}

type Region struct {
	ID   int     `json:"id"`
	Data []Dapil `json:"id"`
}

func main() {
	e := echo.New()

	e.GET("/hasil", func(c echo.Context) error {

		// Path to JSON file

		// Read JSON file
		// jsonFile, err := os.Open("data/region.json")
		jsonFile, err := os.ReadFile("data/region.json")
		if err != nil {
			return err
		}
		// fmt.Println(jsonFile)

		// Parse JSON data into Region struct
		// readJSON, _ := io.ReadAll(jsonFile)

		var regions map[string]Region

		if err := json.Unmarshal([]byte(jsonFile), &regions); err != nil {
			return err
		}

		// Return JSON response with the parsed data
		fmt.Println(regions)

		return c.JSON(http.StatusOK, regions)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
