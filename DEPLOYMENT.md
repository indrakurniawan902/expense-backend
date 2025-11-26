# Deployment Guide - Expense API

Panduan lengkap untuk deploy Expense API ke berbagai platform agar bisa diakses dari luar.

## 1. Docker Setup (Test Locally)

### Build Docker Image
```bash
docker build -t expense-api:latest .
```

### Run dengan Docker
```bash
docker run -p 8080:8080 expense-api:latest
```

### Atau gunakan Docker Compose
```bash
docker-compose up -d
```

Akses API: `http://localhost:8080`

---

## 2. Deploy ke Railway.app (RECOMMENDED - Easiest)

Railway adalah platform paling mudah dan cepat. Gratis untuk 5 project pertama!

### Step 1: Persiapan
1. Create akun di https://railway.app
2. Install Railway CLI:
   ```bash
   npm install -g @railway/cli
   ```

### Step 2: Login & Deploy
```bash
# Login ke Railway
railway login

# Initialize project
railway init

# Deploy
railway up
```

Setelah deploy, Railway akan generate URL public seperti:
```
https://expense-api-production.up.railway.app
```

Mobile developer Anda bisa langsung akses:
```
https://expense-api-production.up.railway.app/health
```

### Configure Environment (jika diperlukan)
Di Railway dashboard, bisa set environment variables atau configure port.

---

## 3. Deploy ke Render.com

Render juga free dan mudah.

### Step 1: Setup
1. Push code ke GitHub
2. Create akun di https://render.com
3. Connect GitHub account

### Step 2: Create Web Service
1. Click "New +" → "Web Service"
2. Select repository
3. Configure:
   - **Name:** expense-api
   - **Environment:** Docker
   - **Build Command:** Leave default
   - **Start Command:** Leave default

### Step 3: Deploy
Click "Create Web Service" dan Render akan auto deploy.

Public URL: `https://expense-api-xxxxx.onrender.com`

---

## 4. Deploy ke Heroku (Classic - Paid)

### Step 1: Setup Heroku
```bash
# Install Heroku CLI
brew install heroku

# Login
heroku login

# Create app
heroku create expense-api
```

### Step 2: Deploy
```bash
# Push ke Heroku
git push heroku main

# View logs
heroku logs --tail
```

Public URL: `https://expense-api.herokuapp.com`

---

## 5. Deploy ke DigitalOcean App Platform

### Step 1: Prepare
1. Push code ke GitHub
2. Create akun di https://www.digitalocean.com

### Step 2: Create App
1. Go to App Platform
2. Click "Create App"
3. Select GitHub repository
4. Configure:
   - **Service Name:** expense-api
   - **Build Command:** Leave blank (auto-detect)
   - **Run Command:** Leave blank (Dockerfile defined)

### Step 3: Deploy
Click "Launch App" dan DigitalOcean akan deploy.

---

## 6. Deploy ke AWS EC2 (VPS)

### Step 1: Launch EC2 Instance
1. Create EC2 instance (t3.micro - free tier eligible)
2. Choose Ubuntu 22.04 AMI
3. Configure security group:
   - Allow port 80 (HTTP)
   - Allow port 443 (HTTPS)
   - Allow port 8080 (custom)

### Step 2: SSH & Setup
```bash
# SSH ke server
ssh -i your-key.pem ec2-user@your-ec2-ip

# Install Docker
sudo yum update -y
sudo yum install docker -y
sudo systemctl start docker

# Clone repo
git clone <your-repo-url>
cd expense-backend

# Build & Run
docker build -t expense-api .
docker run -d -p 8080:8080 --name expense-api expense-api:latest
```

### Step 3: Setup Nginx Reverse Proxy (Optional)
```bash
sudo yum install nginx -y
```

Edit `/etc/nginx/nginx.conf`:
```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

---

## 7. Deploy ke Fly.io

Fly.io sangat bagus untuk aplikasi dengan latency rendah.

### Step 1: Setup
```bash
# Install Fly CLI
brew install flyctl

# Login
flyctl auth login

# Create app
flyctl launch
```

### Step 2: Deploy
```bash
flyctl deploy
```

Public URL: `https://expense-api-xxxxx.fly.dev`

---

## Testing Deployed API

Setelah deploy, test endpoint dengan:

### Health Check
```bash
curl https://your-deployed-url/health
```

### Create Expense
```bash
curl -X POST https://your-deployed-url/expenses \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Test Expense",
    "amount": 100,
    "category": "Test",
    "date": "2025-11-26"
  }'
```

---

## Setup Domain Custom (Optional)

Jika mau punya domain custom (misal: api.mycompany.com):

### Dengan Railway
1. Settings → Domains
2. Add custom domain
3. Update DNS records

### Dengan Render
1. Settings → Custom Domain
2. Add domain
3. Update DNS records

---

## Mobile Developer Integration

Kirim ke mobile developer:

```
Base URL: https://your-deployed-url

Endpoints:
- GET    /health
- POST   /expenses
- GET    /expenses
- GET    /expenses/:id
- PUT    /expenses/:id
- DELETE /expenses/:id

Example:
https://your-deployed-url/expenses
```

Berikut response format yang akan diterima:
```json
{
  "statusCode": 200,
  "message": "All expenses retrieved successfully",
  "data": [...]
}
```

---

## Environment Variables (Jika diperlukan di masa depan)

Create `.env` file:
```
GIN_MODE=release
PORT=8080
```

Atau set di platform deployment masing-masing.

---

## Scaling & Production Considerations

1. **Database** - Saat ini in-memory. Upgrade ke PostgreSQL/MySQL
2. **Authentication** - Add JWT untuk secure endpoints
3. **Rate Limiting** - Protect API dari abuse
4. **HTTPS** - Semua platform di atas support HTTPS otomatis
5. **Monitoring** - Setup logging & monitoring
6. **Backup** - Setup data backup jika pakai database

---

## Recommendation untuk Anda

Untuk quick start dengan minimal setup:
1. **Railway.app** atau **Render.com** → Paling mudah & cepat
2. **Fly.io** → Bagus jika performance penting
3. **AWS/DigitalOcean** → Jika butuh kontrol penuh & scaling

Saya recommend **Railway.app** karena:
- ✅ Sangat mudah
- ✅ Free tier generous
- ✅ Deploy dalam 5 menit
- ✅ Support Docker native
