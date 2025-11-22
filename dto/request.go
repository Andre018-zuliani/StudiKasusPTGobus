package dto

// TicketRequest merepresentasikan input dari pengguna
type TicketRequest struct {
    Name        string
    Destination string
}

// NewRequest adalah konstruktor untuk TicketRequest
func NewRequest(name string, destination string) TicketRequest {
    return TicketRequest{
        Name:        name,
        Destination: destination,
    }
}
