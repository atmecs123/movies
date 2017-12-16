package main

import (
    // Standard library packages
    "context"
    "flag"
    "os"
    "fmt"
    "log"
    "net/http"
    "os/signal"
    "strconv"


    // Third party packages
    "github.com/gorilla/mux"
    "gclassec/movies"
)

func main() {
        mx := mux.NewRouter()

        //Find movie by title or ID by exact value that’s passed in the API.
        mx.HandleFunc("/movie/detailbyIdorTitle", movies.GetMovieByTitleOrId).Methods("GET") //http://localhost:9009/movie/detail?type or id ={title} or {id}

        //API that allows updates to genres and ratings of the movie.
        mx.HandleFunc("/movie/rating/{title}", movies.UpdateRating).Methods("PUT") //http://localhost:9009/movie/rating/{title}?rating={rating}

        //Search movies released in a particular year or given range
        mx.HandleFunc("/movie/year", movies.GetMovieByYear).Methods("GET") //http://localhost:9009/movie/detail?year={year}

        //Search movies with rating higher or lower than passed in value.
        mx.HandleFunc("/movie/ratingrange", movies.GetMovieByRating).Methods("GET") //http://localhost:9009/movie/ratingrange?rating={rating}


        portPtr := flag.Int("port", 9000, "an int")
        flag.Parse()

        portNum := strconv.Itoa(*portPtr)
        fmt.Println("Port:", *portPtr)

        http.Handle("/", mx)

        fmt.Println("Server is on Port: ", *portPtr)

        p := "0.0.0.0:"+portNum
        fmt.Println("Listening at .....",p)
        stop := make(chan os.Signal)
        signal.Notify(stop, os.Interrupt)
        addr := ":" + portNum
        h := &http.Server{Addr: addr}
        slog := log.New(os.Stdout, "", 0)

        go func() {
            slog.Printf("Listening:= %s", addr)
            if err := h.ListenAndServe(); err != nil {
                slog.Fatal(err)
            }
        }()

        <-stop
        h.Shutdown(context.Background())
        }







