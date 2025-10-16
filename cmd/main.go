package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/VeRaven21/GoWeather/internal/api"
	"github.com/VeRaven21/GoWeather/internal/printer"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v3"
)

func main() {
	godotenv.Load()

	cmd := &cli.Command{
		Name:  "GoWeather",
		Usage: "A simple CLI tool to fetch weather; usage: GoWeather [city]",
		Action: func(ctx context.Context, cmd *cli.Command) error {

			city := cmd.Args().Get(0)

			fmt.Printf("Собираю информацию о городе %s \n", city)

			//Получаем положение города
			geoposition, err := api.GetGeolocation(city)

			if err != nil {
				log.Fatal(err)
			}

			weather, err := api.GetWeather(*geoposition)
			if err != nil {
				log.Fatal(err)
			}
			printer.PrintWeather(*weather)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
