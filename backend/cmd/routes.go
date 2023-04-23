package main

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/chat_session"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/messages"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/pages"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/question_answer"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/chat/", pages.ChatPage)

	app.Post("/a/:session_id", messages.InsertMessageToChatSession)

	app.Get("/h/:session_id", chat_session.GetChatMessages)
	app.Post("/c/:session_id", chat_session.CreateChatSession)
	app.Post("/r/:session_id", chat_session.DeleteChatSession)

	app.Post("/create", question_answer.CreateQuestionAnswer)
	app.Post("/remove", question_answer.DeleteQuestionAnswer)
	app.Get("/questions", question_answer.GetQuestionAnswers)
}
