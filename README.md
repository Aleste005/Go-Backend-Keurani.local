# 📦 Go-Backend-Keurani.local

Sebuah proyek backend REST API yang dibangun menggunakan **GoLANG** versi `go1.24.4 windows/amd64`.  
Struktur proyek ini mengikuti arsitektur **MVC (Model - View - Controller)**, sehingga sangat cocok dan mudah dipahami oleh kalangan **pemula hingga menengah**.

---

## 🧠 Teknologi dan Tools
- 🔧 Bahasa Pemrograman: Go (Golang)
- 🌐 Web Framework: Gorilla Mux
- 🛢️ Database: MySQL
- 🧩 Arsitektur: MVC (Model-View-Controller)
- 🧱 Template Engine: `gohtml` (jika dibutuhkan)

---

## 🗂️ Struktur Folder

| Folder/File      | Deskripsi                                                                 |
|------------------|--------------------------------------------------------------------------|
| `config/`        | Konfigurasi global seperti koneksi ke database (`db.go`)                |
| `controller/`    | Logic utama untuk menangani HTTP request dan response                   |
| `middleware/`    | Middleware tambahan (CORS, Logger)                                       |
| `model/`         | Struktur data dan fungsi-fungsi SQL (CRUD)                              |
| `router/`        | Routing aplikasi (menggunakan `Gorilla Mux`)                            |
| `utils/`         | Fungsi pembantu seperti formatter response                              |
| `view/`          | Template HTML (`.gohtml`) jika menggunakan tampilan UI                  |
| `.env`           | Menyimpan variabel lingkungan (jangan upload ke publik)                 |
| `.gitignore`     | Menentukan file/folder yang tidak diikutsertakan dalam Git              |
| `go.mod`         | File deklarasi modul Go dan dependensi                                  |
| `main.go`        | Entry point utama aplikasi                                               |


---

## 🙏 Terima Kasih

Terima kasih telah mengunjungi dan menggunakan proyek ini.  
Semoga struktur dan dokumentasi yang disediakan dapat membantu kamu memahami serta membangun REST API dengan GoLang secara lebih mudah dan menyenangkan.

Jika proyek ini bermanfaat, jangan lupa untuk ⭐️ repository ini dan bagikan ke rekan lainnya!

Salam hangat,  
**Almuhajir,EstE**


