
<template>
  <div id="app" :class="[theme]">
    <header class="header">
      <div class="container">
        <div class="logo">
          <h1>AdminiSoftware</h1>
          <span>Next-Gen Control Panel</span>
        </div>
        <nav class="nav">
          <router-link to="/">Dashboard</router-link>
          <router-link to="/login">Login</router-link>
        </nav>
      </div>
    </header>
    
    <main class="main">
      <router-view />
    </main>
    
    <footer class="footer">
      <div class="container">
        <p>&copy; 2024 AdminiSoftware - <a href="https://admini.tech">admini.tech</a> | Support: <a href="mailto:support@admini.tech">support@admini.tech</a></p>
      </div>
    </footer>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'App',
  setup() {
    const store = useStore()
    const theme = computed(() => store.state.theme)
    
    return {
      theme
    }
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  line-height: 1.6;
  color: #333;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 1rem 0;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

.header .container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo h1 {
  font-size: 1.8rem;
  margin-bottom: 0.2rem;
}

.logo span {
  font-size: 0.9rem;
  opacity: 0.9;
}

.nav {
  display: flex;
  gap: 2rem;
}

.nav a {
  color: white;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 5px;
  transition: background-color 0.3s;
}

.nav a:hover {
  background-color: rgba(255,255,255,0.2);
}

.main {
  min-height: calc(100vh - 140px);
  padding: 2rem 0;
}

.footer {
  background: #2c3e50;
  color: white;
  text-align: center;
  padding: 1rem 0;
}

.footer a {
  color: #3498db;
  text-decoration: none;
}

.footer a:hover {
  text-decoration: underline;
}
</style>
