package server

import (
	"doctorlink/types"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Repository interface {
    SaveUser(u types.User) error
    AllUsersRepo() ([]types.User, error)
    GetUserRepo(id uuid.UUID) (types.User, error)
    SignInRepo(phone string) (types.User, error)
    GenerateSmsRepo(phone string) (int, error)
    GenerateCode() int
    UpdateUserRepo(u types.User) (bool, error)
}

type Bot interface {

}

type Handler struct {
    repo Repository
    bot Bot
}

func NewRouter(repo Repository) *echo.Echo {
	e := echo.New()
    h := Handler{ repo: repo }
    e.POST("/signup", h.SignUp)
    e.GET("/users", h.AllUsers)
    e.GET("/users/:id", h.GetUser)
    e.GET("/signin/:phone", h.GetUser)

    return e
}

func (h *Handler) SignUp(c echo.Context) error {
    u := new(types.User)
    if err := c.Bind(u); err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    user := types.User {
        Id: uuid.New(),
        Phone: u.Phone,
        Name: u.Name,
        Surname: u.Surname,
        Gender: u.Gender,
        Birthday: u.Birthday,
        Subregion: u.Subregion,
        Region: u.Region,
    }

    if err := h.repo.SaveUser(user); err != nil {
        panic(err)
    }

    return c.JSON(http.StatusOK, user.Id)
}

func (h *Handler) AllUsers(c echo.Context) error {
    users, err := h.repo.AllUsersRepo()
    if err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUser(c echo.Context) error {
    id := c.Param("id")
    user, err := h.repo.GetUserRepo(uuid.MustParse(id))
    if err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    return c.JSON(http.StatusOK, user)
}

func (h *Handler) SignIn(c echo.Context) error {
    phone := c.Param("phone")
    user, err := h.repo.SignInRepo(phone)
    if err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    return c.JSON(http.StatusOK, user)
}

func (h *Handler) GenerateSms(c echo.Context) error {
    phone := c.Param("phone")
    code, err := h.repo.GenerateSmsRepo(phone)
    if err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    return c.JSON(http.StatusOK, code)
}

func (h *Handler) UpdateUser(c echo.Context) error {
    u := new(types.User)
    if err := c.Bind(u); err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    a, err := h.repo.UpdateUserRepo(*u)
    if err != nil {
        return c.String(http.StatusBadRequest, "bad request")
    }

    return c.JSON(http.StatusOK, a)
}
