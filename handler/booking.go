package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ruchi-dhore-tw/ticket-booking-system/contract"
	"github.com/ruchi-dhore-tw/ticket-booking-system/model"
)

func BookTicketHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		fmt.Println("Ticket Booked")
		var bookingRequest contract.BookingRequest

		if err := context.ShouldBindBodyWith(&bookingRequest, binding.JSON); err != nil {
			fmt.Printf("Received Error %+v ", err)
			context.JSON(http.StatusCreated, contract.BookingResponse{
				Success: false,
				Errors:  []string{"Invalid request body!!"},
				Data:    nil,
			})
			return
		}

		ticket := model.Ticket{
			Id:      1,
			Catalog: bookingRequest.Catalog,
			Slot:    bookingRequest.Slot,
		}

		context.JSON(http.StatusCreated, contract.BookingResponse{
			Success: true,
			Errors:  []string{},
			Data: []model.Ticket{
				ticket,
			},
		})
	}
}
