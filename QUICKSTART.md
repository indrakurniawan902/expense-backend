# 🚀 Expense API - Quick Start Guide

## Untuk Development

### Run Locally
```bash
# Install dependencies
go mod download

# Build
go build -o expense-api

# Run
./expense-api
```

Server akan jalan di `http://localhost:8080`

---

## Untuk Deployment (Deploy ke Cloud)

### Option 1: Railway.app (RECOMMENDED ⭐)

**Pros:**
- Paling mudah & cepat
- Free tier 5 projects
- Deploy hanya 5 menit
- Support Docker native
- Auto HTTPS

**Steps:**

1. **Push code ke GitHub**
   ```bash
   git init
   git add .
   git commit -m "Initial commit"
   git push -u origin main
   ```

2. **Setup Railway**
   - Go to https://railway.app
   - Sign up / Login
   - Click "New Project" → "Deploy from GitHub"
   - Select repository
   - Railway auto-detect & deploy!

3. **Get Public URL**
   - Cek di Railway Dashboard
   - URL akan seperti: `https://expense-api-prod.up.railway.app`

4. **Share ke Mobile Developer**
   ```
   Base URL: https://expense-api-prod.up.railway.app
   
   Endpoints:
   - GET    /health
   - POST   /expenses
   - GET    /expenses
   - GET    /expenses/:id
   - PUT    /expenses/:id
   - DELETE /expenses/:id
   ```

---

### Option 2: Render.com (Juga Mudah)

1. Connect GitHub account di https://render.com
2. Click "New +" → "Web Service"
3. Select repository
4. Environment: Docker
5. Click "Create Web Service"
6. Done! Public URL auto-generated

---

### Option 3: Fly.io (Bagus untuk Performance)

```bash
# Install CLI
brew install flyctl

# Login & Deploy
flyctl auth login
flyctl launch
flyctl deploy
```

---

### Option 4: Docker (untuk testing lokal atau VPS)

**Build Docker Image**
```bash
docker build -t expense-api:latest .
```

**Run Container**
```bash
docker run -p 8080:8080 expense-api:latest
```

**Atau dengan Docker Compose**
```bash
docker-compose up -d
```

---

### Option 5: VPS (DigitalOcean, AWS EC2, Linode)

1. Create VPS instance
2. SSH & setup
3. Clone repo
4. Run: `./expense-api` atau via Docker
5. Setup Nginx reverse proxy (optional)

Lihat `DEPLOYMENT.md` untuk detail lengkap!

---

## Testing Deployed API

Test dengan curl atau Postman:

### Health Check
```bash
curl https://your-url/health
```

### Create Expense
```bash
curl -X POST https://your-url/expenses \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Coffee",
    "amount": 5.5,
    "category": "Food",
    "date": "2025-11-26"
  }'
```

### List Expenses
```bash
curl https://your-url/expenses
```

---

## Environment Variables

Jika perlu custom port (default 8080):

```bash
# Lokal
PORT=3000 ./expense-api

# Di Railway/Render (set di dashboard settings)
PORT=8080
```

---

## Monitoring & Logs

**Railway:**
- Dashboard → Logs tab

**Render:**
- Dashboard → Logs tab

**Docker:**
```bash
docker logs <container-id>
```

---

## API Response Format

Semua response punya struktur:

```json
{
  "statusCode": 200,
  "message": "Success message",
  "data": { /* actual data */ }
}
```

Error response:
```json
{
  "statusCode": 400,
  "message": "Error description",
  "data": null
}
```

---

## Next Steps (untuk Production)

- [ ] Add database (PostgreSQL)
- [ ] Add authentication (JWT)
- [ ] Add rate limiting
- [ ] Setup monitoring (New Relic, DataDog)
- [ ] Add API documentation (Swagger)
- [ ] Setup CI/CD pipeline
- [ ] Add unit tests

---

## Support

Untuk pertanyaan lebih lanjut, lihat:
- `README.md` - API Documentation
- `DEPLOYMENT.md` - Deployment Guide Details
- `.env.example` - Environment Variables

Semoga berhasil! 🎉
