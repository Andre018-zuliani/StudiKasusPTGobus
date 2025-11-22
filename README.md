# Gobus CLI

Aplikasi CLI sederhana untuk menampilkan harga tiket berdasarkan destinasi.

## Struktur proyek

- [main.go](main.go) — entry point aplikasi
- [go.mod](go.mod) — modul Go
- [data/destinations.go](data/destinations.go) — data destinasi dan harga
- [dto/request.go](dto/request.go) — DTO untuk request
- [dto/response.go](dto/response.go) — DTO untuk response
- [model/ticket.go](model/ticket.go) — model Ticket
- [service/ticket.go](service/ticket.go) — logika bisnis untuk pengecekan harga
- [handler/ticket.go](handler/ticket.go) — handler yang mengeluarkan output ke konsol

## Cara menjalankan

Dari root proyek (folder `gobus-cli`) jalankan:

```sh
go run main.go
```

## Contoh output

Saat menjalankan contoh request di [main.go](main.go) akan menghasilkan:

```
=== Harga Tiket ===
Penumpang: sidik
Tujuan   : Jakarta
Harga    : Rp 150000.00
-------------------------
```

## Penjelasan singkat komponen

- Entry point: [main.go](main.go) — membuat service dan handler, lalu memanggil [`dto.NewRequest`](dto/request.go).
- Service:
  - Konstruktor: [`service.NewTicketService`](service/ticket.go)
  - Fungsi utama: [`service.TicketService.CheckPrice`](service/ticket.go) — mengambil harga dari [`data.Destinations`](data/destinations.go), mengembalikan [`dto.TicketResponse`](dto/response.go).
- Handler:
  - Konstruktor: [`handler.NewTicketHandler`](handler/ticket.go)
  - Menampilkan hasil: [`handler.TicketHandler.GetPrice`](handler/ticket.go)
- Model/DTO:
  - Model: [`model.Ticket`](model/ticket.go)
  - Request DTO: [`dto.TicketRequest`](dto/request.go), dibuat lewat [`dto.NewRequest`](dto/request.go)
  - Response DTO: [`dto.TicketResponse`](dto/response.go)

## Catatan

- Data destinasi disimpan sederhana di [data/destinations.go](data/destinations.go).
- Error jika destinasi tidak ditemukan dikembalikan oleh [`service.TicketService.CheckPrice`](service/ticket.go) dan ditangani oleh [`handler.TicketHandler.GetPrice`](handler/ticket.go).
