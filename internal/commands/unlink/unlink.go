package unlink

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tabletop_matchmaker/internal/db/entities"
	"tabletop_matchmaker/internal/helpers/errors"
)

type Unlink struct{}

func (unlink Unlink) Run(msg *tgbotapi.Message, db *sql.DB) []tgbotapi.Chattable {
	err := entities.DeleteFromBggAccount(db, msg.From.ID)

	if err != nil {
		errors.UnexpectedChatErrorMessage(err, msg.Chat.ID)
	}

	return []tgbotapi.Chattable{tgbotapi.NewMessage(msg.Chat.ID, "Заебумба")}
}

func (unlink Unlink) Callback(_ *tgbotapi.CallbackQuery, _ *sql.DB) []tgbotapi.Chattable {
	return nil
}

func Name() string {
	return "unlink"
}
