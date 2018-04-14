package main

import "testing"

func TestPokeService(t *testing.T) {
    //names := []string{"lapras", "mewtoo", "pikachu", "raichu", "zapdos" }
    pokemons := Pokemons{
        &Pokemon { Name: "charmander", Id: 4 },
        &Pokemon { Name: "zapdos", Id: 145 },
        &Pokemon { Name: "raichu", Id: 26 },
        &Pokemon { Name: "lapras", Id: 131 },
    }
    p := NewPokeService()

    // Test getting pokemon By Name
    //ps := testGetPokemonByName(t, p, names)
    p.sortPokemonByID(pokemons)
    t.Logf("SORTED INIT %s", pokemons)
    // Test Sorting
    /*
    pokemons := p.SortPokemonByIDFromName(names)
    if pokemons == nil {
        t.Errorf("lol")
    }
    */
    //if something != nil {
    //    t.Errotf("Roflcopter")
    //}

}

func testGetPokemonByName(t *testing.T, s *pokeService, names []string) *Pokemons {
    ps := Pokemons{}
    for _, name := range names {
        p := s.getPokemonByName(name)
        t.Logf("I got myself a %s", p)
        if p.Name != name {
            t.Errorf(" Expected %s but got %s", name, p.Name)
        }
        ps = append(ps, &p)
    }
    t.Logf("Gotta catch them all %s", ps)
    return &ps
}
