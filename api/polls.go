package api

import (
	"net/http"

	db "github.com/fayca121/go-poll/db/sqlc"
	"github.com/gin-gonic/gin"
)

// ----------------------- ListPolls API -----------------------
type ListPollsRequest struct {
	PageId   int32 `form:"page" binding:"required,min=1"`
	pageSize int32 `form:"page_size" binding:"required,min=5"`
}

func (server *Server) ListPollsHandler(ctx *gin.Context) {
	var request ListPollsRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := db.ListPollsParams{
		Limit:  request.pageSize,
		Offset: (request.PageId - 1) * request.pageSize,
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
	ID int64 `uri:"id" binding:"required,min=1"`
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

	optionItems := make([]db.Option, len(options))

	for i, opt := range options {
		optionItems[i] = db.Option{
			OptionID:    opt.OptionID,
			OptionValue: opt.OptionValue,
		}
	}

	response := struct {
		db.Poll
		Options []db.Option `json:"options"`
	}{
		Poll:    poll,
		Options: optionItems,
	}

	ctx.JSON(http.StatusOK, response)

}

// ----------------------- createPoll API -----------------------
func (server *Server) CreatePollHander(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, simpleResponse("not implemented"))
}
