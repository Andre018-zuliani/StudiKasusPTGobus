package handler

import (
    "fmt"
    "session-7/dto"
    "session-7/service"
)

type TicketHandler struct {
    TicketService service.TicketService
}

// NewTicketHandler adalah konstruktor untuk TicketHandler
func NewTicketHandler(s service.TicketService) TicketHandler {
    return TicketHandler{
        TicketService: s,
    }
}

// GetPrice memproses permintaan dan mencetak hasilnya ke konsol.
func (h *TicketHandler) GetPrice(req dto.TicketRequest) {
    response, err := h.TicketService.CheckPrice(req)

    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Format output sesuai dengan contoh di soal praktik
    fmt.Println("=== Harga Tiket ===")
    fmt.Printf("Penumpang: %s\n", response.PassengerName)
    fmt.Printf("Tujuan   : %s\n", response.Destination)
    // Menggunakan format %.2f untuk mencetak dua angka di belakang koma (contoh: Rp 150000.00)
    fmt.Printf("Harga    : Rp %.2f\n", response.Price)
}

