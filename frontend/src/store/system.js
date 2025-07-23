
export default {
  namespaced: true,
  state: {
    info: {},
    logs: [],
    metrics: {},
    services: []
  },
  mutations: {
    SET_INFO(state, info) {
      state.info = info
    },
    SET_LOGS(state, logs) {
      state.logs = logs
    },
    SET_METRICS(state, metrics) {
      state.metrics = metrics
    },
    SET_SERVICES(state, services) {
      state.services = services
    }
  },
  actions: {
    async fetchSystemInfo({ commit }) {
      // TODO: Implement API call
      const info = {
        version: '1.0.0',
        os: 'Linux',
        uptime: '15 days, 3 hours',
        load: [0.5, 0.6, 0.7]
      }
      commit('SET_INFO', info)
    },
    async fetchLogs({ commit }) {
      // TODO: Implement API call
      const logs = [
        { id: 1, level: 'info', message: 'System started', timestamp: new Date() },
        { id: 2, level: 'warning', message: 'High memory usage', timestamp: new Date() }
      ]
      commit('SET_LOGS', logs)
    }
  }
}
