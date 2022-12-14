package server

import (
	"SOTBI/telegram-notify/config"
	"context"
	"fmt"
	"github.com/nikoksr/notify/service/telegram"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikoksr/notify"
	"github.com/rs/zerolog/log"

	"SOTBI/telegram-notify/model"
)

const Subject = "JIRA"

type Server struct {
	notifyClient notify.Notifier
}

func New(cfg *config.Config) (*Server, error) {
	telegramClient, err := telegram.New(cfg.Token)
	if err != nil {
		return nil, err
	}

	telegramClient.AddReceivers(cfg.Channel)

	notifyClient := notify.New()
	notifyClient.UseServices(telegramClient)

	return &Server{notifyClient: notifyClient}, nil
}

func (s *Server) NewEvent(c *gin.Context) {
	payload := &model.JiraPayload{}

	if err := c.BindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	message := fmt.Sprintf("Issue: %s\nSummary: %s",
		payload.Issue.Key,
		payload.Issue.Fields.Summary,
	)

	if err := s.notifyClient.Send(context.Background(), Subject, message); err != nil {
		log.Err(err)

		return
	}
}
