package main
import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type movie struct {
	ID string `json:"id"` // Importante la notacion
	Title string `json:"title"`
	Producer string `json:"producer"`
	Year int `json:"year"`
}
var movies= [] movie{
	{ID: "1", Title: "FTX", Producer:"David BU", Year: 2021},
	{ID: "2", Title: "ABC", Producer:"David CX", Year: 2022},
	{ID: "3", Title: "DEC", Producer:"David MO", Year: 2023},
}

// Usar el context de ginx
func getMovies(c *gin.Context){
	c.IndentedJSON(http.StatusOK, movies)
}

// Post para mandar movies
func postMovies(c *gin.Context){
	var newMovie movie
	if err:= c.BindJSON(&newMovie); err!= nil{
		return
	}
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, movies) 
}
// Otro handler
func getMovieByID(c *gin.Context){
	id := c.Param("id")
	for _, a:= range movies {
		if a.ID ==id {
			c.IndentedJSON(http.StatusOK, a)
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"messag":"Movie no encontrada"})
}

func main() {
	//fmt.Println("Hola soy David")
	router := gin.Default()
	router.GET("/movies",getMovies)
	router.POST("/movies",postMovies)
	router.GET("/movies/:id", getMovieByID)
	router.Run("localhost:8080")
}