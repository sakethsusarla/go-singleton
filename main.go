package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type SingletonConfiguration struct {
	Spot                    string
	FavoriteDrink           string
	FavoriteMeal            string
	TemperatureInFahrenheit int
}

var instance *SingletonConfiguration

func GetInstance() *SingletonConfiguration {
	if instance == nil {
		once.Do(
			func() {
				instance = &SingletonConfiguration{
					Spot:                    "Sheldon's spot on the couch",
					FavoriteDrink:           "Diet Virgin Cuba Libre",
					FavoriteMeal:            "Spaghetti with little pieces of hot dog cut up in it",
					TemperatureInFahrenheit: 72,
				}

				fmt.Println("Created instance of SingletonConfiguration")
				fmt.Println("--------------------")
				fmt.Println("Spot: ", instance.Spot)
				fmt.Println("FavoriteDrink: ", instance.FavoriteDrink)
				fmt.Println("FavoriteMeal: ", instance.FavoriteMeal)
				fmt.Println("TemperatureInFahrenheit: ", instance.TemperatureInFahrenheit)
				fmt.Println("--------------------")
			},
		)
	} else {
		fmt.Println("Instance of SingletonConfiguration already exists")
	}

	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		go GetInstance()
	}

	fmt.Scanln()
}
