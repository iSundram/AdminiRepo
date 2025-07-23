
export default {
  namespaced: true,
  state: {
    invoices: [],
    payments: [],
    plans: [],
    currentPlan: null,
    usage: {},
    loading: false
  },
  mutations: {
    SET_INVOICES(state, invoices) {
      state.invoices = invoices
    },
    SET_PAYMENTS(state, payments) {
      state.payments = payments
    },
    SET_PLANS(state, plans) {
      state.plans = plans
    },
    SET_CURRENT_PLAN(state, plan) {
      state.currentPlan = plan
    },
    SET_USAGE(state, usage) {
      state.usage = usage
    },
    SET_LOADING(state, loading) {
      state.loading = loading
    }
  },
  actions: {
    async fetchInvoices({ commit }) {
      commit('SET_LOADING', true)
      try {
        // TODO: Implement API call
        const invoices = []
        commit('SET_INVOICES', invoices)
      } catch (error) {
        console.error('Failed to fetch invoices:', error)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async fetchPlans({ commit }) {
      commit('SET_LOADING', true)
      try {
        // TODO: Implement API call
        const plans = [
          { id: 1, name: 'Basic', price: 9.99 },
          { id: 2, name: 'Pro', price: 29.99 },
          { id: 3, name: 'Enterprise', price: 99.99 }
        ]
        commit('SET_PLANS', plans)
      } catch (error) {
        console.error('Failed to fetch plans:', error)
      } finally {
        commit('SET_LOADING', false)
      }
    }
  },
  getters: {
    overdueInvoices: state => {
      return state.invoices.filter(invoice => 
        invoice.status === 'overdue'
      )
    },
    totalOutstanding: state => {
      return state.invoices
        .filter(invoice => invoice.status !== 'paid')
        .reduce((total, invoice) => total + invoice.amount, 0)
    }
  }
}
