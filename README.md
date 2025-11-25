# Gator — RSS Aggregator CLI 🐊

Gator, RSS feed'lerini terminal üzerinden yönetmeni sağlayan bir **Go tabanlı komut satırı uygulamasıdır.**  
Kullanıcılar feed ekleyebilir, takip edebilir, RSS yazılarını toplayabilir ve kendi veritabanında görüntüleyebilir.

---

## 🚀 Gereksinimler

* Go 1.21+  
* PostgreSQL 15+  

---

## ⚙️ Kurulum

Gator'ı yüklemek için aşağıdaki komutu çalıştır:

```bash
go install github.com/USERNAME/blog-aggregator@latest
```

> `USERNAME` kısmını kendi GitHub kullanıcı adınla değiştir.

---

## 🗄️ Veritabanı & Yapılandırma

Postgres’te `gator` isminde bir veritabanı oluştur:
```bash
createdb gator
```

Config dosyası (`~/.gatorconfig.json`) şu formatta olmalı:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": "unknown"
}
```

---

## 💻 Kullanım

Tüm komutlar `gator` ile başlar.  
Aşağıda bazı örnekler mevcut:

```bash
# veritabanını sıfırla
gator reset

# kullanıcı oluştur
gator register kahya

# giriş yap
gator login kahya

# feed ekle
gator addfeed "Boot.dev Blog" "https://blog.boot.dev/index.xml"

# feed takip et
gator follow "https://news.ycombinator.com/rss"

# takip ettiklerini gör
gator following

# RSS verilerini her dakikada bir çek
gator agg 1m

# takip ettiğin feed'lerden yeni yazıları listele
gator browse 5
```

---

## 📜 Açıklama

Program, RSS kaynaklarını belirli aralıklarla (`agg` komutu) otomatik olarak toplar.  
Toplanan yazılar `posts` tablosuna kaydedilir ve `browse` komutu ile görüntülenebilir.

Tüm veriler PostgreSQL veritabanında saklanır.

---

## 🏗️ Geliştirme

Geliştirme aşamasında çalıştırmak için:

```bash
go run .
```

Yapıyı derleyip global kullanıma almak için:

```bash
go build
```
veya
```bash
go install
```

Bunu yaptıktan sonra program tüm sistemde şu şekilde çalışır:

```bash
gator <command>
```

---

## 🧑💻 Katkı

1. Depoyu klonla  
2. Yeni dal (`branch`) oluştur  
3. Değişikliklerini yap  
4. Pull Request gönder 🎉

---

## 📎 Lisans

MIT Lisansı.  
© 2025 — Gator CLI by [Senin Adın] 🐊
