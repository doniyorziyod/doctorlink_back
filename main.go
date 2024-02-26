package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

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

type Sms struct {
    Phone string `json:phone"" form:"phone" query:"phone"`
    Sms int `json:sms"" form:"sms" query:"sms"`
}

func main() {
	e := echo.New()
    e.POST("/signup", signUp)
    e.GET("/users", allUsers)
    e.GET("/users/:id", getUser)
    e.GET("/signin/:phone", getUser)
	e.Logger.Fatal(e.Start(":1323"))
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

    if err := saveUser(user); err != nil {
        panic(err)
    }

    return c.JSON(http.StatusOK, user.Id)
}

func allUsers(c echo.Context) error {
    users, err := allUsersRepo()
    if err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
    id := c.Param("id")
    user, err := getUserRepo(uuid.MustParse(id))
    if err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    return c.JSON(http.StatusOK, user)
}

func signIn(c echo.Context) error {
    phone := c.Param("phone")
    user, err := signInRepo(phone)
    if err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    return c.JSON(http.StatusOK, user)
}

func generateSms(c echo.Context) error {
    phone := c.Param("phone")
    code, err := generateSmsRepo(phone)
    if err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    return c.JSON(http.StatusOK, code)
}

func updateUser(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    a, err := updateUserRepo(*u)
    if err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    return c.JSON(http.StatusOK, a)
}

func connectDB() (*sqlx.DB, error) {
    db, err := sqlx.Connect("postgres", "user=doniyorziyod dbname=doctorlink sslmode=disable password=7355950d host=localhost")

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Successfully Connected")
    }

    return db, err
}

func saveUser(u User) error {
    db, err := connectDB()
    defer db.Close()
    _, err = db.Exec("INSERT INTO users (id, phone, name, surname, gender, birthday, subregion, region) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
    u.Id, u.Phone, u.Name, u.Surname, u.Gender, u.Birthday, u.Subregion, u.Region)

    return err
}

func allUsersRepo() ([]User, error) {
    db, err := connectDB()
    defer db.Close()
    rows, _ := db.Query("SELECT * from users")
    users := []User{}
    
    for rows.Next() {
        user := User{}
        if err = rows.Scan(&user.Id, &user.Phone, &user.Name, &user.Surname, &user.Gender, &user.Birthday, &user.Subregion, &user.Region); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, err
}

func getUserRepo(id uuid.UUID) (User, error) {
    users, err := allUsersRepo()
    for i := 0; i < len(users); i++ {
       if users[i].Id == id {
           return users[i], err
       } 
    }

    return User{}, err
}

func signInRepo(phone string) (User, error) {
    users, err := allUsersRepo()
    for i := 0; i < len(users); i++ {
       if users[i].Phone == phone {
           return users[i], err
       } 
    }

    return User{}, err
}

func generateSmsRepo(phone string) (int, error) {
    db, err := connectDB()
    defer db.Close()
    sms := Sms {
        Phone: phone,
        Sms: generateCode(),
    }
    _, err = db.Exec("INSERT INTO sms (phone, sms) VALUES ($1, $2)",
    sms.Phone, sms.Sms)

    return sms.Sms, err
}

func generateCode() int {
    rand.New(rand.NewSource(time.Now().UnixNano()))
    return rand.Intn(90000) + 10000
}

func updateUserRepo(u User) (bool, error) {
    db, err := connectDB()
    defer db.Close()
    _, err = db.Exec("UPDATE users SET phone=$1, name=$2, surname=$3, gender=$4, birthday=$5, subregion=$6, region=$7 WHERE id=$1",
    u.Id, u.Phone, u.Name, u.Surname, u.Gender, u.Birthday, u.Subregion, u.Region)

    return true, err
}
