package types

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id" form:"id" query:"id"`
	Username  string    `json:username"" form:"phone" query:"username"`
	Name      string    `json:name"" form:"name" query:"name"`
	Surname   string    `json:surname"" form:"surname" query:"surname"`
	Gender    bool      `json:gender"" form:"gender" query:"gender"`
	Birthday  string    `json:birthday"" form:"birthday" query:"birthday"`
	Subregion string    `json:subregion"" form:"subregion" query:"subregion"`
	Region    string    `json:region"" form:"region" query:"region"`
}

type InputUser struct {
	Username  string `json:username"" form:"username" query:"username"`
	Name      string `json:name"" form:"name" query:"name"`
	Surname   string `json:surname"" form:"surname" query:"surname"`
	Gender    bool   `json:gender"" form:"gender" query:"gender"`
	Birthday  string `json:birthday"" form:"birthday" query:"birthday"`
	Subregion string `json:subregion"" form:"subregion" query:"subregion"`
	Region    string `json:region"" form:"region" query:"region"`
}

type Sms struct {
	Username string `json:username"" form:"username" query:"username"`
	Sms      int    `json:sms"" form:"sms" query:"sms"`
}

type BotUser struct {
	ID                int64    `json:"id"`
	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	IsForum           bool     `json:"is_forum"`
	Username          string   `json:"username"`
	LanguageCode      string   `json:"language_code"`
	IsBot             bool     `json:"is_bot"`
	IsPremium         bool     `json:"is_premium"`
	AddedToMenu       bool     `json:"added_to_attachment_menu"`
	Usernames         []string `json:"active_usernames"`
	CustomEmojiStatus string   `json:"emoji_status_custom_emoji_id"`
	CanJoinGroups     bool     `json:"can_join_groups"`
	CanReadMessages   bool     `json:"can_read_all_group_messages"`
	SupportsInline    bool     `json:"supports_inline_queries"`
}
