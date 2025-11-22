package model

// Ticket merepresentasikan entitas tiket inti, meskipun tidak di-persist (disimpan)
type Ticket struct {
    Destination string
    Price       float64
}
