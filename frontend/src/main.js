
import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import { createStore } from 'vuex'
import App from './App.vue'

// Router configuration
const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: () => import('./views/Dashboard.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('./views/Login.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Store configuration
const store = createStore({
  state: {
    user: null,
    isAuthenticated: false,
    theme: 'light'
  },
  mutations: {
    SET_USER(state, user) {
      state.user = user
      state.isAuthenticated = !!user
    },
    SET_THEME(state, theme) {
      state.theme = theme
    }
  },
  actions: {
    login({ commit }, user) {
      commit('SET_USER', user)
    },
    logout({ commit }) {
      commit('SET_USER', null)
    }
  }
})

// Create and mount app
const app = createApp(App)
app.use(router)
app.use(store)
app.mount('#app')
