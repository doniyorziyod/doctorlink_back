package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type User struct {
    Id uuid.UUID `json:"id" form:"id" query:"id"`
    Phone string `json:phone"" form:"phone" query:"phone"`
    Name string `json:name"" form:"name" query:"name"`
    Surname string `json:surname"" form:"surname" query:"surname"`
    Gender bool `json:gender"" form:"gender" query:"gender"`
    Birthday string `json:birthday"" form:"birthday" query:"birthday"`
    Subregion string `json:subregion"" form:"subregion" query:"subregion"`
    Region string `json:region"" form:"region" query:"region"`
}

type InputUser struct {
    Phone string `json:phone"" form:"phone" query:"phone"`
    Name string `json:name"" form:"name" query:"name"`
    Surname string `json:surname"" form:"surname" query:"surname"`
    Gender bool `json:gender"" form:"gender" query:"gender"`
    Birthday string `json:birthday"" form:"birthday" query:"birthday"`
    Subregion string `json:subregion"" form:"subregion" query:"subregion"`
    Region string `json:region"" form:"region" query:"region"`
}

func main() {
	// e := echo.New()
    // e.POST("/signup", signUp)
	// e.Logger.Fatal(e.Start(":1323"))
    if saveUser() != nil {
        panic(saveUser())
    }
}

func signUp(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    user := User {
        Id: uuid.New(),
        Phone: u.Phone,
        Name: u.Name,
        Surname: u.Surname,
        Gender: u.Gender,
        Birthday: u.Birthday,
        Subregion: u.Subregion,
        Region: u.Region,
    }

    return c.JSON(http.StatusOK, user.Id)
}

func saveUser() error {
    db, err := sqlx.Connect("postgres", "user=doniyorziyod dbname=doctorlink sslmode=disable password=7355950d host=localhost")
    defer db.Close()
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Successfully Connected")
    }

    place := User {}
    rows, _ := db.Queryx("SELECT * FROM users")
    for rows.Next() {
    err := rows.StructScan(&place) // Scan the current row into the "place" variable
    if err != nil {
        log.Fatalln(err)
    }
    log.Printf("%#v\n", place) // Log the content of the "place" struct for each ro
}

    return err
}
