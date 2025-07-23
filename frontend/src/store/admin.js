
export default {
  namespaced: true,
  state: {
    accounts: [],
    systemInfo: {},
    activities: [],
    settings: {}
  },
  mutations: {
    SET_ACCOUNTS(state, accounts) {
      state.accounts = accounts
    },
    SET_SYSTEM_INFO(state, info) {
      state.systemInfo = info
    },
    SET_ACTIVITIES(state, activities) {
      state.activities = activities
    },
    SET_SETTINGS(state, settings) {
      state.settings = settings
    }
  },
  actions: {
    async fetchAccounts({ commit }) {
      // TODO: Implement API call
      const accounts = [
        { id: 1, username: 'user1', domain: 'example1.com', package: 'Starter' },
        { id: 2, username: 'user2', domain: 'example2.com', package: 'Pro' }
      ]
      commit('SET_ACCOUNTS', accounts)
    },
    async fetchSystemInfo({ commit }) {
      // TODO: Implement API call
      const info = {
        cpuUsage: 45,
        memoryUsage: 67,
        diskUsage: 23,
        uptime: '15 days'
      }
      commit('SET_SYSTEM_INFO', info)
    }
  }
}
