package api

import (
	"net/http"

	db "github.com/fbz-tec/go-poll/db/sqlc"
	"github.com/gin-gonic/gin"
)

// ----------------------- ListPolls API -----------------------
type ListPollsRequest struct {
	PageId   int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) ListPollsHandler(ctx *gin.Context) {
	var request ListPollsRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := db.ListPollsParams{
		Limit:  request.PageSize,
		Offset: (request.PageId - 1) * request.PageSize,
	}
	polls, err := server.store.ListPolls(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, polls)
}

// ----------------------- GetPoll API -----------------------
type getPollRequest struct {
	ID int64 `uri:"pollId" binding:"required,min=1"`
}

func (server *Server) GetPollHander(ctx *gin.Context) {
	var request getPollRequest
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// 2 queries to get poll and options data
	poll, err := server.store.GetPoll(ctx, request.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	options, err := server.store.GetOptions(ctx, poll.PollID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := struct {
		db.Poll
		Options []db.Option `json:"options"`
	}{
		Poll:    poll,
		Options: options,
	}

	ctx.JSON(http.StatusOK, response)

}

// ----------------------- createPoll API -----------------------}
type createPollRequest struct {
	Question string   `json:"question"`
	Owner    string   `json:"owner"`
	Options  []string `json:"options"`
}

func (server *Server) CreatePollHander(ctx *gin.Context) {

	var request createPollRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//todo: use transaction
	poll, err := server.store.CreatePoll(ctx, db.CreatePollParams{
		Question: request.Question,
		Owner:    request.Owner,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	options := make([]db.Option, len(request.Options))

	for _, option := range request.Options {
		opt, err := server.store.CreateOption(ctx, db.CreateOptionParams{
			PollID:      poll.PollID,
			OptionValue: option,
		})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		options = append(options, opt)
	}

	response := struct {
		db.Poll
		Options []db.Option `json:"options"`
	}{
		Poll:    poll,
		Options: options,
	}

	ctx.JSON(http.StatusCreated, response)
}
