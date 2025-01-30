🎬 Movies CRUD App
Welcome to the Movies CRUD App! This application is built using Golang and allows you to perform CRUD (Create, Read, Update, Delete) operations on a collection of movies. The app does not use any database; instead, it utilizes structs to manage movie data.

🚀 Features
Get All Movies: Retrieve a list of all movies.
Get Movie by ID: Fetch details of a specific movie by its ID.
Create Movie: Add a new movie to the collection.
Update Movie: Modify details of an existing movie.
Delete Movie: Remove a movie from the collection.
🛠️ Endpoints
GET /movies - Get all movies
GET /movies/{id} - Get a movie by ID
POST /movies - Create a new movie
PUT /movies/{id} - Update a movie by ID
DELETE /movies/{id} - Delete a movie by ID
📦 Installation
Clone the repository:
git clone https://github.com/yourusername/movies-crud-app.git

Navigate to the project directory:
cd movies-crud-app

Run the application:
go run main.go

📋 Usage
Use the following cURL commands to interact with the API:

Get All Movies:
curl -X GET http://localhost:8080/movies

Get Movie by ID:
curl -X GET http://localhost:8080/movies/{id}

Create Movie:
curl -X POST -H "Content-Type: application/json" -d '{"id":"1","title":"Inception","director":"Christopher Nolan"}' http://localhost:8080/movies

Update Movie:
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Inception","director":"Christopher Nolan"}' http://localhost:8080/movies/{id}

Delete Movie:
curl -X DELETE http://localhost:8080/movies/{id}

📚 Code Snippets
Here’s a glimpse of the main routes in the application:

Go

r.HandleFunc("/movies", getMovies).Methods("GET")
r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
r.HandleFunc("/movies", createMovie).Methods("POST")
r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
AI-generated code. Review and use carefully. More info on FAQ.
🤝 Contributing
Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.
