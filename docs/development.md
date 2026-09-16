# Panduan Pengembangan

## Prinsip arsitektur

Project ini menggunakan layered architecture dengan dependency rule berikut:

```text
HTTP handler -> use case -> repository interface
                                      ^
                                      |
                         memory/database adapter
```

- `domain` berisi entity dan aturan bisnis yang tidak bergantung pada framework.
- `usecase` mengorkestrasi business logic dan hanya bergantung pada interface.
- `repository` mendefinisikan kontrak akses data. Implementasi storage berada di subpackage adapter.
- `delivery/http` menerjemahkan request HTTP menjadi input use case dan hasil use case menjadi response HTTP.
- `cmd/api` hanya melakukan wiring dependency dan menjalankan server.

Package di bawah `internal` hanya dapat diakses oleh module ini. Ini membantu menjaga batas public API.

## Aturan dependency

1. Domain tidak boleh mengimpor HTTP, database driver, atau package delivery.
2. Use case tidak boleh mengetahui detail JSON, status code, atau concrete storage.
3. Handler tidak boleh mengakses repository secara langsung; gunakan use case.
4. Implementasi repository harus memenuhi interface yang didefinisikan oleh layer repository.
5. Constructor digunakan untuk dependency injection agar unit test dapat memakai fake repository.

## Alur menambah fitur

1. Definisikan atau ubah entity dan error bisnis di `internal/domain`.
2. Tambahkan kontrak repository bila fitur membutuhkan persistence baru.
3. Implementasikan business logic di `internal/usecase`.
4. Tambahkan adapter repository, misalnya `internal/repository/postgres`.
5. Tambahkan endpoint dan mapping request/response di `internal/delivery/http`.
6. Wire dependency baru di `cmd/api/main.go`.
7. Tambahkan unit test untuk use case dan handler test untuk kontrak HTTP.
8. Jalankan formatter, test, dan static checks sebelum membuka pull request.

## Perintah harian

```bash
gofmt -w .
go test ./...
go vet ./...
go run ./cmd/api
```

Gunakan `go test -race ./...` ketika mengubah kode yang menyentuh concurrency atau shared state.

## Konvensi kode

- Gunakan nama package pendek dan jelas.
- Return error, jangan panic untuk error yang dapat dipulihkan.
- Bungkus error dengan konteks bila layer membutuhkan informasi tambahan.
- Terima `context.Context` dari boundary dan teruskan ke dependency.
- Jangan menaruh konfigurasi, koneksi database, atau global mutable state di package domain.
- Perubahan API harus disertai dokumentasi dan test yang relevan.

## Langkah berikutnya untuk production

- Ganti `repository/memory` dengan adapter database yang menerapkan `repository.URLRepository`.
- Tambahkan konfigurasi terstruktur dan graceful shutdown.
- Tambahkan request ID, structured logging, metrics, dan health/readiness endpoint.
- Tambahkan validasi collision serta strategi retry untuk short code.
- Tambahkan migrasi database dan integration test dengan database terisolasi.
