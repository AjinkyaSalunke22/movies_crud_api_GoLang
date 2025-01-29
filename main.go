package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string   `json:"id"`
	Isbn     string   `json:"isbn"`
	Title    string   `json:"title"`
	Duration int      `json:"duration"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(movies) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"message": "No Movies Found"})
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "The movies are"})
		json.NewEncoder(w).Encode(movies)
		w.WriteHeader(http.StatusOK)
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Movie deleted with given ID"})
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "No Movie to delete with this ID"})
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(map[string]string{"message": "Movie with given ID is:"})
			json.NewEncoder(w).Encode(item)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "No Movie with this ID"})
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie.ID = generateID()
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(map[string]string{"message": "New movie created"})
	json.NewEncoder(w).Encode(movie)
	w.WriteHeader(http.StatusCreated)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
				json.NewEncoder(w).Encode(map[string]string{"message": "Cannot decode the given data"})
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(map[string]string{"message": "Movie updated with this ID"})
			json.NewEncoder(w).Encode(movie)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Movie not found with the given ID"})
}

func generateID() string {
	return "id" + strconv.Itoa(rand.Intn(100000)) + "aj" + strconv.Itoa(rand.Intn(100000)) + "sal"
}

func main() {
	r := mux.NewRouter()

	// movies = append(movies, Movie{ID: "1", Isbn: "324332", Title: "Movie One", Duration: 180, Director: &Director{Firstname: "Ajinkya", Lastname: "Salunke"}})
	// movies = append(movies, Movie{ID: "2", Isbn: "324232", Title: "Movie Two", Duration: 170, Director: &Director{Firstname: "Ronak", Lastname: "Jain"}})
	// movies = append(movies, Movie{ID: "3", Isbn: "324432", Title: "Movie Three", Duration: 190, Director: &Director{Firstname: "Sarthak", Lastname: "Kamble"}})
	// movies = append(movies, Movie{ID: "1", Isbn: "324332", Title: "Movie One", Duration: 180, Director: &Director{Firstname: "Ajinkya", Lastname: "Salunke"}})
	// movies = append(movies, Movie{ID: "2", Isbn: "324232", Title: "Movie Two", Duration: 170, Director: &Director{Firstname: "Ronak", Lastname: "Jain"}})
	// movies = append(movies, Movie{ID: "3", Isbn: "324432", Title: "Movie Three", Duration: 190, Director: &Director{Firstname: "Sarthak", Lastname: "Kamble"}})
	// movies = append(movies, Movie{ID: "1", Isbn: "324332", Title: "Movie One", Duration: 180, Director: &Director{Firstname: "Ajinkya", Lastname: "Salunke"}})
	// movies = append(movies, Movie{ID: "2", Isbn: "324232", Title: "Movie Two", Duration: 170, Director: &Director{Firstname: "Ronak", Lastname: "Jain"}})
	// movies = append(movies, Movie{ID: "3", Isbn: "324432", Title: "Movie Three", Duration: 190, Director: &Director{Firstname: "Sarthak", Lastname: "Kamble"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server is running on port http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}