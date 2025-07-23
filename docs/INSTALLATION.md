
# AdminiSoftware Installation Guide

## System Requirements

### Minimum Requirements
- **OS:** CentOS 7+, Ubuntu 18.04+, Debian 9+, or Rocky Linux 8+
- **CPU:** 2 cores
- **RAM:** 4GB
- **Storage:** 20GB SSD
- **Network:** 100Mbps connection

### Recommended Requirements
- **OS:** Rocky Linux 9+ or Ubuntu 22.04+
- **CPU:** 4+ cores
- **RAM:** 8GB+
- **Storage:** 50GB+ NVMe SSD
- **Network:** 1Gbps connection

## Prerequisites

### Required Software
- Go 1.21+
- Node.js 18+
- PostgreSQL 13+
- Redis 6+
- Nginx

## Installation Methods

### Method 1: Automatic Installation (Recommended)

```bash
# Download and run the installer
curl -sSL https://raw.githubusercontent.com/adminisoftware/installer/main/install.sh | bash
```

### Method 2: Manual Installation

#### Step 1: Install System Dependencies

**CentOS/Rocky Linux:**
```bash
# Update system
sudo dnf update -y

# Install required packages
sudo dnf install -y git curl wget unzip nginx postgresql-server postgresql-contrib redis

# Install Go
wget https://golang.org/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Install Node.js
curl -fsSL https://rpm.nodesource.com/setup_18.x | sudo bash -
sudo dnf install -y nodejs
```

**Ubuntu/Debian:**
```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install required packages
sudo apt install -y git curl wget unzip nginx postgresql postgresql-contrib redis-server

# Install Go
wget https://golang.org/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Install Node.js
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt install -y nodejs
```

#### Step 2: Configure Database

```bash
# Initialize PostgreSQL (CentOS/Rocky only)
sudo postgresql-setup --initdb
sudo systemctl enable postgresql
sudo systemctl start postgresql

# Create database and user
sudo -u postgres psql -c "CREATE DATABASE adminisoftware;"
sudo -u postgres psql -c "CREATE USER adminisoftware WITH PASSWORD 'your_secure_password';"
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE adminisoftware TO adminisoftware;"
```

#### Step 3: Download and Build AdminiSoftware

```bash
# Clone repository
git clone https://github.com/adminisoftware/adminisoftware.git
cd adminisoftware

# Build backend
cd backend
cp .env.example .env
# Edit .env file with your configuration
go mod tidy
go build -o adminisoftwared cmd/adminisoftwared.go

# Build frontend
cd ../frontend
npm install
npm run build

# Return to project root
cd ..
```

#### Step 4: Configure Environment

```bash
# Edit backend configuration
nano backend/.env
```

**Example .env configuration:**
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=adminisoftware
DB_USER=adminisoftware
DB_PASSWORD=your_secure_password

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT Configuration
JWT_SECRET=your_jwt_secret_key_here

# Server Configuration
PORT=5000
DEBUG=false

# Email Configuration
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your_email@gmail.com
SMTP_PASSWORD=your_app_password

# SSL Configuration
SSL_ENABLED=true
SSL_CERT_PATH=/etc/ssl/certs/adminisoftware.crt
SSL_KEY_PATH=/etc/ssl/private/adminisoftware.key

# Backup Configuration
BACKUP_PATH=/var/backups/adminisoftware
BACKUP_RETENTION_DAYS=30

# Security Configuration
RATE_LIMIT_ENABLED=true
RATE_LIMIT_REQUESTS=1000
RATE_LIMIT_WINDOW=3600

# External APIs
CLOUDFLARE_API_KEY=your_cloudflare_api_key
AWS_ACCESS_KEY_ID=your_aws_access_key
AWS_SECRET_ACCESS_KEY=your_aws_secret_key
```

#### Step 5: Configure Nginx

```bash
# Create Nginx configuration
sudo nano /etc/nginx/sites-available/adminisoftware
```

**Nginx configuration:**
```nginx
server {
    listen 80;
    server_name your-domain.com;
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl http2;
    server_name your-domain.com;

    ssl_certificate /etc/ssl/certs/adminisoftware.crt;
    ssl_certificate_key /etc/ssl/private/adminisoftware.key;

    # SSL Security
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES256-GCM-SHA512:DHE-RSA-AES256-GCM-SHA512;
    ssl_prefer_server_ciphers off;
    ssl_session_cache shared:SSL:10m;

    # Security Headers
    add_header X-Frame-Options DENY;
    add_header X-Content-Type-Options nosniff;
    add_header X-XSS-Protection "1; mode=block";
    add_header Strict-Transport-Security "max-age=63072000";

    # Proxy to AdminiSoftware
    location / {
        proxy_pass http://127.0.0.1:5000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # API routes
    location /api/ {
        proxy_pass http://127.0.0.1:5000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

```bash
# Enable site
sudo ln -s /etc/nginx/sites-available/adminisoftware /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

#### Step 6: Create Systemd Service

```bash
# Create systemd service file
sudo nano /etc/systemd/system/adminisoftware.service
```

**Service configuration:**
```ini
[Unit]
Description=AdminiSoftware Control Panel
After=network.target postgresql.service redis.service

[Service]
Type=simple
User=adminisoftware
Group=adminisoftware
WorkingDirectory=/opt/adminisoftware
ExecStart=/opt/adminisoftware/backend/adminisoftwared
Restart=always
RestartSec=5
Environment=GIN_MODE=release

[Install]
WantedBy=multi-user.target
```

```bash
# Create user and set permissions
sudo useradd -r -s /bin/false adminisoftware
sudo mkdir -p /opt/adminisoftware
sudo cp -r . /opt/adminisoftware/
sudo chown -R adminisoftware:adminisoftware /opt/adminisoftware

# Enable and start service
sudo systemctl daemon-reload
sudo systemctl enable adminisoftware
sudo systemctl start adminisoftware
```

#### Step 7: Run Database Migrations

```bash
# Run initial migrations
cd /opt/adminisoftware/backend
sudo -u adminisoftware ./adminisoftwared migrate
```

## Post-Installation Setup

### Initial Configuration

1. **Access the web interface:**
   ```
   https://your-domain.com
   ```

2. **Default login credentials:**
   - Username: `admin`
   - Password: `admin123`

3. **Change default password immediately after first login**

### SSL Certificate Setup

**Using Let's Encrypt:**
```bash
# Install Certbot
sudo apt install certbot python3-certbot-nginx  # Ubuntu/Debian
sudo dnf install certbot python3-certbot-nginx  # CentOS/Rocky

# Obtain certificate
sudo certbot --nginx -d your-domain.com

# Auto-renewal
sudo crontab -e
# Add: 0 12 * * * /usr/bin/certbot renew --quiet
```

### Firewall Configuration

```bash
# UFW (Ubuntu/Debian)
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable

# Firewalld (CentOS/Rocky)
sudo firewall-cmd --permanent --add-service=ssh
sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https
sudo firewall-cmd --reload
```

## Verification

### Check Service Status
```bash
sudo systemctl status adminisoftware
sudo systemctl status nginx
sudo systemctl status postgresql
sudo systemctl status redis
```

### View Logs
```bash
# AdminiSoftware logs
sudo journalctl -u adminisoftware -f

# Nginx logs
sudo tail -f /var/log/nginx/access.log
sudo tail -f /var/log/nginx/error.log
```

### Test API
```bash
curl -k https://your-domain.com/api/v1/health
```

## Troubleshooting

### Common Issues

1. **Service won't start:**
   - Check configuration file syntax
   - Verify database connection
   - Check file permissions

2. **Database connection errors:**
   - Verify PostgreSQL is running
   - Check database credentials
   - Ensure database exists

3. **Permission denied errors:**
   - Check file ownership
   - Verify user permissions
   - Check SELinux/AppArmor policies

### Getting Help

- **Documentation:** https://docs.admini.tech
- **Support:** support@admini.tech
- **Community:** https://community.admini.tech
- **GitHub Issues:** https://github.com/adminisoftware/adminisoftware/issues

## Maintenance

### Regular Tasks

1. **Update system packages weekly**
2. **Monitor disk space and logs**
3. **Review security alerts**
4. **Backup configuration and data**
5. **Test disaster recovery procedures**

### Automated Backups

```bash
# Add to crontab
0 2 * * * /opt/adminisoftware/scripts/backup.sh
```

This completes the installation of AdminiSoftware. The system should now be fully operational and accessible via your domain.
