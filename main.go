package main

import (
    "fmt"
    "session-7/dto"
    "session-7/handler"
    "session-7/service"
)

func main() {
    // Inisialisasi Service dan Handler
    ticketService := service.NewTicketService()
    ticketHandler := handler.NewTicketHandler(ticketService)

    // Contoh Request 1: Sesuai dengan soal praktik
    req1 := dto.NewRequest("sidik", "Jakarta")
    ticketHandler.GetPrice(req1)

    fmt.Println("-------------------------")
}
