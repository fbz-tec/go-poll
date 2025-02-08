package api

import (
	"net/http"
	"strconv"

	db "github.com/fbz-tec/go-poll/db/sqlc"
	"github.com/gin-gonic/gin"
)

// -------------- CreateVote API -------------------------------
type CreateVoteRequest struct {
	OptionID int64  `json:"option_id"`
	Voter    string `json:"voter"`
}

func (server *Server) CreateVoteHandler(ctx *gin.Context) {
	var req CreateVoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	vote, err := server.store.CreateVote(ctx, db.CreateVoteParams{
		OptionID: req.OptionID,
		Voter:    req.Voter,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, vote)
}

// -------------------- ListVotes API -----------------------
type ListVotesRequest struct {
	PageId   int32 `form:"page" binding: "required,min=1"`
	PageSize int32 `form:"page_size" binding: "required,min=5,max=20"`
}

func (server *Server) ListVotesHandler(ctx *gin.Context) {
	var req ListVotesRequest
	pollPram := ctx.Param("pollId")
	pollId, err := strconv.Atoi(pollPram)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	votes, err := server.store.GetVotesByPoll(ctx, db.GetVotesByPollParams{
		PollID: int64(pollId),
		Limit:  req.PageId,
		Offset: (req.PageId - 1) * req.PageSize,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, votes)

}

// ------------------- GetVoteResult ----------------
type GetVoteResultRequest struct {
	PollId int64 `uri:"pollId" binding:"required,min=1"`
}

func (server *Server) GetVoteResultHandler(ctx *gin.Context) {
	var req GetVoteResultRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	result, err := server.store.GetTotalVotes(ctx, req.PollId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
