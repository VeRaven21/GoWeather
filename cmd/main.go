package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/VeRaven21/GoWeather/internal/api"
	"github.com/VeRaven21/GoWeather/internal/config"
	"github.com/VeRaven21/GoWeather/internal/printer"

	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "GoWeather",
		Usage: "A simple CLI tool to fetch weather; usage: GoWeather [city]",

		Action: func(ctx context.Context, cmd *cli.Command) error {
			defaultCity := cmd.String("default-city")

			// Если передан флаг --default-city — сохраняем его в конфиг и используем
			if defaultCity != "" {
				err := config.SaveConfig(&config.Config{
					DefaultCity: defaultCity,
					Language:    "ru",
				})
				if err != nil {
					log.Fatalf("Failed to save config: %v", err)
				}
				// Используем указанный город как целевой
				city := defaultCity

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
			}

			// Если флаг не передан — работаем как обычно
			city := cmd.Args().Get(0)

			if city == "" {
				cfg, err := config.LoadConfig()
				if err != nil {
					log.Fatal(err)
				}

				if cfg.DefaultCity != "" {
					city = cfg.DefaultCity
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

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "default-city",
				Aliases: []string{"c"},
				Value:   "",
				Usage:   "Add default city to config",
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
