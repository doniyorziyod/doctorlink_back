package server

import (
	"doctorlink/types"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Repository interface {
	SaveUser(u types.User) error
	AllUsersRepo() ([]types.User, error)
	GetUserRepo(id uuid.UUID) (types.User, error)
	SignInRepo(phone string) (uuid.UUID, error)
	GenerateSmsRepo(username string) (int, error)
	GenerateCode() int
	UpdateUserRepo(u types.User) (bool, error)
	SaveBotUser(u types.BotUser) error
}

type Handler struct {
	repo Repository
}

func NewRouter(repo Repository) *echo.Echo {
	e := echo.New()
	h := Handler{repo: repo}
	e.POST("/signup", h.SignUp)
	e.GET("/users", h.AllUsers)
	e.GET("/users/:id", h.GetUser)
	e.GET("/signin/:username", h.SignIn)

	return e
}

func (h *Handler) SignUp(c echo.Context) error {
	u := new(types.User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user := types.User{
		Id:        uuid.New(),
		Username:  u.Username,
		Name:      u.Name,
		Surname:   u.Surname,
		Gender:    u.Gender,
		Birthday:  u.Birthday,
		Subregion: u.Subregion,
		Region:    u.Region,
	}

	if err := h.repo.SaveUser(user); err != nil {
		log.Fatal(err)
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
	username := c.Param("username")
	id, err := h.repo.SignInRepo(username)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, id)
}

func (h *Handler) GenerateSms(c echo.Context) error {
	username := c.Param("username")
	code, err := h.repo.GenerateSmsRepo(username)
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
