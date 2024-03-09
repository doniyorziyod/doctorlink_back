package bot

import (
	"doctorlink/types"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	tele "gopkg.in/telebot.v3"
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

var (
	hello = "👋"
	world = " Assalomu alaykum! Doctorlink platformasiga xush kelibsiz"
)

func NewBot(repo Repository) *tele.Bot {
	pref := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	h := Handler{repo: repo}
	b.Handle("/start", h.Start)

	return b
}

func (h *Handler) Start(c tele.Context) error {
	u := c.Sender()
	user := types.BotUser{
		ID:                u.ID,
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		IsForum:           u.IsForum,
		Username:          u.Username,
		LanguageCode:      u.LanguageCode,
		IsBot:             u.IsBot,
		IsPremium:         u.IsPremium,
		AddedToMenu:       u.AddedToMenu,
		Usernames:         u.Usernames,
		CustomEmojiStatus: u.CustomEmojiStatus,
		CanJoinGroups:     u.CanJoinGroups,
		CanReadMessages:   u.CanReadMessages,
		SupportsInline:    u.SupportsInline,
	}

	if err := h.repo.SaveBotUser(user); err != nil {
		log.Fatal(err)
	}

	return c.Send(hello + world)
}
