1. Inicializar modulo
```go
go mod init github.com/gin_02
Ya te debe aparecer el maejador de modulos
```
2. Creamos el archivo `main.go` con
```go
package main
import "fmt"

func main() {
	fmt.Println("Hola soy David")
}
```
3. Testaemos que todo ande bien
```bash
go run .
```
Te debe salir el mensaje que programamos
4. Ahora debemos installar gin, puedes ir a google y colocar gin golang `https://gin-gonic.com/`, con este comando lo puedes descargar en la terminal
```bash
go get -u github.com/gin-gonic/gin
```
5. Ahora modificamos el `main.go`
```go
package main
import (
	"fmt"
	"github.com/gin-gonic/gin",
	"http"
)
type movie struct {
	ID string `json:"id"` // Importante la notacion
	Title string `json:"title"`
	Producer string `json:"producer"`
	Year int `json:"year"`
}
var movies= [] movie{
	{ID: "1", Title: "FTX", Producer:"David BU", Year: 2021},
	{ID: "2", Title: "ABC", Producer:"David CX", Year: 2021}
	{ID: "2", Title: "DEC", Producer:"David MO", Year: 2021}
}

// Usar el ocntext ge ginx
func getMovies(c *gin.Context){
	c.IndentedJSON(http.StatusOk, movies)
}

func main() {
	//fmt.Println("Hola soy David")
	router := gin.Default()
	router.GET("/movies",getMovies)
	router.Run("localhost:8080")

}
```
6. Instalar ThunderClient en los complementos de VS Code 
7. Ahora levantamos el servicio con
```bash
go run.
```
8. Vas al metodo GET `http://localhost:8080/movies` en thunderClient y deberas obtener la estructura con la data
```json
[
  {
    "id": "1",
    "title": "FTX",
    "producer": "David BU",
    "year": 2021
  },
  {
    "id": "2",
    "title": "ABC",
    "producer": "David CX",
    "year": 2022
  },
  {
    "id": "3",
    "title": "DEC",
    "producer": "David MO",
    "year": 2023
  }
]
```
9. Ahora modificamos el `main.go` para agregar un metodo POST
```go
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
	c.BindJSON(&newMovie)
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, movies) 
}
func main() {
	//fmt.Println("Hola soy David")
	router := gin.Default()
	router.GET("/movies",getMovies)
	router.POST("/movies",postMovies)
	router.Run("localhost:8080")

}
```
10. Debes detener el servivor en la terminal y volver a levantalro
```bash
go run .
```
11. Ahora puedes ir a `http://localhost:8080/movies` con metodo POST y en el body poner:
```json
{
    "id": "4",
    "title": "CCF",
    "producer": "Cidys FT",
    "year": 2024
  }
```
Te debera devolver el `Status: 201 Created` y ademas todos los registros
12. Puedes comprobarlo con el GET `http://localhost:8080/movies`
13. Pero OJO la data se te va a perder cuando bajes el server porque esta en memoria
14. Modificamos los metodos por si hay algun error en el input en el `main.go`
```go
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
```
15. Detienes el server y lo vuelves a levantar 
```bash
go run .
```
16. Ahora puedes ir al metodo GET `http://localhost:8080/movies/1` y te debe devolver el dato filtrado
```json
{
  "id": "1",
  "title": "FTX",
  "producer": "David BU",
  "year": 2021
}
```
En caso contrario si pones un id que no existe por eg. 6 te debe devolver un 404 Not found 
```json
{
  "messag": "Movie no encontrada"
}
```
