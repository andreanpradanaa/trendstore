# 📦 Daftar API Trendstore

### 🛠️ Admin Panel

| Method | Endpoint             | Deskripsi                             |
|--------|----------------------|----------------------------------------|
| GET    | /admin/users         | List semua user                        |
| GET    | /admin/users/:id     | Lihat detail user tertentu             |
| GET    | /admin/orders        | List semua pesanan                     |
| GET    | /admin/orders/:id    | Detail pesanan                         |
| GET    | /admin/products      | List semua produk (termasuk non-aktif) |
| POST   | /admin/products      | Tambah produk baru                     |
| PUT    | /admin/products/:id  | Update produk                          |
| DELETE | /admin/products/:id  | Hapus produk                           |

> 🔐 Semua endpoint `/admin/**` hanya bisa diakses oleh user dengan role `admin`.

---

### 🔐 Auth
| Method | Endpoint        | Deskripsi                     |
|--------|-----------------|-------------------------------|
| POST   | /auth/login     | Login user                    |
| POST   | /auth/register  | Register akun baru            |
| POST   | /auth/logout    | Logout dan revoke token       |

---

### 👤 User

| Method | Endpoint       | Deskripsi                        |
|--------|----------------|----------------------------------|
| GET    | /users/me      | Ambil data profil user saat ini |
| PUT    | /users/me      | Update profil user              |
| GET    | /users/:id     | Detail user berdasarkan ID      |

---

### 🛍️ Produk

| Method | Endpoint        | Deskripsi                        |
|--------|-----------------|----------------------------------|
| GET    | /products       | List semua produk                |
| GET    | /products/:id   | Detail produk                    |
| POST   | /products       | Tambah produk baru (admin only) |
| PUT    | /products/:id   | Edit produk (admin only)        |
| DELETE | /products/:id   | Hapus produk (admin only)       |

---

### 🧺 Keranjang (Cart)

| Method | Endpoint         | Deskripsi                        |
|--------|------------------|----------------------------------|
| GET    | /cart            | Lihat isi keranjang              |
| POST   | /cart            | Tambah item ke keranjang        |
| PUT    | /cart/:itemId    | Update jumlah item              |
| DELETE | /cart/:itemId    | Hapus item dari keranjang       |

---

### 📦 Checkout & Order

| Method | Endpoint        | Deskripsi                         |
|--------|-----------------|-----------------------------------|
| POST   | /checkout       | Proses checkout keranjang         |
| GET    | /orders         | Lihat semua order user            |
| GET    | /orders/:id     | Detail order                      |

---

### 💳 Pembayaran

| Method | Endpoint           | Deskripsi                    |
|--------|--------------------|------------------------------|
| POST   | /payments/pay      | Proses pembayaran            |
| GET    | /payments/status   | Cek status pembayaran        |

---

### 🚚 Pengiriman (Shipping)

| Method | Endpoint                | Deskripsi                           |
|--------|-------------------------|-------------------------------------|
| GET    | /shipping/track/:id     | Lacak status pengiriman berdasarkan order ID |

---

### 🎁 Promosi (Voucher / Diskon)

| Method | Endpoint         | Deskripsi                             |
|--------|------------------|----------------------------------------|
| GET    | /promos          | List semua promo aktif                 |
| GET    | /promos/:code    | Cek detail promo berdasarkan kode      |
| POST   | /promos/apply    | Apply kode promo ke keranjang          |

---

### 💖 Wishlist

| Method | Endpoint              | Deskripsi                            |
|--------|-----------------------|--------------------------------------|
| GET    | /wishlist             | Lihat semua produk di wishlist       |
| POST   | /wishlist             | Tambahkan produk ke wishlist         |
| DELETE | /wishlist/:productId  | Hapus produk dari wishlist           |

---

### 🌟 Rating & Review

| Method | Endpoint                | Deskripsi                              |
|--------|-------------------------|----------------------------------------|
| POST   | /products/:id/reviews   | Beri rating & ulasan ke produk         |
| GET    | /products/:id/reviews   | Lihat semua ulasan untuk produk        |

---

### 🔔 Notifikasi

| Method | Endpoint           | Deskripsi                             |
|--------|--------------------|----------------------------------------|
| GET    | /notifications     | Lihat semua notifikasi user           |
| PUT    | /notifications/:id | Tandai notifikasi sebagai dibaca      |
