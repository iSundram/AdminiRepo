
<template>
  <div class="dashboard">
    <div class="container">
      <h2>AdminiSoftware Dashboard</h2>
      
      <div class="welcome-banner">
        <h3>Welcome to the Future of Web Hosting Control</h3>
        <p>Experience the most advanced control panel with AI-powered features, automated security, and next-gen tools.</p>
      </div>
      
      <div class="features-grid">
        <div class="feature-card">
          <h4>🤖 AI-Powered Automation</h4>
          <p>Intelligent server health prediction and self-healing infrastructure</p>
        </div>
        
        <div class="feature-card">
          <h4>🌐 Cloud-Native Integration</h4>
          <p>One-click CDN integration and global load balancing</p>
        </div>
        
        <div class="feature-card">
          <h4>🔒 Advanced Security</h4>
          <p>AI-driven threat detection and automated malware removal</p>
        </div>
        
        <div class="feature-card">
          <h4>📊 Smart Analytics</h4>
          <p>Real-time monitoring with predictive insights</p>
        </div>
      </div>
      
      <div class="status-section">
        <h3>System Status</h3>
        <div class="status-grid">
          <div class="status-item">
            <span class="status-indicator online"></span>
            <span>API Service: Online</span>
          </div>
          <div class="status-item">
            <span class="status-indicator online"></span>
            <span>Database: Connected</span>
          </div>
          <div class="status-item">
            <span class="status-indicator online"></span>
            <span>Security: Active</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dashboard'
}
</script>

<style scoped>
.dashboard {
  padding: 2rem 0;
}

.welcome-banner {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 2rem;
  border-radius: 10px;
  margin: 2rem 0;
  text-align: center;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1.5rem;
  margin: 2rem 0;
}

.feature-card {
  background: white;
  padding: 1.5rem;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
  border: 1px solid #eee;
}

.feature-card h4 {
  margin-bottom: 0.5rem;
  color: #2c3e50;
}

.status-section {
  margin: 3rem 0;
}

.status-grid {
  display: flex;
  gap: 2rem;
  flex-wrap: wrap;
}

.status-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: white;
  border-radius: 5px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}

.status-indicator {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}

.status-indicator.online {
  background: #27ae60;
}
</style>
