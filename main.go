// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// //book struct(model)
// type Book struct {
// 	ID     string  `json: "id"`
// 	Isbn   string  `json: "isbn"`
// 	Title  string  `json: "title"`
// 	Author *Author `json: "author"`
// }

// //author struct
// type Author struct {
// 	Firstname string `json: "firstname"`
// 	Lastname  string `json: "lastname"`
// }

// // init books var as a slice book struct
// var books []Book

// //get all books
// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// //get single book
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) // get params
// 	//loopthrough books and find with id
// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }

// //create book
// func createBook(w http.ResponseWriter, r *http.Request) {
// 	rand.Seed(time.Now().UnixNano())
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	json.NewDecoder(r.Body).Decode(&book)

// 	// len := r.ContentLength
// 	// body := make([]int, len)
// 	// fmt.Println(w, string(len))
// 	len := r.ContentLength
// 	body := make([]byte, len)
// 	r.Body.Read(body)
// 	fmt.Fprintln(w, string(body))
// 	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock id - not safe
// 	book.Isbn = strconv.Itoa(rand.Intn(100000))
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)

// }
// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = strconv.Itoa(rand.Intn(10000000)) // mock id - not safe
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)

// }
// func deleteBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)

// }

// func main() {
// 	//creating digigs and strings for Isbn and Title
// 	// randomMaker()
// 	//init router
// 	r := mux.NewRouter()

// 	// mock data - @todo - implement data base
// 	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
// 	books = append(books, Book{ID: "2", Isbn: "84756", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

// 	//route handlers / endpoints
// 	r.HandleFunc("/api/books", getBooks).Methods("GET")
// 	r.HandleFunc("/api/books/{isbn}", getBook).Methods("GET")
// 	r.HandleFunc("/api/books", createBook).Methods("POST")
// 	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
// 	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

// 	log.Fatal(http.ListenAndServe(":8000", r))
// }
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Age struct {
	Id   int  `json:"id"`
	User User `json:"user"`
}

var users []User

func main() {
	users = append(users, User{Id: "1", Name: "riku", Email: "riku.kuri1234@"})
	users = append(users, User{Id: "2", Name: "lu", Email: "lutianye89@"})
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", mainHandler)
	e.POST("/post", post)
	e.GET("/delete/:id", deleteByHTML)
	//htmlはDELETE使えない？e.DELETEではなくe.GETにしたら動いた
	e.DELETE("/delete/:id", delete)
	e.GET("/get", show)
	e.Logger.Fatal(e.Start(":1323")) //e.loggerがe.post,e.getより先に書かれているとmessage not foundとなる。なぜか。
}

//ハンドラーを定義

func mainHandler(c echo.Context) error {
	t, err := template.ParseFiles("main.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(c.Response(), users); err != nil {
		panic(err.Error())
	}
	return c.String(http.StatusOK, "")
}

// func (t *Template) ParseFiles(filenames ...string) (*Template, error)
// func (t *Template) Execute(wr io.Writer, data interface{}) error
func show(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
func post(c echo.Context) error {
	var u User
	//htmlからのPOST
	name := c.FormValue("name")
	email := c.FormValue("email")
	if err := c.Bind(&u); err != nil {
		// error handling
	}
	if name != "" && email != "" {
		fmt.Println("done")
		u.Name = name
		u.Email = email
	}
	//curlからのPOST
	userId := len(users) + 1
	u.Id = strconv.Itoa(userId)
	users = append(users, u)
	fmt.Println(users)
	fmt.Println("name: ", u.Name, "email: ", u.Email)
	fmt.Println(name, email)
	return c.JSON(http.StatusOK, u)
}
func deleteByHTML(c echo.Context) error {
	//htmlからのDELETE
	id := c.FormValue("id")
	for index, item := range users {
		if item.Id == id {
			fmt.Println("done by html", id)
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	return c.JSON(http.StatusOK, users)

}
func delete(c echo.Context) error {
	fmt.Println("done")
	//htmlからのDELETE
	id := c.FormValue("id")
	for index, item := range users {
		if item.Id == id {
			fmt.Println("done by html of delete", id)
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	//curlからのDELETE
	params := c.Param("id")
	for index, item := range users {
		//"delete/:idなのに、なぜcurlでは:をつけずにidを指定してdeleteしなければならないのか
		if item.Id == params {
			fmt.Println("done by curl")
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	return c.JSON(http.StatusOK, users)
}
