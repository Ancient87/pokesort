package main

import (
//        "errors"
//        "strings"
        "net/http"
        "encoding/json"
        "fmt"
        "time"
        "sort"
        "log"
        "io/ioutil"
)

var pokeApiEndpoint = "http://pokeapi.co/api/v2"
var pokeApiName = "/pokemon/"



// Data structure to present a simple Pokemon
type Pokemon struct {
    Name string `json:"name"`
    Id int `json:"id"`
}

// Define a slice of monsters as a type
type Pokemons []*Pokemon

type By func(p1, p2 *Pokemon) bool


// Method on By to sort to sort argument slice
func (by By) Sort(pokemon Pokemons) {
    ps := &pokeSorter {
        pokemon: pokemon,
        by:      by, // 
    }
    sort.Sort(ps) // Invoke sort library on pokeSorter
}


// Implement sort interface by returning length of array
func (s pokeSorter) Len() int { return len(s.pokemon) }
// Implement sort interface by providing a way to swap pokemon in an array slice
func (s pokeSorter) Swap(i, j int) { s.pokemon[i], s.pokemon[j] = s.pokemon[j], s.pokemon[i] }

// byId implements Sort.Interface as Pokemons have Len and Swap and byId provides a Less
//type byId struct { Pokemons }

// Complete sort interface implementation for ByID
func (s pokeSorter) Less(i, j int) bool { return s.pokemon[i].Id < s.pokemon[j].Id }

// struct that sorts a slice of pokemon based on the function provided
// Satisfies sorter interface by virtue of defining a function of type by
type pokeSorter struct {
    pokemon Pokemons
    by      By
}

// Knows how to talk to the PokeAPI
type querier interface {
    getPokemonByName(pokemon string) string
}

// PokeService does things with pokemon
type Service interface {
    Sort(pokemonNames... string) (Pokemons, error)
}

// pokeService is a struct that implements Service int
type pokeService struct {
    q querier
}

// Constructor for pokeService return snew pokeService
func NewPokeService() *pokeService {
    return &pokeService{
        q:   nil,
    }
}

// Take list of pokemonnames and sort them by ID
func (s *pokeService) SortPokemonByIDFromName(pokemonNames []string) *Pokemons {
    ps := s.getPokemonsByName(pokemonNames)
    s.sortPokemonByID(ps)
    return &ps
}

// Take a bunch of pokemon and soer them by their ID
func (s *pokeService) sortPokemonByID(p Pokemons)  {
        ps := &pokeSorter {
        pokemon: p,
//        by: byId,
    }
    // Invoke sort library on the planetsorter
    sort.Sort(ps)
}


// Sorts pokemon and returns a list of sorted pokemon
// Makes pokeService implement Service interface
func (s *pokeService) getPokemonsByName(pokemonNames []string) Pokemons {
    var monsters Pokemons
    // Invoke makePokemonFromName for each name passed
    for _, name := range pokemonNames {
        p := s.getPokemonByName(name)
        monsters = append(monsters, &p)
    }
    return monsters
}

// Queries PokeAPI and returns Pokemon Struct for one Pokemon by Name
func (s *pokeService) getPokemonByName(pokemonName string) Pokemon {
    // Call Poke API and get pokemon json
    client := http.Client{
        Timeout: time.Second *20,
    }
    url := fmt.Sprintf("%s%s%s", pokeApiEndpoint, pokeApiName, pokemonName)
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
   }
   req.Header.Set("User-Agent", "Go Test")

   res, getErr := client.Do(req)
   if getErr != nil {
        log.Fatal(getErr)
   }

   body, readErr := ioutil.ReadAll(res.Body)
   if readErr != nil {
        log.Fatal(readErr)
   }
   pokemon := Pokemon{}
   jsonErr := json.Unmarshal(body, &pokemon)

   if jsonErr != nil {
        log.Fatal(jsonErr)
   }
   return pokemon
}


/*
func makeQuerier() querier {
    pokeApiClient := http.Client{
        Timeout: time.Second *2,
   }
*/


