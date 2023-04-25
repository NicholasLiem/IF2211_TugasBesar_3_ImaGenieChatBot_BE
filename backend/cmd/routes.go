package main

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/chat_session"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/home"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/messages"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/question_answer"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", home.GetHome)
	// Chat sessions
	app.Post("/chat-sessions", chat_session.CreateChatSession)
	app.Delete("/chat-sessions/:session_id", chat_session.DeleteChatSession)
	app.Get("/chat-sessions", chat_session.GetChatSessions)

	// Chat messages
	app.Post("/chat-sessions/:session_id/messages", messages.MessageHandler)
	app.Get("/chat-sessions/:session_id/messages", messages.GetChatMessages)

	// Question answers
	app.Get("/question-answers", question_answer.GetQuestionAnswers)

	// Get all chat session IDs
	app.Get("/chat-sessions-id", chat_session.GetChatSessionIDs)
}
