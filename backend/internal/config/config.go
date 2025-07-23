
package config

// Configuration management
type Config struct {
    Database DatabaseConfig `json:"database"`
    Server   ServerConfig   `json:"server"`
    Security SecurityConfig `json:"security"`
}

type DatabaseConfig struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
    Database string `json:"database"`
}

type ServerConfig struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

type SecurityConfig struct {
    JWTSecret string `json:"jwt_secret"`
    BCryptCost int   `json:"bcrypt_cost"`
}
