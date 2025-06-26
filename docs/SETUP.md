# 🚀 Setup Guide

Complete installation and setup guide for the Debt Tracker application.

## ⚠️ System Requirements

**This project requires Linux/Ubuntu environment:**

- **Ubuntu 20.04 LTS or higher** (recommended)
- **WSL with Ubuntu** (for Windows users)
- **macOS** (with additional configuration)

> **Windows Users**: Install WSL2 with Ubuntu from Microsoft Store, then follow this guide within WSL.

## 📋 Prerequisites

Install the following software on your Ubuntu/Linux system:

### 1. Go (Golang)
```bash
# Remove any existing Go installation
sudo rm -rf /usr/local/go

# Download and install Go (check for latest version)
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Add to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verify installation
go version
```

### 2. Git
```bash
# Install Git
sudo apt update
sudo apt install git

# Verify installation
git --version

# Configure Git (if not already done)
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"
```

### 3. Additional Tools
```bash
# Install essential tools
sudo apt install curl sqlite3 build-essential

# Optional: Install VS Code
wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > packages.microsoft.gpg
sudo install -o root -g root -m 644 packages.microsoft.gpg /etc/apt/trusted.gpg.d/
sudo sh -c 'echo "deb [arch=amd64,arm64,armhf signed-by=/etc/apt/trusted.gpg.d/packages.microsoft.gpg] https://packages.microsoft.com/repos/code stable main" > /etc/apt/sources.list.d/vscode.list'
sudo apt update
sudo apt install code
```

## 🔄 Project Setup

### 1. Clone the Repository

**For Contributors (Fork first):**
```bash
# Fork the repository on GitHub, then:
git clone https://github.com/YOUR_USERNAME/debt-tracker-app.git
cd debt-tracker-app

# Add upstream remote for syncing
git remote add upstream https://github.com/MAIN_USERNAME/debt-tracker-app.git
```

**For Direct Collaborators:**
```bash
git clone https://github.com/MAIN_USERNAME/debt-tracker-app.git
cd debt-tracker-app
```

### 2. Backend Setup

```bash
# Navigate to backend directory
cd backend

# Initialize Go module (if not already done)
go mod init debt-tracker-backend

# Install dependencies
go get github.com/gin-gonic/gin
go get github.com/golang-jwt/jwt/v4
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv
go get modernc.org/sqlite

# Tidy up modules
go mod tidy
```

### 3. Create Project Structure

```bash
# Ensure you're in the backend directory
cd backend

# Create directory structure
mkdir -p cmd/server
mkdir -p internal/{handlers,models,database,middleware}
mkdir -p configs

# Create essential files (if they don't exist)
touch cmd/server/main.go
touch internal/database/database.go
touch internal/models/models.go
touch internal/handlers/auth.go
touch internal/handlers/contacts.go
touch internal/handlers/debts.go
touch internal/handlers/transactions.go
touch internal/middleware/auth.go
touch configs/config.go
touch .gitignore
```

### 4. Environment Configuration

Create `.env` file in the backend directory:

```bash
cd backend
cat > .env << EOF
DATABASE_URL=debt_tracker.db
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production-$(openssl rand -hex 32)
ENVIRONMENT=development
PORT=8080
EOF
```

### 5. Frontend Setup

```bash
# Navigate to frontend directory
cd ../frontend

# Create frontend structure (if not exists)
mkdir -p assets/{css,js,images}
touch index.html

# The main dashboard file should already exist from the repository
```

## 🏃‍♂️ Running the Application

### 1. Start Backend Server

```bash
cd backend

# Run the application
go run cmd/server/main.go
```

Expected output:
```
[GIN-debug] [WARNING] Running in "debug" mode
[GIN-debug] GET    /health                   --> main.main.func1
[GIN-debug] POST   /api/v1/auth/register     --> ...
[GIN-debug] POST   /api/v1/auth/login        --> ...
...
Server starting on port 8080
```

### 2. Start Frontend Server

Open a new terminal:

```bash
cd frontend

# Option 1: Python HTTP server
python3 -m http.server 3000

# Option 2: Node.js http-server (if installed)
npx http-server -p 3000

# Option 3: Direct file access
# Open index.html directly in your browser
```

### 3. Access the Application

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Health Check**: http://localhost:8080/health

## 🧪 Testing the Setup

### 1. Health Check
```bash
curl http://localhost:8080/health
```

Expected response:
```json
{"status":"healthy","service":"debt-tracker-api"}
```

### 2. User Registration
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123",
    "name": "Test User",
    "phone": "+1234567890"
  }'
```

### 3. Frontend Access
1. Open http://localhost:3000 in your browser
2. Register a new account
3. Add a contact
4. Create a debt
5. Add a transaction

## 🛠️ Development Tools

### VS Code Extensions (Recommended)

```bash
# Install VS Code extensions
code --install-extension golang.Go
code --install-extension ms-vscode.vscode-json
code --install-extension bradlc.vscode-tailwindcss
code --install-extension formulahendry.auto-rename-tag
```

### VS Code Workspace Settings

Create `.vscode/settings.json`:
```json
{
    "go.useLanguageServer": true,
    "go.formatTool": "goimports",
    "go.lintTool": "golangci-lint",
    "go.testFlags": ["-v"],
    "files.autoSave": "onFocusChange",
    "editor.formatOnSave": true
}
```

## 🐛 Troubleshooting

### Common Issues

**1. Port Already in Use**
```bash
# Find and kill process using port 8080
sudo lsof -i :8080
sudo kill -9 <PID>

# Or use a different port
PORT=8081 go run cmd/server/main.go
```

**2. Go Module Issues**
```bash
# Clean module cache
go clean -modcache
go mod download
go mod tidy
```

**3. Permission Issues**
```bash
# Fix file permissions
chmod +x scripts/*.sh
sudo chown -R $USER:$USER .
```

**4. Database Issues**
```bash
# Reset database (will lose data)
rm backend/debt_tracker.db
# Restart server to recreate
```

**5. Frontend Not Loading**
- Ensure backend is running on port 8080
- Check browser console for errors
- Verify API calls in Network tab
- Try incognito/private browsing mode

### WSL-Specific Issues

**File System Performance:**
```bash
# Store project in WSL file system, not Windows
# Good: /home/username/projects/debt-tracker-app
# Bad: /mnt/c/Users/username/projects/debt-tracker-app
```

**Network Issues:**
```bash
# If localhost doesn't work, try WSL IP
ip addr show eth0
# Use the inet address instead of localhost
```

## 🔧 Development Workflow

### 1. Daily Development
```bash
# Pull latest changes
git pull upstream main

# Create feature branch
git checkout -b feature/your-feature-name

# Start development servers
# Terminal 1: Backend
cd backend && go run cmd/server/main.go

# Terminal 2: Frontend  
cd frontend && python3 -m http.server 3000

# Terminal 3: Development work
code .
```

### 2. Code Quality
```bash
# Format Go code
go fmt ./...

# Run tests (when available)
go test ./...

# Check for security issues
go mod download
go list -json -m all | nancy sleuth
```

### 3. Git Workflow
```bash
# Stage changes
git add .

# Commit with descriptive message
git commit -m "feat: add transaction filtering feature"

# Push to your fork
git push origin feature/your-feature-name

# Create pull request on GitHub
```

## 📦 Database Management

### Backup Database
```bash
# Create backup
cp backend/debt_tracker.db backend/debt_tracker_backup_$(date +%Y%m%d).db
```

### Reset Database
```bash
# Stop server
# Remove database file
rm backend/debt_tracker.db
# Restart server (will recreate tables)
```

### View Database
```bash
# Install SQLite browser
sudo apt install sqlitebrowser

# Open database
sqlitebrowser backend/debt_tracker.db

# Or use command line
sqlite3 backend/debt_tracker.db
.tables
.schema users
.quit
```

## 🚀 Next Steps

After successful setup:

1. **Read the codebase**: Understand the project structure
2. **Check issues**: Look at GitHub issues for tasks to work on  
3. **Read ROADMAP.md**: Understand planned features
4. **Start contributing**: Pick a feature or bug to work on

## 📞 Getting Help

If you encounter issues:

1. **Check this guide**: Follow each step carefully
2. **Search existing issues**: Someone might have faced the same problem
3. **Create an issue**: Describe your problem with system details
4. **Ask in discussions**: Use GitHub Discussions for questions

---

**Happy coding! 🚀**