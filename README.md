
# AdminiSoftware - Next-Gen Control Panel

AdminiSoftware is an ultra-modern, AI-powered web hosting control panel designed as a comprehensive alternative to cPanel, Webuzo, and WHM. Built with cutting-edge technology and forward-thinking features.

## 🌟 Features

- **AI-Powered Automation**: Intelligent server health prediction and self-healing infrastructure
- **Cloud-Native Integration**: One-click CDN integration and global load balancing  
- **Advanced Security**: AI-driven threat detection and automated malware removal
- **Smart Analytics**: Real-time monitoring with predictive insights
- **Modern UI/UX**: Responsive design with dark mode support
- **Multi-tenancy**: Support for resellers and end-users
- **Extensive API**: RESTful API for all operations

## 🚀 Quick Start

1. **Backend Setup**:
   ```bash
   cd backend
   cp .env.example .env
   go mod tidy
   go run cmd/adminisoftwared.go
   ```

2. **Frontend Setup**:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

3. **Access the Panel**:
   - Open your browser to `http://localhost:5000`
   - Use demo credentials: `admin@admini.tech` / `admin123`

## 🏗️ Architecture

- **Backend**: Go with Gin framework
- **Frontend**: Vue.js 3 with Vite
- **Database**: PostgreSQL with GORM
- **Authentication**: JWT with 2FA support
- **Security**: Advanced rate limiting and brute force protection

## 📧 Contact

- **Website**: [admini.tech](https://admini.tech)
- **Support**: support@admini.tech

## 📄 License

AdminiSoftware - All rights reserved.
