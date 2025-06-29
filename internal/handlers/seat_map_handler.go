package handlers

import (
	"net/http"
	"strconv"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/handlers/responses"
	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/services"
	e "github.com/alifrahmadian/alif-bookcabin-coding-test/pkg/errors"
	"github.com/gin-gonic/gin"
)

type SeatMapHandler struct {
	SeatMapService services.SeatMapService
}

func NewSeatMapHandler(seatMapService *services.SeatMapService) *SeatMapHandler {
	return &SeatMapHandler{
		SeatMapService: *seatMapService,
	}
}

func (h *SeatMapHandler) GetSeatMap(c *gin.Context) {
	seatsItineraryPartIDParam := c.Param("id")

	id, err := strconv.ParseInt(seatsItineraryPartIDParam, 10, 64)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, e.ErrInvalidSeatsItineraryPartID.Error())
		return
	}

	resp, err := h.SeatMapService.GetSeatMapBySeatsItineraryPartID(id)
	if err != nil {
		if err == e.ErrSeatMapSeatsItineraryPartNotExist {
			responses.ErrorResponse(c, http.StatusBadRequest, e.ErrSeatMapSeatsItineraryPartNotExist.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.SuccessResponse(c, "success", resp)
}
