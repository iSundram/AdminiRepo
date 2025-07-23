
#!/bin/bash

# AdminiSoftware Metrics Collector
# Collects system metrics for AI analysis and monitoring

LOG_FILE="/var/log/adminisoftware/metrics.log"
METRICS_DIR="/tmp/adminisoftware/metrics"

# Create directories if they don't exist
mkdir -p "$(dirname "$LOG_FILE")"
mkdir -p "$METRICS_DIR"

echo "$(date): Starting metrics collection..." >> "$LOG_FILE"

# Collect CPU metrics
cpu_usage=$(top -bn1 | grep "Cpu(s)" | sed "s/.*, *\([0-9.]*\)%* id.*/\1/" | awk '{print 100 - $1}')
echo "cpu_usage:$cpu_usage" > "$METRICS_DIR/cpu.metric"

# Collect memory metrics
memory_usage=$(free | grep Mem | awk '{printf "%.2f", $3/$2 * 100.0}')
echo "memory_usage:$memory_usage" > "$METRICS_DIR/memory.metric"

# Collect disk usage
disk_usage=$(df -h / | awk 'NR==2 {print $5}' | sed 's/%//')
echo "disk_usage:$disk_usage" > "$METRICS_DIR/disk.metric"

# Collect network metrics
network_rx=$(cat /proc/net/dev | grep eth0 | awk '{print $2}')
network_tx=$(cat /proc/net/dev | grep eth0 | awk '{print $10}')
echo "network_rx:$network_rx" > "$METRICS_DIR/network_rx.metric"
echo "network_tx:$network_tx" > "$METRICS_DIR/network_tx.metric"

# Collect load average
load_avg=$(uptime | awk -F'load average:' '{print $2}' | awk -F, '{print $1}' | tr -d ' ')
echo "load_average:$load_avg" > "$METRICS_DIR/load.metric"

# Collect process count
process_count=$(ps aux | wc -l)
echo "process_count:$process_count" > "$METRICS_DIR/processes.metric"

echo "$(date): Metrics collection completed" >> "$LOG_FILE"

# Send metrics to AI analysis service (if enabled)
if [ "$AI_ANALYSIS_ENABLED" = "true" ]; then
    curl -X POST http://localhost:5000/api/v1/ai/metrics \
         -H "Content-Type: application/json" \
         -d @"$METRICS_DIR/combined.json" \
         >> "$LOG_FILE" 2>&1
fi
