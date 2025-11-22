package service

import (
	"fmt"
	"session-7/data"
	"session-7/dto"
	"session-7/model" // Import model
)

type TicketService struct{}

// NewTicketService adalah konstruktor untuk TicketService
func NewTicketService() TicketService {
	return TicketService{}
}

// CheckPrice mencari harga tiket dan mengembalikan Model.Ticket
func (s *TicketService) CheckPrice(req dto.TicketRequest) (dto.TicketResponse, error) {
	price, exists := data.Destinations[req.Destination]

	if !exists {
		// Jika destinasi tidak ditemukan
		return dto.TicketResponse{
			PassengerName: req.Name,
			Destination:   req.Destination,
		}, fmt.Errorf("destinasi '%s' tidak ditemukan", req.Destination)
	}

	// Model.Ticket (opsional: dapat dikonversi ke DTO di sini atau di Handler)
	ticketModel := model.Ticket{
		Destination: req.Destination,
		Price:       price,
	}

	// Mapping ke DTO Response untuk dikirim ke Handler
	response := dto.TicketResponse{
		PassengerName: req.Name,
		Destination:   ticketModel.Destination,
		Price:         ticketModel.Price,
	}

	return response, nil
}
