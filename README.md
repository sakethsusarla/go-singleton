## Singleton Example - ShelBot Configuration

### Context
This example demonstrates how to implement a Singleton pattern in Go to create a configuration for ShelBot. The Singleton pattern ensures that a class has only one instance and provides a global point of access to it.

### Problem
In this example, we want to create ShelBot, and we need some configuration for that. The configuration should be read only once.

### Solution
The `SingletonConfiguration` struct holds the configuration details for ShelBot:

- `Spot`: Sheldon's spot on the couch
- `FavoriteDrink`: Diet Virgin Cuba Libre
- `FavoriteMeal`: Spaghetti with little pieces of hot dog cut up in it
- `TemperatureInFahrenheit`: 72

The `GetInstance` function returns the instance of `SingletonConfiguration`. It uses the `sync.Once` type to ensure that the configuration is initialized only once.