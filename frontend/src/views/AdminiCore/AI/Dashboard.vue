
<template>
  <div class="ai-dashboard">
    <div class="header">
      <h1>🤖 AI Control Center</h1>
      <p>Intelligent automation and predictive insights</p>
    </div>

    <div class="ai-grid">
      <!-- AI Copilot -->
      <div class="card copilot-card">
        <h3>💬 AI Assistant</h3>
        <div class="copilot-chat">
          <div class="chat-messages">
            <div v-for="message in conversation" :key="message.id" 
                 :class="['message', message.type]">
              <span class="content">{{ message.content }}</span>
              <span class="timestamp">{{ formatTime(message.timestamp) }}</span>
            </div>
          </div>
          <div class="chat-input">
            <input v-model="chatInput" @keyup.enter="askCopilot" 
                   placeholder="Ask AI anything about your server..." />
            <button @click="askCopilot">Send</button>
          </div>
        </div>
      </div>

      <!-- System Health Prediction -->
      <div class="card prediction-card">
        <h3>📊 Health Prediction</h3>
        <div class="health-score">
          <div class="score-circle" :style="{ '--score': healthScore }">
            <span class="score-value">{{ healthScore }}</span>
            <span class="score-label">Health Score</span>
          </div>
        </div>
        <div class="predictions">
          <div class="prediction-item">
            <span class="label">CPU (Next Hour)</span>
            <span class="value">{{ resourcePrediction.cpu }}%</span>
          </div>
          <div class="prediction-item">
            <span class="label">Memory (Next Hour)</span>
            <span class="value">{{ resourcePrediction.memory }}%</span>
          </div>
          <div class="prediction-item">
            <span class="label">Storage (Next Week)</span>
            <span class="value">{{ resourcePrediction.storage }}%</span>
          </div>
        </div>
      </div>

      <!-- Anomaly Detection -->
      <div class="card anomaly-card">
        <h3>🔍 Anomaly Detection</h3>
        <div class="anomaly-status">
          <span class="status-indicator" :class="anomalyStatus">{{ anomalyStatus }}</span>
          <span class="status-text">System Behavior Normal</span>
        </div>
        <div class="recent-anomalies">
          <div v-for="anomaly in recentAnomalies" :key="anomaly.id" class="anomaly-item">
            <span class="type">{{ anomaly.type }}</span>
            <span class="severity" :class="anomaly.severity">{{ anomaly.severity }}</span>
            <span class="time">{{ formatTime(anomaly.timestamp) }}</span>
          </div>
        </div>
      </div>

      <!-- AI Recommendations -->
      <div class="card recommendations-card">
        <h3>💡 AI Recommendations</h3>
        <div class="recommendations">
          <div v-for="rec in recommendations" :key="rec.id" class="recommendation">
            <div class="rec-icon">{{ rec.icon }}</div>
            <div class="rec-content">
              <h4>{{ rec.title }}</h4>
              <p>{{ rec.description }}</p>
              <div class="rec-actions">
                <button @click="applyRecommendation(rec)" class="apply-btn">Apply</button>
                <button @click="dismissRecommendation(rec)" class="dismiss-btn">Dismiss</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- AI Models Status -->
      <div class="card models-card">
        <h3>🧠 AI Models</h3>
        <div class="models-list">
          <div v-for="model in aiModels" :key="model.id" class="model-item">
            <span class="model-name">{{ model.name }}</span>
            <span class="model-accuracy">{{ model.accuracy }}% accurate</span>
            <span class="model-status" :class="model.status">{{ model.status }}</span>
          </div>
        </div>
      </div>

      <!-- Automation Status -->
      <div class="card automation-card">
        <h3>⚙️ AI Automation</h3>
        <div class="automation-stats">
          <div class="stat">
            <span class="stat-value">{{ automationStats.tasksAutomated }}</span>
            <span class="stat-label">Tasks Automated</span>
          </div>
          <div class="stat">
            <span class="stat-value">{{ automationStats.timeSaved }}</span>
            <span class="stat-label">Hours Saved</span>
          </div>
          <div class="stat">
            <span class="stat-value">{{ automationStats.efficiency }}</span>
            <span class="stat-label">Efficiency Gain</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions, mapGetters } from 'vuex'

export default {
  name: 'AIDashboard',
  data() {
    return {
      chatInput: '',
      conversation: [
        {
          id: 1,
          type: 'assistant',
          content: 'Hello! I\'m your AI assistant. How can I help you optimize your server today?',
          timestamp: new Date()
        }
      ],
      resourcePrediction: {
        cpu: 85.5,
        memory: 67.2,
        storage: 45.8
      },
      anomalyStatus: 'normal',
      recentAnomalies: [
        { id: 1, type: 'CPU Spike', severity: 'low', timestamp: new Date() },
        { id: 2, type: 'Memory Usage', severity: 'medium', timestamp: new Date() }
      ],
      recommendations: [
        {
          id: 1,
          icon: '🚀',
          title: 'Upgrade PHP Version',
          description: 'Switching to PHP 8.1 could improve performance by 25%'
        },
        {
          id: 2,
          icon: '🔒',
          title: 'SSL Certificate Renewal',
          description: 'Your SSL certificate expires in 30 days'
        }
      ],
      aiModels: [
        { id: 1, name: 'Resource Predictor', accuracy: 94.2, status: 'active' },
        { id: 2, name: 'Anomaly Detector', accuracy: 88.7, status: 'active' },
        { id: 3, name: 'Security Analyzer', accuracy: 91.5, status: 'training' }
      ],
      automationStats: {
        tasksAutomated: 247,
        timeSaved: 38.5,
        efficiency: '+42%'
      }
    }
  },
  computed: {
    ...mapGetters('ai', ['healthScore']),
    healthScore() {
      return 95 // Mock data
    }
  },
  methods: {
    askCopilot() {
      if (!this.chatInput.trim()) return
      
      this.conversation.push({
        id: Date.now(),
        type: 'user',
        content: this.chatInput,
        timestamp: new Date()
      })
      
      // Simulate AI response
      setTimeout(() => {
        this.conversation.push({
          id: Date.now() + 1,
          type: 'assistant',
          content: 'Based on your query, I recommend checking your server logs and optimizing your database queries.',
          timestamp: new Date()
        })
      }, 1000)
      
      this.chatInput = ''
    },
    applyRecommendation(rec) {
      console.log('Applying recommendation:', rec.title)
      // TODO: Implement recommendation application
    },
    dismissRecommendation(rec) {
      this.recommendations = this.recommendations.filter(r => r.id !== rec.id)
    },
    formatTime(date) {
      return new Date(date).toLocaleTimeString()
    }
  }
}
</script>

<style scoped>
.ai-dashboard {
  padding: 2rem;
}

.ai-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 2rem;
}

.card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 4px 20px rgba(0,0,0,0.1);
}

.copilot-chat {
  height: 300px;
  display: flex;
  flex-direction: column;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  margin-bottom: 1rem;
}

.message {
  margin-bottom: 1rem;
  padding: 0.5rem;
  border-radius: 8px;
}

.message.user {
  background: #667eea;
  color: white;
  margin-left: 2rem;
}

.message.assistant {
  background: #f8f9fa;
  margin-right: 2rem;
}

.chat-input {
  display: flex;
  gap: 0.5rem;
}

.chat-input input {
  flex: 1;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.health-score {
  text-align: center;
  margin-bottom: 1rem;
}

.score-circle {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  background: conic-gradient(#667eea 0deg, #667eea calc(var(--score) * 3.6deg), #e9ecef calc(var(--score) * 3.6deg));
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  margin: 0 auto;
  position: relative;
}

.score-circle::before {
  content: '';
  position: absolute;
  width: 80px;
  height: 80px;
  background: white;
  border-radius: 50%;
}

.score-value {
  font-size: 2rem;
  font-weight: bold;
  color: #667eea;
  z-index: 1;
}

.score-label {
  font-size: 0.8rem;
  color: #666;
  z-index: 1;
}

.prediction-item, .anomaly-item, .model-item {
  display: flex;
  justify-content: space-between;
  padding: 0.5rem 0;
  border-bottom: 1px solid #eee;
}

.status-indicator {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  display: inline-block;
  margin-right: 0.5rem;
}

.status-indicator.normal {
  background: #28a745;
}

.recommendation {
  display: flex;
  gap: 1rem;
  padding: 1rem;
  border: 1px solid #eee;
  border-radius: 8px;
  margin-bottom: 1rem;
}

.rec-icon {
  font-size: 2rem;
}

.rec-actions {
  margin-top: 0.5rem;
  display: flex;
  gap: 0.5rem;
}

.apply-btn, .dismiss-btn {
  padding: 0.25rem 0.75rem;
  border-radius: 4px;
  border: none;
  cursor: pointer;
}

.apply-btn {
  background: #667eea;
  color: white;
}

.dismiss-btn {
  background: #6c757d;
  color: white;
}

.automation-stats {
  display: flex;
  justify-content: space-around;
  text-align: center;
}

.stat-value {
  display: block;
  font-size: 2rem;
  font-weight: bold;
  color: #667eea;
}

.stat-label {
  font-size: 0.8rem;
  color: #666;
}
</style>
