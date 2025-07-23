
<template>
  <div class="admin-layout">
    <nav class="sidebar">
      <div class="logo">
        <h2>AdminiCore</h2>
      </div>
      <ul class="nav-menu">
        <li><router-link to="/admin/dashboard">Dashboard</router-link></li>
        <li><router-link to="/admin/accounts">Accounts</router-link></li>
        <li><router-link to="/admin/dns">DNS</router-link></li>
        <li><router-link to="/admin/email">Email</router-link></li>
        <li><router-link to="/admin/files">Files</router-link></li>
        <li><router-link to="/admin/security">Security</router-link></li>
        <li><router-link to="/admin/monitoring">Monitoring</router-link></li>
        <li><router-link to="/admin/ai">AI Assistant</router-link></li>
        <li><router-link to="/admin/system">System</router-link></li>
      </ul>
    </nav>
    <main class="main-content">
      <header class="top-bar">
        <div class="breadcrumb">
          <span>AdminiCore</span>
        </div>
        <div class="user-menu">
          <span>Admin User</span>
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
  name: 'AdminCoreLayout',
  methods: {
    logout() {
      this.$store.dispatch('logout')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>
.admin-layout {
  display: flex;
  height: 100vh;
}

.sidebar {
  width: 250px;
  background: #2c3e50;
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
  background: #34495e;
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
