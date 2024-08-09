# Pokedex CLI

This is a CLI application developed with learning process, as part of the backend developer roadmap of [boot.dev](https://www.boot.dev/tracks/backend)

# Commands

- `help` -> Retrieve a list of commands
- `map` -> Retrieve a list of locations
- `map-back` -> Goes back on listed locations*
- `explore {location}` -> Retrieve a list of pokemons in a location
- `catch {pokemon}` -> Catches a pokemon
- `inspect` -> Retrieves information of a catched pokemon
- `pokedex` -> Retrieves a list of catched pokemons
- `exit` -> Exits the application

*The application makes usage of [PokéAPI](https://pokeapi.co/) to retrieve the list of pokemons and locations, when using `map` and `map-back` commands results as cached to avoid multiple requests to the API and improve the performance of the application.

Example:

```shell
$ map
```

Output:

```shell
- canalave-city-area
- eterna-city-area
- pastoria-city-area
- sunyshore-city-area
- sinnoh-pokemon-league-area
- oreburgh-mine-1f
- oreburgh-mine-b1f
- valley-windworks-area
- eterna-forest-area
- fuego-ironworks-area
- mt-coronet-1f-route-207
```

```shell
$ explore canalave-city-area
```

Output:

```shell
Pokemons in canalave-city-area:
 - tentacool
 - tentacruel
 - staryu
 - magikarp
 - gyarados
 - wingull
 - pelipper
 - shellos
 - gastrodon
 - finneon
 - lumineon
```

```shell
$ catch gyarados
``` 

Output:

```shell
Pokemon in tentacool:
Trying to catch the pokemon...C - tentacool
You caught tentacool!
```





