
#!/bin/bash

# AdminiSoftware Cleanup Script
# Runs periodic cleanup tasks

LOG_FILE="/var/log/adminisoftware/cleanup.log"

echo "$(date): Starting cleanup tasks..." >> "$LOG_FILE"

# Clean temporary files
find /tmp/adminisoftware -type f -mtime +7 -delete 2>/dev/null || true

# Clean old log files
find /var/log/adminisoftware -name "*.log" -mtime +30 -delete 2>/dev/null || true

# Clean expired sessions
# TODO: Implement session cleanup

# Clean old metrics
find /tmp/adminisoftware/metrics -type f -mtime +1 -delete 2>/dev/null || true

echo "$(date): Cleanup tasks completed" >> "$LOG_FILE"
