
export default {
  namespaced: true,
  state: {
    domains: [],
    emails: [],
    files: [],
    databases: [],
    stats: {}
  },
  mutations: {
    SET_DOMAINS(state, domains) {
      state.domains = domains
    },
    SET_EMAILS(state, emails) {
      state.emails = emails
    },
    SET_FILES(state, files) {
      state.files = files
    },
    SET_DATABASES(state, databases) {
      state.databases = databases
    },
    SET_STATS(state, stats) {
      state.stats = stats
    }
  },
  actions: {
    async fetchDomains({ commit }) {
      // TODO: Implement API call
      const domains = [
        { id: 1, name: 'example.com', type: 'main', status: 'active' },
        { id: 2, name: 'subdomain.example.com', type: 'subdomain', status: 'active' }
      ]
      commit('SET_DOMAINS', domains)
    },
    async fetchEmails({ commit }) {
      // TODO: Implement API call
      const emails = [
        { id: 1, address: 'admin@example.com', quota: 1024, used: 256 },
        { id: 2, address: 'info@example.com', quota: 512, used: 128 }
      ]
      commit('SET_EMAILS', emails)
    }
  }
}
