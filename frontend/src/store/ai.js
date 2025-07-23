
import { createStore } from 'vuex'

export default {
  namespaced: true,
  state: {
    copilot: {
      isActive: false,
      suggestions: [],
      conversation: []
    },
    predictions: {
      resource: null,
      failures: null,
      security: null
    },
    models: [],
    systemAnalysis: null,
    anomalies: []
  },
  mutations: {
    SET_COPILOT_ACTIVE(state, isActive) {
      state.copilot.isActive = isActive
    },
    ADD_COPILOT_MESSAGE(state, message) {
      state.copilot.conversation.push(message)
    },
    SET_PREDICTIONS(state, { type, data }) {
      state.predictions[type] = data
    },
    SET_MODELS(state, models) {
      state.models = models
    },
    SET_SYSTEM_ANALYSIS(state, analysis) {
      state.systemAnalysis = analysis
    },
    ADD_ANOMALY(state, anomaly) {
      state.anomalies.unshift(anomaly)
    }
  },
  actions: {
    async askCopilot({ commit }, query) {
      commit('ADD_COPILOT_MESSAGE', { type: 'user', content: query, timestamp: new Date() })
      
      // TODO: Call AI API
      const response = "This is an AI-powered suggestion based on your query."
      
      commit('ADD_COPILOT_MESSAGE', { 
        type: 'assistant', 
        content: response, 
        timestamp: new Date() 
      })
    },
    async fetchPredictions({ commit }, type) {
      // TODO: Call prediction API
      const mockData = {
        cpu: 85.5,
        memory: 67.2,
        storage: 45.8,
        confidence: 0.92
      }
      commit('SET_PREDICTIONS', { type, data: mockData })
    },
    async analyzeSystem({ commit }) {
      // TODO: Call system analysis API
      const analysis = {
        healthScore: 95,
        recommendations: [
          'Consider upgrading PHP version for better performance',
          'SSL certificate expires in 30 days'
        ]
      }
      commit('SET_SYSTEM_ANALYSIS', analysis)
    }
  },
  getters: {
    healthScore: state => state.systemAnalysis?.healthScore || 0,
    recentAnomalies: state => state.anomalies.slice(0, 5),
    activePredictions: state => Object.values(state.predictions).filter(p => p !== null)
  }
}
