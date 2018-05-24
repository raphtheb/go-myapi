package main

// Imports
import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "strconv"
)

// The person Type (more like an object)
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person



// functions


// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

// FizzBuzz Web
func FizzBuzzWeb(w http.ResponseWriter, r *http.Request) {
    for out := range FizzBuzz(100) {
        json.NewEncoder(w).Encode(out)
    }
}

// FizzBuzz Custom
func FizzBuzzCustom(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    amount, err := strconv.Atoi(params["amount"])
    if err == nil {
        for out := range FizzBuzz(amount) {
            json.NewEncoder(w).Encode(out)
        }
    } else {
      json.NewEncoder(w).Encode("ERROR: Your custom fizzbuzz parameter isn't an integer.")
    }
}

//FizzBuzz routine
func FizzBuzz(amount int) <-chan string {
    out := make(chan string, amount)
    go func() {
        for i :=1; i <= amount; i++ {
            result := ""
            if i%3 == 0 {result += "Fizz" }
            if i%5 == 0 {result += "Buzz" }
            if result == "" { result = fmt.Sprintf("%v", i) }
            out <- result
        }
        close(out)
    }()
    return out
}


// main function to boot up everything
func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
    people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/fizzbuzz", FizzBuzzWeb).Methods("GET")
    router.HandleFunc("/fizzbuzz/{amount}", FizzBuzzCustom).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}
