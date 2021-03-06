package logic

import (
  "net/http"
)

var (
  ErrInvalidTimeFormat = ApiError(http.StatusBadRequest, "invalid time format")
  ErrInvalidUser = ApiError(http.StatusBadRequest,"invalid user")
  ErrNotFriend = ApiError(http.StatusBadRequest,"you can only view friend's todo list")
)
