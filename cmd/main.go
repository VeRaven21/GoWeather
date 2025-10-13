package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/VeRaven21/GoWeather/internal/api"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "GoWeather",
		Usage: "A simple CLI tool to fetch weather; usage: GoWeather [city]",
		Action: func(ctx context.Context, cmd *cli.Command) error {

			apikey := ""

			city := cmd.Args().Get(0)

			fmt.Printf("Собираю информацию о городе %s \n", city)

			//Получаем положение города
			geoposition, err := api.GetGeolocation(ctx, city, apikey)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Местоположение города %s - %v, %v", city, geoposition.Lat, geoposition.Lon)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
