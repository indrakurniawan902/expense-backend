#!/bin/bash

# Expense API - Deployment Quick Reference
# Copy & paste commands untuk deploy ke berbagai platform

echo "🚀 Expense API - Deployment Commands"
echo "====================================="
echo ""

# ============================================================================
# OPTION 1: RAILWAY.APP (RECOMMENDED)
# ============================================================================
echo "📌 OPTION 1: Railway.app (EASIEST - 5 minutes)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "1. Create account: https://railway.app"
echo ""
echo "2. Install Railway CLI:"
echo "   npm install -g @railway/cli"
echo ""
echo "3. Login & Deploy:"
echo "   railway login"
echo "   railway init"
echo "   railway up"
echo ""
echo "4. Get public URL from Railway dashboard"
echo ""

# ============================================================================
# OPTION 2: DOCKER BUILD & RUN
# ============================================================================
echo ""
echo "📌 OPTION 2: Docker (For Local Testing or VPS)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Build Docker Image:"
echo "  docker build -t expense-api:latest ."
echo ""
echo "Run Container:"
echo "  docker run -p 8080:8080 expense-api:latest"
echo ""
echo "Or with Docker Compose:"
echo "  docker-compose up -d"
echo ""

# ============================================================================
# OPTION 3: RENDER.COM
# ============================================================================
echo ""
echo "📌 OPTION 3: Render.com"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "1. Go to: https://render.com"
echo "2. Click 'New +' → 'Web Service'"
echo "3. Select GitHub repository"
echo "4. Environment: Docker"
echo "5. Click 'Create Web Service'"
echo "6. Get public URL from dashboard"
echo ""

# ============================================================================
# OPTION 4: FLY.IO
# ============================================================================
echo ""
echo "📌 OPTION 4: Fly.io"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "1. Install Fly CLI:"
echo "   brew install flyctl"
echo ""
echo "2. Deploy:"
echo "   flyctl auth login"
echo "   flyctl launch"
echo "   flyctl deploy"
echo ""

# ============================================================================
# TESTING DEPLOYED API
# ============================================================================
echo ""
echo "✅ TESTING DEPLOYED API"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Replace YOUR_URL with actual deployed URL"
echo ""
echo "1. Health Check:"
echo "   curl https://YOUR_URL/health"
echo ""
echo "2. Create Expense:"
echo "   curl -X POST https://YOUR_URL/expenses \\"
echo "     -H 'Content-Type: application/json' \\"
echo "     -d '{\"description\":\"Coffee\",\"amount\":5.5,\"category\":\"Food\",\"date\":\"2025-11-26\"}'"
echo ""
echo "3. Get All Expenses:"
echo "   curl https://YOUR_URL/expenses"
echo ""
echo "4. Get Single Expense:"
echo "   curl https://YOUR_URL/expenses/1"
echo ""
echo "5. Update Expense:"
echo "   curl -X PUT https://YOUR_URL/expenses/1 \\"
echo "     -H 'Content-Type: application/json' \\"
echo "     -d '{\"amount\":6.5}'"
echo ""
echo "6. Delete Expense:"
echo "   curl -X DELETE https://YOUR_URL/expenses/1"
echo ""

# ============================================================================
# LOCAL DEVELOPMENT
# ============================================================================
echo ""
echo "🔨 LOCAL DEVELOPMENT"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Build:"
echo "  go build -o expense-api"
echo ""
echo "Run:"
echo "  ./expense-api"
echo ""
echo "With custom port:"
echo "  PORT=3000 ./expense-api"
echo ""
echo "Test (in another terminal):"
echo "  curl http://localhost:8080/health"
echo ""

# ============================================================================
# GIT SETUP
# ============================================================================
echo ""
echo "📝 GIT SETUP (For GitHub)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Initialize git:"
echo "  git init"
echo "  git add ."
echo "  git commit -m 'Initial commit: Expense API'"
echo ""
echo "Push to GitHub:"
echo "  git remote add origin https://github.com/YOUR_USERNAME/expense-backend.git"
echo "  git branch -M main"
echo "  git push -u origin main"
echo ""

# ============================================================================
# IMPORTANT FILES
# ============================================================================
echo ""
echo "📚 DOCUMENTATION"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Read these files:"
echo "  1. README.md          - API documentation"
echo "  2. API_GUIDE.md       - Complete guide for mobile dev"
echo "  3. QUICKSTART.md      - Quick start guide"
echo "  4. DEPLOYMENT.md      - Detailed deployment guide"
echo "  5. PROJECT_STATUS.md  - Project status & checklist"
echo ""

# ============================================================================
# RECOMMENDATION
# ============================================================================
echo ""
echo "⭐ RECOMMENDATION"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "For FASTEST deployment: Use Railway.app"
echo ""
echo "Why Railway.app?"
echo "  ✅ Easiest setup (5 minutes)"
echo "  ✅ Free tier for 5 projects"
echo "  ✅ Auto HTTPS"
echo "  ✅ Good uptime"
echo "  ✅ Auto redeploy on push"
echo ""

# ============================================================================
# SUMMARY
# ============================================================================
echo ""
echo "📋 NEXT STEPS"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "1. Choose deployment option above"
echo "2. Follow the steps for your chosen platform"
echo "3. Get the public URL"
echo "4. Test with curl/Postman"
echo "5. Share URL with mobile developer"
echo ""
echo "Good luck! 🚀"
echo ""
