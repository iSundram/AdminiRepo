
export default {
  namespaced: true,
  state: {
    clients: [],
    packages: [],
    commissions: [],
    stats: {}
  },
  mutations: {
    SET_CLIENTS(state, clients) {
      state.clients = clients
    },
    SET_PACKAGES(state, packages) {
      state.packages = packages
    },
    SET_COMMISSIONS(state, commissions) {
      state.commissions = commissions
    },
    SET_STATS(state, stats) {
      state.stats = stats
    }
  },
  actions: {
    async fetchClients({ commit }) {
      // TODO: Implement API call
      const clients = [
        { id: 1, username: 'client1', domain: 'client1.com', package: 'Basic' },
        { id: 2, username: 'client2', domain: 'client2.com', package: 'Premium' }
      ]
      commit('SET_CLIENTS', clients)
    },
    async fetchPackages({ commit }) {
      // TODO: Implement API call
      const packages = [
        { id: 1, name: 'Basic', price: 9.99, accounts: 5 },
        { id: 2, name: 'Premium', price: 19.99, accounts: 15 }
      ]
      commit('SET_PACKAGES', packages)
    }
  }
}
