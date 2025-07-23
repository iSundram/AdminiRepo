
#!/bin/bash

# AdminiSoftware Health Check Script
# Monitors system health and services

HEALTH_ENDPOINT="http://localhost:5000/health"
LOG_FILE="/var/log/adminisoftware/health.log"

echo "$(date): Running health check..." >> "$LOG_FILE"

# Check API health
if curl -s "$HEALTH_ENDPOINT" > /dev/null; then
    echo "$(date): API health check passed" >> "$LOG_FILE"
else
    echo "$(date): API health check failed" >> "$LOG_FILE"
    exit 1
fi

# Check disk space
DISK_USAGE=$(df / | awk 'NR==2 {print $5}' | sed 's/%//')
if [ "$DISK_USAGE" -gt 90 ]; then
    echo "$(date): WARNING - Disk usage is ${DISK_USAGE}%" >> "$LOG_FILE"
fi

# Check memory usage
MEMORY_USAGE=$(free | grep Mem | awk '{printf "%.0f", $3/$2 * 100.0}')
if [ "$MEMORY_USAGE" -gt 90 ]; then
    echo "$(date): WARNING - Memory usage is ${MEMORY_USAGE}%" >> "$LOG_FILE"
fi

echo "$(date): Health check completed" >> "$LOG_FILE"
