package api

import (
	db "github.com/fbz-tec/go-poll/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Queries
	router *gin.Engine
}

func NewServer(store *db.Queries) (*Server, error) {
	server := &Server{
		store:  store,
		router: gin.Default(),
	}
	router := server.router

	v1 := router.Group("/api/v1")
	{
		polls := v1.Group("polls")
		{
			polls.GET("", server.ListPollsHandler)
			polls.POST("", server.CreatePollHander)
			polls.GET(":pollId", server.GetPollHander)
			polls.GET(":pollId/votes", server.ListVotesHandler)
			polls.GET(":pollId/results", server.GetVoteResultHandler)
		}
		votes := v1.Group("votes")
		{
			votes.POST("", server.CreateVoteHandler)
		}
		users := v1.Group("users")
		{
			users.POST("", server.CreateUserHandler)
			users.GET(":user_id", nil)
			users.GET(":user_id/votes", nil)
		}
	}
	return server, nil
}

func (server *Server) Start(listenAddr string) error {
	return server.router.Run(listenAddr)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func simpleResponse(msg string) gin.H {
	return gin.H{
		"msg": msg,
	}
}
