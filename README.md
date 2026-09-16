# URL Shortener Go

Backend URL shortener dengan Go, layered architecture, dan dependency inversion.

## Struktur

```text
.
├── cmd/api/                       # Entry point aplikasi
├── internal/
│   ├── domain/                    # Entity dan aturan inti domain
│   ├── usecase/                   # Business logic dan application service
│   ├── repository/                # Kontrak persistence dan implementasi
│   │   └── memory/                # Adapter storage in-memory untuk development
│   └── delivery/http/handler/     # HTTP handler dan routing
├── docs/development.md            # Panduan pengembangan
└── go.mod
```

## Menjalankan

```bash
go run ./cmd/api
```

Server berjalan di `http://localhost:8080`. Port dapat diubah dengan environment variable `PORT`.

## Contoh API

Membuat short URL:

```bash
curl -X POST http://localhost:8080/api/v1/urls \
  -H "Content-Type: application/json" \
  -d '{"target_url":"https://example.com/docs"}'
```

Buka short code yang dikembalikan melalui `GET /{short_code}`.

## Pengujian

```bash
go test ./...
```

Detail keputusan arsitektur, aturan dependency, dan alur kontribusi ada di [docs/development.md](docs/development.md).
