package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	// every movie has a director
	//in the Movie struct it means that each Movie object contains a field named Director, which is a pointer to a Director struct.
    
	//each movie object can have associated information about its director

	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// w-> used to send http response data back to client
// r->contains info about the incoming http request
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content type", "application/json")
	//json.NewEncoder(w)-> this creates a new json encoder that will write 
    //encoded data to 'w'

	//.Encode(movies)-> This encodes the movies data into JSON format using the JSON encoder created in the previous step. 


	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	//it informs the client that the response body will be in json format

	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r) //to extract the variables from the request url

	for index, item := range movies {
		if item.ID == params["id"] {
			//The ... operator is used to unpack the elements of a slice into the append function as individual arguments.
			movies = append(movies[:index], movies[index+1:]...)
			break

		}
		json.NewEncoder(w).Encode(movies)

	}
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}

}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	var movie Movie
	//The & operator is used to pass a pointer to the movie variable so that Decode can modify its contents.

	_ = json.NewDecoder(r.Body).Decode(&movie)
	//This line generates a unique ID for the new movie. It uses rand.Intn to generate a random integer between 0 and 100,000,000, and then strconv.Itoa to convert that integer to a string.

	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content type","applicatioin/json")
	params:=mux.Vars(r)
	for index,item:=range movies{
		if item.ID==params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			var movie Movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			movies=append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func main() {
	//mux is a popular HTTP request router and dispatcher for Go.

	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
	//defining routes

	//HandleFunc is a method (or function) provided by the router or web framework to define routes and their associated handler functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	//by using {id} in the URL pattern, you're telling the router or framework to capture whatever value comes after "/movies/" in the URL

	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
