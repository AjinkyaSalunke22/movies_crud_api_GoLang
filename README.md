# 🎬 Movies CRUD API

A simple RESTful API built with Go (Golang) that allows you to manage a collection of movies. This application implements basic CRUD (Create, Read, Update, Delete) operations using in-memory storage with structs.

## 🚀 Features

- Get all movies
- Get a specific movie by ID
- Create new movies
- Update existing movies
- Delete movies
- No database required - uses in-memory storage

## 🛠️ Prerequisites

- Go (1.16 or later)
- Gorilla Mux router

## 📦 Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/movies-crud-app.git
```

2. Navigate to the project directory:
```bash
cd movies-crud-app
```

3. Install dependencies:
```bash
go mod tidy
```

4. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## 🔄 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/movies` | Get all movies |
| GET | `/movies/{id}` | Get a movie by ID |
| POST | `/movies` | Create a new movie |
| PUT | `/movies/{id}` | Update a movie by ID |
| DELETE | `/movies/{id}` | Delete a movie by ID |

## 📝 API Usage

### Get All Movies
```bash
curl -X GET http://localhost:8080/movies
```

### Get Movie by ID
```bash
curl -X GET http://localhost:8080/movies/1
```

### Create Movie
```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "id": "4",
    "isbn": "324532",
    "title": "New Movie",
    "duration": 165,
    "director": {
        "firstname": "John",
        "lastname": "Doe"
    }
}' http://localhost:8080/movies
```

### Update Movie
```bash
curl -X PUT -H "Content-Type: application/json" -d '{
    "isbn": "324532",
    "title": "Updated Movie",
    "duration": 170,
    "director": {
        "firstname": "John",
        "lastname": "Doe"
    }
}' http://localhost:8080/movies/1
```

### Delete Movie
```bash
curl -X DELETE http://localhost:8080/movies/1
```

## 📊 Data Structure

```go
type Movie struct {
    ID       string    `json:"id"`
    Isbn     string    `json:"isbn"`
    Title    string    `json:"title"`
    Duration int       `json:"duration"`
    Director *Director `json:"director"`
}

type Director struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
}
```

## 🔧 Sample Data

The application comes with some sample movie data:

```go
movies = append(movies, Movie{
    ID: "1",
    Isbn: "324332",
    Title: "Movie One",
    Duration: 180,
    Director: &Director{
        Firstname: "Ajinkya",
        Lastname: "Salunke"
    }
})
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 📞 Support

For support, please open an issue in the GitHub repository.
