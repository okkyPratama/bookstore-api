# Bookstore API

Bookstore API adalah sebuah aplikasi backend untuk mengelola data buku dan kategorinya. Aplikasi ini dibangun menggunakan Go dengan framework Gin dan menggunakan PostgreSQL sebagai database.

## Fitur

- Autentikasi pengguna menggunakan JWT
- Manajemen buku (CRUD operations)
- Manajemen kategori buku (CRUD operations)
- Pencarian buku berdasarkan kategori

## Prasyarat

Sebelum menjalankan aplikasi, pastikan Anda telah menginstal:

- Go (versi 1.16 atau lebih baru)
- PostgreSQL
- Git

## Instalasi

1. Clone repositori ini:
   ```
   git clone https://github.com/username/bookstore-api.git
   cd bookstore-api
   ```

2. Install dependensi:
   ```
   go mod tidy
   ```

3. Salin file `.env.example` menjadi `.env` dan sesuaikan konfigurasi database:
   ```
   cp .env.example .env
   ```

4. Jalankan migrasi database:
   ```
   go run main.go migrate
   ```

5. Jalankan aplikasi:
   ```
   go run main.go
   ```

Aplikasi akan berjalan di `http://localhost:8080`.

## Penggunaan API

### Autentikasi

Semua endpoint (kecuali registrasi dan login) memerlukan token JWT yang valid. Sertakan token di header `Authorization` dengan format `Bearer {token}`.

### Endpoints

#### User

- `POST /api/users/register`: Registrasi pengguna baru
- `POST /api/users/login`: Login pengguna

#### Buku

- `GET /api/books`: Mendapatkan semua buku
- `POST /api/books`: Menambahkan buku baru
- `GET /api/books/:id`: Mendapatkan detail buku
- `PUT /api/books/:id`: Mengupdate buku
- `DELETE /api/books/:id`: Menghapus buku

#### Kategori

- `GET /api/categories`: Mendapatkan semua kategori
- `POST /api/categories`: Menambahkan kategori baru
- `GET /api/categories/:id`: Mendapatkan detail kategori
- `PUT /api/categories/:id`: Mengupdate kategori
- `DELETE /api/categories/:id`: Menghapus kategori
- `GET /api/categories/:id/books`: Mendapatkan buku berdasarkan kategori

### Contoh Penggunaan

1. Registrasi Pengguna:
   ```
   POST /api/users/register
   {
     "username": "john_doe",
     "password": "securepassword123"
   }
   ```

2. Login:
   ```
   POST /api/users/login
   {
     "username": "john_doe",
     "password": "securepassword123"
   }
   ```

3. Menambahkan Buku Baru:
   ```
   POST /api/books
   Authorization: Bearer {your_jwt_token}
   {
     "title": "The Go Programming Language",
     "description": "A comprehensive guide to Go",
     "image_url": "https://example.com/go-book.jpg",
     "release_year": 2015,
     "price": 4999,
     "total_page": 380,
     "category_id": 1
   }
   ```

## Validasi

- Release year buku harus antara 1980 dan 2024.
- Thickness buku akan otomatis dihitung berdasarkan total halaman:
  - "tipis" jika kurang dari atau sama dengan 100 halaman
  - "tebal" jika lebih dari 100 halaman

## Deployment

Aplikasi ini siap untuk di-deploy ke Railway. Pastikan untuk mengatur variabel lingkungan yang diperlukan di dashboard Railway.

