package movies

import (
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"encoding/json"

	"github.com/gorilla/mux"
	"os"
)

type Movies struct {
	Title string `gorm:"column:title"`
	Released_Year 	string `gorm:"column:released_year"`
	Rating string 	`gorm:"column:rating"`
	Id string 		`gorm:"column:id"`
	Generes string 	`gorm:"column:generes"`

}

type MoviesResp struct {
	Title string
	Released_Year 	string
	Rating string
	Id string
	Generes string

}


func GetMovieByTitleOrId(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open("mysql", "root:root@/movies?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		fmt.Println("Inside error")
	}
	searchtype:= r.URL.Query().Get("type")
	fmt.Println("Connected",db)
	movies := Movies{}
	db.SingularTable(true)

	db.Where("title = ?", searchtype).Or("id = ?", searchtype).Find(&movies)

	movieresp:= MoviesResp{}

	movieresp.Id=movies.Id
	movieresp.Title=movies.Title
	movieresp.Rating=movies.Rating
	movieresp.Released_Year=movies.Released_Year
	movieresp.Generes=movies.Generes
	_ = json.NewEncoder(w).Encode(&movieresp)

	defer db.Close()
}

func UpdateRating(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open("mysql", "root:root@/movies?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		fmt.Println("Inside error")
	}

	rating:= r.URL.Query().Get("rating")
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Println("Connected",db)
	movies := Movies{}
	db.SingularTable(true)
	_ = json.NewEncoder(os.Stdout).Encode(db.Where("title = ?", title).Find(&movies))
	movies.Rating=rating
	db.Model(&movies).Updates(movies)
	movieresp:= MoviesResp{}

	movieresp.Id=movies.Id
	movieresp.Title=movies.Title
	movieresp.Rating=movies.Rating
	movieresp.Released_Year=movies.Released_Year
	movieresp.Generes=movies.Generes
	_ = json.NewEncoder(w).Encode(&movieresp)

	defer db.Close()

}

func GetMovieByYear(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open("mysql", "root:root@/movies?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		fmt.Println("Inside error")
	}
	year:= r.URL.Query().Get("year")
	fmt.Println("Connected",db)
	movies := []Movies{}
	db.SingularTable(true)
	_ = json.NewEncoder(os.Stdout).Encode(db.Where("released_year = ?", year).Find(&movies))
	movie:= []MoviesResp{}
	for _,element:= range movies{
		movieresp:= MoviesResp{}
		movieresp.Id=element.Id
		movieresp.Title=element.Title
		movieresp.Rating=element.Rating
		movieresp.Released_Year=element.Released_Year
		movieresp.Generes=element.Generes
		movie=append(movie,movieresp)
	}
	_ = json.NewEncoder(w).Encode(&movie)
	defer db.Close()
}

func GetMovieByRating(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open("mysql", "root:root@/movies?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		fmt.Println("Inside error")
	}
	rating:= r.URL.Query().Get("rating")
	fmt.Println("Connected",db)
	movies := []Movies{}
	db.SingularTable(true)
	_ = json.NewEncoder(os.Stdout).Encode(db.Where("rating >= ?", rating).Find(&movies))
	movie:= []MoviesResp{}
	for _,element:= range movies{
		movieresp:= MoviesResp{}
		movieresp.Id=element.Id
		movieresp.Title=element.Title
		movieresp.Rating=element.Rating
		movieresp.Released_Year=element.Released_Year
		movieresp.Generes=element.Generes
		movie=append(movie,movieresp)
	}
	_ = json.NewEncoder(w).Encode(&movie)
	defer db.Close()
}


