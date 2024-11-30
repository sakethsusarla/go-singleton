package main

import (
	"fmt"
	"sync"
)

// Synchronization to ensure single initialization
var once sync.Once

// Singleton struct to encapsulate Sheldon's preferences
type SheldonConfiguration struct {
	Spot                    string
	FavoriteDrink           string
	FavoriteMeal            string
	TemperatureInFahrenheit int
}

// Singleton instance of Sheldon's preferences
var instance *SheldonConfiguration

// GetSheldonConfiguration ensures only one instance is created
func GetSheldonConfiguration() *SheldonConfiguration {
	if instance == nil {
		once.Do(func() {
			instance = &SheldonConfiguration{
				Spot:                    "Sheldon's spot on the couch",
				FavoriteDrink:           "Diet Virgin Cuba Libre",
				FavoriteMeal:            "Spaghetti with little pieces of hot dog cut up in it",
				TemperatureInFahrenheit: 72,
			}

			fmt.Println("Created instance of SheldonConfiguration")
			fmt.Println("--------------------")
			fmt.Println("Spot: ", instance.Spot)
			fmt.Println("FavoriteDrink: ", instance.FavoriteDrink)
			fmt.Println("FavoriteMeal: ", instance.FavoriteMeal)
			fmt.Println("TemperatureInFahrenheit: ", instance.TemperatureInFahrenheit)
			fmt.Println("--------------------")
		})
	} else {
		fmt.Println("Instance of SheldonConfiguration already exists")
	}

	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		go GetSheldonConfiguration()
	}

	fmt.Scanln()
}
