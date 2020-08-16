package handler

import (
	"net/http"
	"strconv"

	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/output"

	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/repository"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	roomRepo repository.Room
}

func NewRoomHandler(roomRepo repository.Room) *RoomHandler {
	return &RoomHandler{roomRepo: roomRepo}
}

func (h *RoomHandler) GetRoom(c echo.Context) error {
	roomID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apiResponseError(c, http.StatusBadRequest, ErrBadRequest, err)
	}

	room, err := h.roomRepo.GetRoom(roomID)
	if err != nil {
		return apiResponseError(c, http.StatusInternalServerError, ErrInternalServerError, err)
	}

	return apiResponseOK(c, output.ToGetRoomResponse(room))
}
