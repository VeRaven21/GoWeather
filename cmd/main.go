package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/VeRaven21/GoWeather/internal/api"
	"github.com/VeRaven21/GoWeather/internal/config"
	"github.com/VeRaven21/GoWeather/internal/printer"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v3"
)

func main() {
	godotenv.Load()

	cmd := &cli.Command{
		Name:  "GoWeather",
		Usage: "A simple CLI tool to fetch weather; usage: GoWeather [city]",

		// Main program function
		Action: func(ctx context.Context, cmd *cli.Command) error {

			city := cmd.Args().Get(0)

			if city == "" {
				config, err := config.LoadConfig("config.yaml")
				if err != nil {
					log.Fatal(err)
				}

				if config.DefaultCity != "" {
					city = config.DefaultCity
				} else {
					fmt.Println("Please provide a city name as an argument or set a default city in config.yaml")
					return nil
				}

			}

			fmt.Printf("Собираю информацию о городе %s \n", city)

			geoposition, err := api.GetGeolocation(city)
			if err != nil && err.Error() == "city not found" {
				fmt.Printf("Город %s не найден\n", city)
				return nil
			} else if err != nil {
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
