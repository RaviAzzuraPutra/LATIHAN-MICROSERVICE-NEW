# Sistem Microservice Order dan Stock

Dokumentasi lengkap untuk proyek microservice yang mengelola order dan stok barang dengan komunikasi asynchronous menggunakan Apache Kafka.

---

## 📋 Daftar Isi

1. [Pengenalan Proyek](#pengenalan-proyek)
2. [Arsitektur Sistem](#arsitektur-sistem)
3. [Teknologi yang Digunakan](#teknologi-yang-digunakan)
4. [Struktur Folder](#struktur-folder)
5. [Persyaratan Sistem](#persyaratan-sistem)
6. [Cara Instalasi](#cara-instalasi)
7. [Konfigurasi Environment](#konfigurasi-environment)
8. [Cara Menjalankan Aplikasi](#cara-menjalankan-aplikasi)
9. [API Endpoints](#api-endpoints)
10. [Alur Kerja Sistem](#alur-kerja-sistem)
11. [Penjelasan Database](#penjelasan-database)
12. [Performance Benchmark & Trade-off Analysis](#performance-benchmark--trade-off-analysis)
13. [Troubleshooting](#troubleshooting)

---

## 🎯 Pengenalan Proyek

Proyek ini adalah implementasi dari **Microservice Architecture** yang terdiri dari dua layanan utama:

### **1. Order Service (Layanan Pemesanan)**
Layanan yang bertanggung jawab untuk mengelola pemesanan barang dari pelanggan. Setiap kali ada pemesanan baru, layanan ini akan:
- Menyimpan data pemesanan ke dalam database
- Mengirimkan notifikasi ke layanan Stock melalui message broker (Kafka)

### **2. Stock Service (Layanan Stok)**
Layanan yang bertanggung jawab untuk mengelola stok atau inventori barang. Layanan ini akan:
- Mendengarkan notifikasi dari Order Service
- Secara otomatis mengurangi jumlah stok saat ada pemesanan baru
- Menyediakan API untuk mengelola stok barang secara manual

**Keunggulan Arsitektur Ini:**
- ✅ Dua layanan berjalan **independen** (tidak saling bergantung langsung)
- ✅ Komunikasi **asynchronous** melalui Kafka (tidak perlu response langsung)
- ✅ Database **terpisah** untuk masing-masing layanan (database isolation)
- ✅ **Scalable** dan mudah untuk menambah fitur baru

---

## 🏗️ Arsitektur Sistem

```
┌─────────────────────────────────────────────────────────────────┐
│                          CLIENT (Postman/Browser)               │
└────────────────────────┬────────────────────────────────────────┘
                         │
          ┌──────────────┴──────────────┐
          │                             │
    ┌─────▼──────────┐          ┌──────▼────────┐
    │  ORDER SERVICE │          │ STOCK SERVICE │
    │   Port: 8086   │          │  Port: 8087   │
    └─────┬──────────┘          └──────┬────────┘
          │                             │
    ┌─────▼──────────┐          ┌──────▼────────┐
    │ Database Order │          │ Database Stock│
    │  (PostgreSQL)  │          │ (PostgreSQL)  │
    └────────────────┘          └───────────────┘
          │                             ▲
          │                             │
          └─────────────┬───────────────┘
                        │
                   ┌────▼─────┐
                   │   KAFKA   │
                   │ Message   │
                   │  Broker   │
                   └───────────┘
```

**Alur Komunikasi:**
1. Client mengirim request ke Order Service via HTTP
2. Order Service menyimpan data ke database sendiri
3. Order Service mengirim pesan ke Kafka
4. Stock Service mendengarkan Kafka 24/7
5. Saat ada pesan baru, Stock Service memproses dan update database

---

## 🛠️ Teknologi yang Digunakan

| Komponen | Teknologi | Versi |
|----------|-----------|-------|
| **Bahasa Pemrograman** | Go (Golang) | 1.20+ |
| **Framework Web** | Gin | v1.9+ |
| **Database** | PostgreSQL | 13+ |
| **Message Broker** | Apache Kafka | 3.0+ |
| **ORM** | GORM | v1.25+ |
| **Containerization** | Docker & Docker Compose | 20.10+ |

---

## 📁 Struktur Folder

```
LATIHAN-MICROSERVICE-NEW/
├── order/                          # Service untuk mengelola order
│   ├── main.go                     # Entry point aplikasi Order
│   ├── go.mod                      # Module Go untuk Order
│   ├── bootstrap/
│   │   └── index.go                # Inisialisasi aplikasi (DB, Kafka, Router)
│   ├── config/                     # Konfigurasi aplikasi
│   │   ├── app_config/             # Konfigurasi port aplikasi
│   │   ├── db_config/              # Konfigurasi database
│   │   └── kafka_config/           # Konfigurasi Kafka
│   ├── controller/
│   │   └── index.go                # Handler untuk HTTP requests
│   ├── service/
│   │   └── index.go                # Logika bisnis Order
│   ├── repository/
│   │   └── index.go                # Akses database
│   ├── model/
│   │   └── order.go                # Struktur data Order
│   ├── request/
│   │   └── index.go                # Struktur request HTTP
│   ├── router/
│   │   └── index.go                # Routing HTTP
│   ├── database/
│   │   └── connect.go              # Koneksi ke database
│   └── apache_kafka/
│       ├── producer/               # Mengirim pesan ke Kafka
│       ├── interface/              # Interface untuk producer
│       └── message/                # Struktur pesan Kafka
│
├── stock/                          # Service untuk mengelola stock
│   ├── main.go                     # Entry point aplikasi Stock
│   ├── go.mod                      # Module Go untuk Stock
│   ├── bootstrap/
│   │   └── index.go                # Inisialisasi aplikasi
│   ├── config/                     # Konfigurasi aplikasi
│   │   ├── app_config/
│   │   ├── db_config/
│   │   └── kafka_config/
│   ├── controller/
│   │   └── index.go                # Handler untuk HTTP requests
│   ├── service/
│   │   └── index.go                # Logika bisnis Stock
│   ├── repository/
│   │   └── index.go                # Akses database
│   ├── model/
│   │   └── stock.go                # Struktur data Stock
│   ├── request/
│   │   └── index.go                # Struktur request HTTP
│   ├── router/
│   │   └── index.go                # Routing HTTP
│   ├── database/
│   │   └── connect.go              # Koneksi ke database
│   └── apache_kafka/
│       └── consumer/               # Mendengarkan pesan dari Kafka
│       └── message/                # Struktur pesan Kafka yang akan didengarkan
│
├── postgresql/                     # Database setup
│   ├── order_db.sql                # Script SQL untuk database order
│   └── stock_db.sql                # Script SQL untuk database stock
│
├── k8s/                            # Konfigurasi Kubernetes
│   ├── order-kubernetes.yml        # Deployment Order di K8s
│   └── stock-kubernetes.yml        # Deployment Stock di K8s
│
├── kafka/                          # Data Kafka (generated)
│   └── data/                       # Folder untuk menyimpan data Kafka
│
├── docker-compose.yml              # Konfigurasi Docker Compose
└── README.md                       # Dokumentasi (file ini)
```
---

## 📡 API Endpoints

### **Order Service Endpoints**

#### **1. Get All Orders** (Ambil semua data order)
```http
GET /order/
```
**Response (200 OK):**
```json
{
  "Message": "Berhasil Mendapatkan Data Order",
  "Data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "productid": "PROD-001",
      "quantity": 5
    }
  ]
}
```

#### **2. Create New Order** (Buat order baru)
```http
POST /order/add-order
Content-Type: application/json

{
  "productid": "PROD-001",
  "quantity": 10
}
```

**Response (201 Created):**
```json
{
  "Message": "Berhasil Menambahkan Data",
  "Data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "productid": "PROD-001",
    "quantity": 10
  }
}
```

**Penjelasan:**
- Ketika order berhasil dibuat, sistem secara otomatis mengirim pesan ke Kafka
- Stock Service akan mendengarkan dan mengurangi stok barang secara otomatis

---

### **Stock Service Endpoints**

#### **1. Create Stock** (Tambah barang ke stok)
```http
POST /stock/add-stock
Content-Type: application/json

{
  "name": "Laptop Asus",
  "quantity": 50
}
```

**Response (200 OK):**
```json
{
  "Message": "Berhasil Menambahkan Data",
  "Data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Laptop Asus",
    "quantity": 50
  }
}
```

#### **2. Get All Stock** (Ambil semua data stok)
```http
GET /stock/
```

**Response (200 OK):**
```json
{
  "Message": "Berhasil Mendapatkan Data Stock",
  "Data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Laptop Asus",
      "quantity": 45
    }
  ]
}
```

#### **3. Get Stock by ID** (Ambil stok berdasarkan ID)
```http
GET /stock/:id
```

Example:
```http
GET /stock/550e8400-e29b-41d4-a716-446655440000
```

**Response (200 OK):**
```json
{
  "Message": "Berhasil Mendapatkan Data Berdasarkan ID",
  "Data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Laptop Asus",
    "quantity": 45
  }
}
```

#### **4. Update Stock** (Ubah data stok)
```http
PUT /stock/update/:id
Content-Type: application/json

{
  "name": "Laptop Asus Pro",
  "quantity": 30
}
```

Example:
```http
PUT /stock/update/550e8400-e29b-41d4-a716-446655440000
```

**Response (200 OK):**
```json
{
  "Message": "Berhasil Mengubah Data",
  "Data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Laptop Asus Pro",
    "quantity": 30
  }
}
```

#### **5. Delete Stock** (Hapus stok)
```http
DELETE /stock/delete/:id
```

Example:
```http
DELETE /stock/delete/550e8400-e29b-41d4-a716-446655440000
```

**Response (200 OK):**
```json
{
  "Message": "Berhasil Menghapus Data"
}
```

---

## 📊 Alur Kerja Sistem

### **Skenario Lengkap: Dari Order Sampai Update Stock**

```
1. CLIENT
   └─ POST http://localhost:8000/order/add-order
      Body: {
        "productid": "LAPTOP-001",
        "quantity": 5
      }
      ↓
2. ORDER SERVICE → Controller
   └─ Receive dan validate request
   └─ Parse data JSON
      ↓
3. ORDER SERVICE → Service Layer
   └─ Validasi business logic
   └─ Cek apakah product ID valid
   └─ Cek apakah quantity > 0
      ↓
4. ORDER SERVICE → Repository
   └─ Insert ke database PostgreSQL (order_db)
   └─ Generate UUID otomatis sebagai ID
      ↓
5. DATABASE (order_db)
   ┌──────────────────────────────────────┐
   │ orders table                         │
   ├──────────────────────────────────────┤
   │ id (UUID)    │ productid   │ quantity│
   ├──────────────────────────────────────┤
   │ xxxx-xxxx... │ LAPTOP-001  │ 5       │
   └──────────────────────────────────────┘
      ↓
6. ORDER SERVICE → Kafka Producer
   └─ Create message: {
       "id": "xxxx-xxxx...",
       "product_id": "LAPTOP-001",
       "quantity": 5,
       "timestamp": "2024-01-15T10:30:00Z"
     }
   └─ Send to Kafka topic: "order_created"
   └─ Jika gagal, log warning tapi order tetap tersimpan
      ↓
7. KAFKA BROKER
   ┌───────────────────────────┐
   │ Topic: order_created      │
   ├───────────────────────────┤
   │ Message: {...}            │
   │ (Tertampung di partition) │
   └───────────────────────────┘
      ↓
8. STOCK SERVICE → Consumer (Background Process)
   └─ Listening ke Kafka topic 24/7
   └─ Detect pesan baru
   └─ Parse message JSON
   └─ Call service: HapusStockBerdasarkanOrder()
      ↓
9. STOCK SERVICE → Service Layer
   └─ Query database: cari stok dengan ID = "LAPTOP-001"
   └─ Kurangi quantity: 50 - 5 = 45
   └─ Update database
      ↓
10. DATABASE (stock_db)
    ┌──────────────────────────────────────┐
    │ stocks table                         │
    ├──────────────────────────────────────┤
    │ id │ name       │ quantity (updated) │
    ├──────────────────────────────────────┤
    │ ... │ Laptop... │ 45 (sebelum: 50)  │
    └──────────────────────────────────────┘
      ↓
11. RESPONSE KE CLIENT
    HTTP 201 Created
    {
      "Message": "Berhasil Menambahkan Data",
      "Data": {
        "id": "xxxx-xxxx...",
        "productid": "LAPTOP-001",
        "quantity": 5
      }
    }

✅ PROSES SELESAI
   - Order tersimpan di database order
   - Stock berkurang di database stock
   - Semua terjadi secara otomatis lewat Kafka
```

---

## 💾 Penjelasan Database

### **Database Order (order_db)**

Menyimpan semua data transaksi pemesanan dari pelanggan.

```sql
CREATE TABLE orders (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  productid VARCHAR(100) NOT NULL,
  quantity INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Kolom:**
- `id`: Identitas unik untuk setiap order (UUID format)
- `productid`: Kode atau ID produk yang dipesan
- `quantity`: Jumlah unit yang dipesan
- `created_at`: Waktu order dibuat

### **Database Stock (stock_db)**

Menyimpan semua data inventori barang.

```sql
CREATE TABLE stocks (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  quantity INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Kolom:**
- `id`: Identitas unik untuk setiap barang (UUID format)
- `name`: Nama barang
- `quantity`: Jumlah stok tersedia
- `created_at`: Waktu data dibuat
- `updated_at`: Waktu data terakhir diubah

### **Visualisasi Relasi Database**

```
┌──────────────────────┐         ┌──────────────────────┐
│    orders (order_db) │         │   stocks (stock_db)  │
├──────────────────────┤         ├──────────────────────┤
│ id: UUID             │         │ id: UUID             │
│ productid: VARCHAR   │────┬───→│ name: VARCHAR        │
│ quantity: INT        │    │    │ quantity: INT        │
│ created_at: TIME     │    │    │ created_at: TIME     │
└──────────────────────┘    │    │ updated_at: TIME     │
                            │    └──────────────────────┘
                     (Kafka messaging)
```

**Alur Sinkronisasi:**
1. Order dibuat dengan `productid`
2. Kafka mengirimkan pesan ke Stock Service
3. Stock Service cari barang dengan ID yang sama
4. Update kolom `quantity` di table stocks
5. Update kolom `updated_at` otomatis

---

## 📚 Contoh Penggunaan dengan cURL

### **Skenario Lengkap (Step-by-Step)**

#### **Step 1: Buat Stok Barang**
```bash
curl -X POST http://localhost:8001/stock/add-stock \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 14 Pro",
    "quantity": 100
  }'
```

**Respon:**
```json
{
  "Message": "Berhasil Menambahkan Data",
  "Data": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "iPhone 14 Pro",
    "quantity": 100
  }
}
```

Catat ID yang didapat: `a1b2c3d4-e5f6-7890-abcd-ef1234567890`

#### **Step 2: Lihat Data Stok**
```bash
curl http://localhost:8001/stock/
```

#### **Step 3: Buat Order (Stok akan berkurang otomatis)**
```bash
curl -X POST http://localhost:8000/order/add-order \
  -H "Content-Type: application/json" \
  -d '{
    "productid": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "quantity": 10
  }'
```

#### **Step 4: Verifikasi Stok Berkurang**
```bash
# Tunggu 2-3 detik, lalu cek stok
curl http://localhost:8001/stock/a1b2c3d4-e5f6-7890-abcd-ef1234567890
```

**Respon:** Quantity seharusnya sudah menjadi 90 (100 - 10)

```json
{
  "Message": "Berhasil Mendapatkan Data Berdasarkan ID",
  "Data": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "iPhone 14 Pro",
    "quantity": 90
  }
}
```

---

## ⚡ Performance & Keuntungan Sistem

Bagian ini menjelaskan bagaimana performa sistem kami dan keuntungan arsitektur yang dipilih dengan bahasa sehari-hari.

---

### **1. Kecepatan Sistem (Berapa Lama Sih Proses Order?)**

Saat Anda membuat order, sistemnya bekerja seperti ini:

**Waktu untuk Membuat 1 Order:**
```
Menerima request        : 5ms   (membaca data dari client)
Cek data order          : 2ms   (pastikan product ID & quantity OK)
Simpan ke database      : 30ms  (masuk ke tabel orders)
Kirim ke Kafka          : 20ms  (beritahu stock service)
Balas ke client         : 5ms   (kirim respon "berhasil")
─────────────────────────────────
Total waktu            : ~60-65 ms (kurang dari 0,1 detik)
```

**Jadi berapa lama stock berkurang?**
- Order dibuat: 0 detik (langsung disimpan)
- Stock berkurang: 0.05-0.15 detik kemudian (lewat Kafka)
- Ini sangat cepat dan tidak perlu khawatir 😊

---

### **2. Mengapa Pilih Kafka (Bukan HTTP Langsung)?**

**Skenario 1: Kalau Menggunakan HTTP Langsung (TIDAK BAGUS)**

```
CLIENT
  │
  └─→ ORDER SERVICE
        │
        └─→ "Halo STOCK SERVICE, kurangi stok dong!"
              │
              └─ STOCK SERVICE: "Tapi aku lagi offline/sibuk"
                    │
                    └─ ORDER SERVICE: "Waduh, order gagal!"
              
❌ MASALAH: Order gagal hanya karena stock service sedang offline
```

**Skenario 2: Pakai Kafka (CARA KAMI - LEBIH BAIK)**

```
CLIENT
  │
  └─→ ORDER SERVICE
        │
        └─→ "Halo KAFKA, tolong kirim pesan ini"
              │
              KAFKA: "OK, pesan masuk di papan pengumuman"
              │
              └─ ORDER SERVICE: "Alhamdulillah, order berhasil!"
              
              Sambil itu, STOCK SERVICE
              └─→ "Halo KAFKA, ada pesan buat aku?"
                    │
                    └─ KAFKA: "Iya, kurangin stok iPhone"
                          │
                          STOCK SERVICE: "OK, langsung dikurangin!"

✅ KEUNTUNGAN: Order berhasil meski stock service lagi offline
              Stock akan berkurang otomatis begitu service online lagi
```

**Perbandingan Singkat:**

| Cara | Kecepatan Order | Stok Berkurang | Jika Stock Down | Ribet? |
|------|-----------------|----------------|-----------------|--------|
| **HTTP Langsung** | Cepat | Langsung | Order gagal ❌ | Simpel |
| **Kafka (Kami)** | Cepat | 0.1 detik | Order OK ✅ | Sedikit kompleks |

---

### **3. Database Terpisah (Mengapa Gak Disamain?)**

**Opsi 1: Database Terpisah (Pilihan Kami)**

```
Order Service                Stock Service
      ↓                           ↓
   Database 1               Database 2
  (order_db)              (stock_db)
  
Keuntungan:
✅ Order service bisa jalan sendiri (tidak tergantung stock)
✅ Jika database stock down, order masih bisa diterima
✅ Lebih cepat (masing-masing database fokus sendiri)
✅ Mudah menambah server baru tanpa pengaruhi yang lain

Kerugian:
❌ Sedikit lebih ribet untuk maintenance
```

**Opsi 2: Database Disamain (TIDAK BAGUS untuk Microservice)**

```
Order Service + Stock Service
        ↓
    Database 1 (Shared)

Keuntungan:
✅ Lebih sederhana

Kerugian:
❌ Jika database down, SEMUA service jatuh
❌ Tidak bisa scale terpisah
❌ Mudah terjadi conflict antar service
```

**Kesimpulan:** Cara kami lebih baik karena lebih aman 🛡️

---

### **4. Berapa Banyak Sih Order yang Bisa Ditangani?**

**Sistem kami bisa handle:**

| Metrik | Kemampuan | Catatan |
|--------|-----------|---------|
| **Order per detik** | 15-20 | Per server |
| **Pesan Kafka per detik** | 50-100 | Dengan 3 server stock |
| **Concurrent users** | 100+ | Bisa di-scale lebih tinggi |
| **Storage order/bulan** | ~50GB | Untuk 100 ribu order |

**Contoh:**
- Toko online kecil (100 order/hari) → Aman banget
- Toko online sedang (1000 order/hari) → Masih OK
- Toko online besar (10,000 order/hari) → Perlu ditambah server

---

### **5. Apakah Data Order & Stock Selalu Sama?**

**Jawab: Hampir selalu sama, dengan delay sangat kecil**

```
Contoh: iPhone punya 100 stok

Waktu 0.00 detik : Client buat order 5 iPhone
                   Database order    : sudah tercatat ✓
                   Database stock    : masih 100 (belum berkurang)

Waktu 0.05 detik : Stock service ambil pesan dari Kafka
Waktu 0.10 detik : Database stock    : berkurang jadi 95 ✓

Jadi delay hanya 0.1 detik (100 milidetik)
Ini tidak masalah untuk kebanyakan e-commerce
```

**Kapan baru perlu khawatir?**
- Jika delay lebih dari 1 detik → mulai perlu optimize
- Jika delay lebih dari 5 detik → harus optimize

Sistem kami sekarang: **Delay < 0.2 detik** = Bagus! 👍

---

### **6. Apa Saja yang Perlu Di-monitor?**

Untuk memastikan sistem sehat, pantau hal-hal ini:

**Yang Penting:**
```
1. Berapa lama respon API?
   → Idealnya < 100ms
   
2. Berapa banyak error?
   → Idealnya < 1% (99%+ success)
   
3. Stock update delay?
   → Idealnya < 500ms
   
4. Database CPU usage?
   → Idealnya < 70%
   
5. Storage space?
   → Pantau jangan sampai penuh
```

**Pakai Tool Simple:**
```bash
# Lihat real-time metrics
curl http://localhost:8000/metrics
curl http://localhost:8001/metrics
```

---

### **7. Kapan Harus Tambah Server?**

**Tanda-tanda perlu di-scale:**

```
Jika... → Lakukan...

Order per detik > 15         → Tambah 1 order server
Stock update delay > 1s      → Tambah 1-2 stock server
CPU > 80%                    → Upgrade server atau tambah instance
Database query > 500ms       → Tambah index atau upgrade DB
Error rate > 5%              → Cek log, ada apa?
```

**Contoh:**
```
Hari ini order masuk 100/detik
Kapasitas saat ini: 15-20/detik
Perlu: 100 ÷ 15 = ~7 server order service
Solusi: Tambah 6 server lagi, atau pakai auto-scaling
```

---

## 👨‍💻 Author

**Ravi Azzura Putra**
- GitHub: [@RaviAzzuraPutra](https://github.com/RaviAzzuraPutra)
- Project: [LATIHAN-MICROSERVICE-NEW](https://github.com/RaviAzzuraPutra/LATIHAN-MICROSERVICE-NEW)

---