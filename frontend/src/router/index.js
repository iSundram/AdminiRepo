
import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Dashboard from '../views/Dashboard.vue'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: { requiresAuth: true }
  },
  {
    path: '/admin',
    name: 'AdminCore',
    component: () => import('../views/AdminiCore/Dashboard.vue'),
    meta: { requiresAuth: true, role: 'admin' }
  },
  {
    path: '/reseller',
    name: 'AdminReseller',
    component: () => import('../views/AdminiReseller/Dashboard.vue'),
    meta: { requiresAuth: true, role: 'reseller' }
  },
  {
    path: '/panel',
    name: 'AdminPanel',
    component: () => import('../views/AdminiPanel/Dashboard.vue'),
    meta: { requiresAuth: true, role: 'user' }
  },
  {
    path: '/marketplace',
    name: 'Marketplace',
    component: () => import('../views/Marketplace/Index.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/experimental',
    name: 'Experimental',
    component: () => import('../views/Experimental/BlockchainDNS.vue'),
    meta: { requiresAuth: true, role: 'admin' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem('token')
  
  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
  } else {
    next()
  }
})

export default router
