
# AdminiSoftware API Documentation

## Base URL
```
https://your-domain.com/api/v1
```

## Authentication
All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your-token>
```

## Response Format
All API responses follow this format:
```json
{
  "status": "success|error",
  "message": "Response message",
  "data": {} // Response data (for successful requests)
}
```

## Authentication Endpoints

### POST /auth/login
Login to the system.

**Request:**
```json
{
  "username": "admin",
  "password": "admin123"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_at": 1705651200,
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@admini.tech",
      "role": "admin"
    }
  }
}
```

### POST /auth/logout
Logout from the system.

### GET /auth/profile
Get current user profile (requires authentication).

### PUT /auth/profile
Update user profile (requires authentication).

## Admin Endpoints (Admin Role Required)

### GET /admin/users
Get all system users.

### POST /admin/users
Create a new user.

**Request:**
```json
{
  "username": "newuser",
  "email": "user@example.com",
  "password": "password123",
  "role": "user"
}
```

### PUT /admin/users/:id
Update a user.

### DELETE /admin/users/:id
Delete a user.

### GET /admin/accounts
Get all hosting accounts.

### POST /admin/accounts
Create a new hosting account.

**Request:**
```json
{
  "domain": "example.com",
  "username": "user123",
  "password": "password123",
  "email": "user@example.com",
  "package": "Basic"
}
```

### GET /admin/packages
Get all hosting packages.

### POST /admin/packages
Create a new hosting package.

**Request:**
```json
{
  "name": "Premium",
  "disk_quota": 50000,
  "bandwidth": 500000,
  "email_quota": 1000,
  "ftp_quota": 50,
  "db_quota": 50,
  "price": 49.99,
  "features": ["SSL Certificate", "Email Accounts", "Backup Service"]
}
```

### GET /admin/system/info
Get system information.

### GET /admin/system/logs
Get system logs.

**Query Parameters:**
- `type` - Log type filter
- `limit` - Number of logs to return (default: 100)

## User Panel Endpoints (User Role Required)

### GET /user/dashboard
Get user dashboard data.

### GET /user/domains
Get user's domains.

### POST /user/domains
Add a new domain/subdomain.

**Request:**
```json
{
  "domain": "subdomain.example.com",
  "type": "subdomain"
}
```

### GET /user/domains/:id/dns
Get DNS records for a domain.

### POST /user/domains/:id/dns
Add a DNS record.

**Request:**
```json
{
  "name": "www",
  "type": "CNAME",
  "value": "example.com",
  "ttl": 3600
}
```

### GET /user/files
Get files in a directory.

**Query Parameters:**
- `path` - Directory path (default: "/")

### POST /user/files/upload
Upload a file.

**Form Data:**
- `file` - File to upload
- `path` - Target directory path

### POST /user/files/folder
Create a new folder.

**Request:**
```json
{
  "path": "/",
  "folder_name": "new-folder"
}
```

### GET /user/email/accounts
Get email accounts.

### POST /user/email/accounts
Create an email account.

**Request:**
```json
{
  "email": "info@example.com",
  "password": "password123",
  "quota": 1000
}
```

### GET /user/databases
Get databases.

### POST /user/databases
Create a database.

**Request:**
```json
{
  "name": "user123_newdb",
  "charset": "utf8mb4",
  "collation": "utf8mb4_unicode_ci"
}
```

### GET /user/ssl
Get SSL certificates.

### POST /user/ssl
Request an SSL certificate.

**Request:**
```json
{
  "domain": "example.com",
  "type": "letsencrypt",
  "auto_renew": true
}
```

## Reseller Endpoints (Reseller Role Required)

### GET /reseller/dashboard
Get reseller dashboard data.

### GET /reseller/accounts
Get client accounts.

### POST /reseller/accounts
Create a client account.

**Request:**
```json
{
  "domain": "client.com",
  "username": "client123",
  "password": "password123",
  "email": "client@example.com",
  "package": "Professional"
}
```

## AI Assistant Endpoints

### POST /ai/assistant
Get AI assistant response.

**Request:**
```json
{
  "query": "How can I optimize my server performance?"
}
```

### GET /ai/analysis
Get AI-powered system analysis.

### GET /ai/suggestions
Get optimization suggestions.

### GET /ai/predictions
Get resource usage predictions.

## Plugin Marketplace Endpoints

### GET /plugins
Get available plugins.

**Query Parameters:**
- `category` - Filter by category
- `search` - Search term

### GET /plugins/:id
Get plugin details.

### POST /plugins/:id/install
Install a plugin.

### DELETE /plugins/:id
Uninstall a plugin.

### GET /plugins/installed
Get installed plugins.

## Error Codes

- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `422` - Validation Error
- `500` - Internal Server Error

## Rate Limiting

API requests are rate limited to prevent abuse:
- 1000 requests per hour per IP address
- 100 requests per minute per authenticated user

When rate limit is exceeded, you'll receive a `429 Too Many Requests` response.
