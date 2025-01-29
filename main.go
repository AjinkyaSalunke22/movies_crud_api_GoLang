package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)

// "fmt"
// "log"
// "encoding/json"
// "math/rand"
// "net/http"
// "strconv"
// "github.com/gorilla/mux"

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"Laastname"`
}

var movies []Movie 


// Get all movies fun

func getMovies(w http.ResponseWriter, r *http.Request){ // Pointer is due to 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) 
}

// Delete a movie
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


func createMovie (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie 
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


func updateMovie(w http.ResponseWriter, r http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main(){
	r := mux.NewRouter() // create a new router

	movies = append(movies, Movie{ID: "1", Isbn: "324332", Title: "Movie One", Director: &Director{Firstname: "Ajinkya", Lastname: "Salunke"}})
	movies = append(movies, Movie{ID: "2", Isbn: "324232", Title: "Movie Two", Director: &Director{Firstname: "Ronak", Lastname: "Jain"}})
	movies = append(movies, Movie{ID: "3", Isbn: "324432", Title: "Movie Three", Director: &Director{Firstname: "Sarthak", Lastname: "kamble"}})

	// Routes and Handlers
	r.handleFunc("/movies", getMovies).Methods("GET")
	r.handleFunc("movies/id", getMovie).Methods("GET")
	r.handleFunc("movies", createMovie).Methods("POST")
	r.handleFunc("movies/id", updateMovie).Methods("PUT")
	r.handleFunc("movies/id", deleteMovie).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r)) // logout the error if any
	
}