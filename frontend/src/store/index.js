
import { createStore } from 'vuex'
import admin from './admin'
import user from './user'
import reseller from './reseller'
import ai from './ai'
import analytics from './analytics'
import billing from './billing'
import system from './system'

export default createStore({
  state: {
    isAuthenticated: false,
    user: null,
    theme: 'light',
    notifications: []
  },
  mutations: {
    SET_AUTHENTICATION(state, status) {
      state.isAuthenticated = status
    },
    SET_USER(state, user) {
      state.user = user
    },
    SET_THEME(state, theme) {
      state.theme = theme
    },
    ADD_NOTIFICATION(state, notification) {
      state.notifications.push({
        id: Date.now(),
        ...notification
      })
    },
    REMOVE_NOTIFICATION(state, id) {
      state.notifications = state.notifications.filter(n => n.id !== id)
    }
  },
  actions: {
    login({ commit }, credentials) {
      // TODO: Implement login logic
      commit('SET_AUTHENTICATION', true)
      commit('SET_USER', { name: credentials.email })
    },
    logout({ commit }) {
      commit('SET_AUTHENTICATION', false)
      commit('SET_USER', null)
      localStorage.removeItem('token')
    },
    showNotification({ commit }, notification) {
      commit('ADD_NOTIFICATION', notification)
      setTimeout(() => {
        commit('REMOVE_NOTIFICATION', notification.id)
      }, 5000)
    }
  },
  modules: {
    admin,
    user,
    reseller,
    ai,
    analytics,
    billing,
    system
  }
})
