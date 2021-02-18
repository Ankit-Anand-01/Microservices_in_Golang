package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

//user
type user struct {
	Name     string
	Password string
}

// Book ...
type Book struct {
	Book   string
	Author string
	Genre  string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	books := make([]Book, 0)
	/*	books = append(books, Book{
			Book:   "It",
			Author: "Stephen King",
		})
	*/
	fmt.Println(books)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//horror
	r.GET("/horror", func(c *gin.Context) {
		fmt.Println("hi welcome to my db")
		//opening of database
		db, err := sql.Open("mysql", "root:lucifer007@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}
		//select the table
		rows, err := db.Query("SELECT Book,Author FROM books where genre='Horror'")
		if err != nil {
			panic(err)
		}
		var book []Book
		for rows.Next() {
			var b1 Book
			err = rows.Scan(&b1.Book, &b1.Author)
			if err != nil {
				panic(err)
			}
			book = append(book, b1)
		}
		//closing of database
		defer db.Close()

		c.HTML(http.StatusOK, "horror.html", gin.H{"books": book})
	})

	//sci-fi
	r.GET("/sci-fi", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:lucifer007@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}
		//select the table
		rows, err := db.Query("SELECT Book,Author FROM books where genre='Sci-fi'")
		if err != nil {
			panic(err)
		}
		var book2 []Book
		for rows.Next() {
			var b2 Book
			err = rows.Scan(&b2.Book, &b2.Author)
			if err != nil {
				panic(err)
			}
			book2 = append(book2, b2)
		}
		//closing of database
		defer db.Close()
		c.HTML(http.StatusOK, "Sci-fi.html", gin.H{"books": book2})
	})

	//Fantasy
	r.GET("/fantasy", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:lucifer007@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}
		//select the table
		rows, err := db.Query("SELECT Book,Author FROM books where genre='Fantasy'")
		if err != nil {
			panic(err)
		}
		var book3 []Book
		for rows.Next() {
			var b3 Book
			err = rows.Scan(&b3.Book, &b3.Author)
			if err != nil {
				panic(err)
			}
			book3 = append(book3, b3)
		}
		//closing of database
		defer db.Close()
		c.HTML(http.StatusOK, "fantasy.html", gin.H{"books": book3})
	})

	r.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add.html", nil)
	})

	r.POST("/adds", func(c *gin.Context) {
		Book := c.PostForm("Book")
		Author := c.PostForm("Author")
		Genre := c.PostForm("Genre")

		db, err := sql.Open("mysql", "root:lucifer007@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err.Error())
		}
		k := "insert into books values('" + Book + "','" + Author + "','" + Genre + "' )"
		result, err := db.Query(k)
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()
		c.HTML(http.StatusOK, "adds.html", nil)
	})

	//register

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	r.POST("/register-successful", func(c *gin.Context) {
		Name := c.PostForm("Name")
		Password := c.PostForm("Password")

		db, err := sql.Open("mysql", "root:lucifer007@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err.Error())
		}
		k := "insert into Users values('" + Name + "','" + Password + "' )"
		result, err := db.Query(k)
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()
		c.HTML(http.StatusOK, "register-successful.html", nil)
	})

	//print user and password
	r.POST("/login-successful", func(c *gin.Context) {
		a := 0
		Name := c.PostForm("Name")
		Password := c.PostForm("Password")
		fmt.Println(Name, Password)
		db, err := sql.Open("mysql", "root:lucifer007@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}
		//select the table
		result, err := db.Query("SELECT * from Users")
		if err != nil {
			panic(err)
		}

		for result.Next() {
			var detail user
			err = result.Scan(&detail.Name, &detail.Password)
			if err != nil {
				panic(err)
			}

			if detail.Name == Name && detail.Password == Password {
				a = 1
			}

		}

		//closing of database
		defer db.Close()
		if a == 1 {
			c.HTML(http.StatusOK, "add.html", nil)
		} else {

			c.HTML(http.StatusOK, "invalid.html", nil)
		}
	})

	//login
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.Run(":8000")
}
