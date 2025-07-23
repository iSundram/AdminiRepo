
<template>
  <div class="user-layout">
    <nav class="sidebar">
      <div class="logo">
        <h2>AdminiPanel</h2>
      </div>
      <ul class="nav-menu">
        <li><router-link to="/user/dashboard">Dashboard</router-link></li>
        <li><router-link to="/user/domains">Domains</router-link></li>
        <li><router-link to="/user/files">Files</router-link></li>
        <li><router-link to="/user/email">Email</router-link></li>
        <li><router-link to="/user/databases">Databases</router-link></li>
        <li><router-link to="/user/applications">Applications</router-link></li>
        <li><router-link to="/user/security">Security</router-link></li>
        <li><router-link to="/user/stats">Statistics</router-link></li>
        <li><router-link to="/user/backups">Backups</router-link></li>
        <li><router-link to="/user/ai">AI Assistant</router-link></li>
      </ul>
    </nav>
    <main class="main-content">
      <header class="top-bar">
        <div class="breadcrumb">
          <span>AdminiPanel</span>
        </div>
        <div class="user-menu">
          <span>{{ user.name || 'User' }}</span>
          <button @click="logout">Logout</button>
        </div>
      </header>
      <div class="content">
        <router-view />
      </div>
    </main>
  </div>
</template>

<script>
export default {
  name: 'UserPanelLayout',
  computed: {
    user() {
      return this.$store.state.user || {}
    }
  },
  methods: {
    logout() {
      this.$store.dispatch('logout')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>
.user-layout {
  display: flex;
  height: 100vh;
}

.sidebar {
  width: 250px;
  background: #3498db;
  color: white;
  padding: 20px 0;
}

.logo {
  text-align: center;
  margin-bottom: 30px;
}

.nav-menu {
  list-style: none;
  padding: 0;
}

.nav-menu li {
  margin: 5px 0;
}

.nav-menu a {
  color: white;
  text-decoration: none;
  padding: 10px 20px;
  display: block;
  transition: background 0.3s;
}

.nav-menu a:hover {
  background: #2980b9;
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.top-bar {
  background: #ecf0f1;
  padding: 15px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #bdc3c7;
}

.content {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
}
</style>
