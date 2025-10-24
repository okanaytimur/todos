# 🧩 TodoX – Go + SQLite REST API

Basit, hafif ve Docker içinde çalışan bir **Go tabanlı Todo servisi**.  
SQLite veritabanı kullanır ve REST API üzerinden CRUD işlemleri yapılabilir.

---

## 🚀 Özellikler

- Go ile yazılmış minimalist REST API  
- SQLite veritabanı (otomatik oluşturulur)  
- CRUD işlemleri:
  - `GET /todos` → Tüm görevleri listele  
  - `POST /todos` → Yeni görev ekle  
  - `PUT /todos/{id}` → Görevi güncelle  
  - `DELETE /todos/{id}` → Görevi sil  
- Docker ile kolay kurulum  
- Kalıcı veri için volume (host makinedeki `/data` klasörü)

---

## 🧱 Gereksinimler

- [Docker](https://www.docker.com/)  
- (İsteğe bağlı) `curl` veya Postman gibi bir API istemcisi  

---

## 🧰 Kurulum

### 1️⃣ Kaynak kodu klonla

```bash
git clone https://github.com/<kullanici-adi>/todox.git
cd todox
2️⃣ Docker imajını oluştur
bash
Kodu kopyala
docker build -t todox .
3️⃣ Çalıştır (Windows örneği)
CMD kullanıyorsan:

bash
Kodu kopyala
docker run -d -p 8080:8080 -v C:\Users\<kullanici-adi>\todox\data:/app/data todox
PowerShell kullanıyorsan:

powershell
Kodu kopyala
docker run -d -p 8080:8080 -v "$PWD\data:/app/data" todox
Linux / macOS:

bash
Kodu kopyala
docker run -d -p 8080:8080 -v $(pwd)/data:/app/data todox
🧩 API Kullanımı
🔹 Tüm görevleri listele
bash
Kodu kopyala
curl http://localhost:8080/todos
🔹 Yeni görev ekle
Windows CMD’de " karakterleri dikkatli kullanılmalı!

bash
Kodu kopyala
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d "{\"title\":\"Deneme görevi\",\"details\":\"Test açıklaması\"}"
🔹 Görevi güncelle
bash
Kodu kopyala
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d "{\"title\":\"Yeni başlık\",\"details\":\"Güncellendi\"}"
🔹 Görevi sil
bash
Kodu kopyala
curl -X DELETE http://localhost:8080/todos/1
🗃️ Dosya Yapısı
bash
Kodu kopyala
.
├── main.go          # Uygulama ana dosyası
├── go.mod           # Go modül tanımı
├── Dockerfile       # Docker imajı oluşturmak için
└── data/            # SQLite veritabanı burada saklanır
🧑‍💻 Geliştirici Notları
Uygulama 8080 portunda çalışır.

Veritabanı /app/data/todos.db konumundadır.

Container silinse bile -v parametresiyle volume bağlı olduğu sürece veriler korunur.

📜 Lisans
Bu proje GPLv3 lisansı altında paylaşılmıştır.
Dilediğiniz gibi kullanabilir, değiştirebilir ve geliştirebilirsiniz.

✍️ Hazırlayan: Okan Aytimur
🌐 okanaytimur.com.tr
🕓 Son güncelleme: 24 Ekim 2025 (İstanbul)
