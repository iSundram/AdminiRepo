
<template>
  <div class="experimental-layout">
    <nav class="sidebar">
      <div class="logo">
        <h2>Experimental</h2>
        <span class="beta-badge">BETA</span>
      </div>
      <ul class="nav-menu">
        <li><router-link to="/experimental/blockchain-dns">Blockchain DNS</router-link></li>
        <li><router-link to="/experimental/quantum-tls">Quantum TLS</router-link></li>
        <li><router-link to="/experimental/ar-vr">AR/VR Dashboard</router-link></li>
        <li><router-link to="/experimental/voice-assistant">Voice Assistant</router-link></li>
        <li><router-link to="/experimental/web3-wallet">Web3 Wallet</router-link></li>
      </ul>
    </nav>
    <main class="main-content">
      <header class="top-bar">
        <div class="warning-banner">
          ⚠️ Experimental Features - Use with caution in production environments
        </div>
        <div class="user-menu">
          <button @click="goBack">← Back to Main Panel</button>
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
  name: 'ExperimentalLayout',
  methods: {
    goBack() {
      this.$router.push('/dashboard')
    }
  }
}
</script>

<style scoped>
.experimental-layout {
  display: flex;
  height: 100vh;
}

.sidebar {
  width: 250px;
  background: linear-gradient(135deg, #8e2de2 0%, #4a00e0 100%);
  color: white;
  padding: 20px 0;
}

.logo {
  text-align: center;
  margin-bottom: 30px;
  position: relative;
}

.beta-badge {
  background: #ff6b6b;
  padding: 2px 8px;
  border-radius: 10px;
  font-size: 0.7em;
  position: absolute;
  top: -5px;
  right: 20px;
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
  background: rgba(255, 255, 255, 0.1);
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.top-bar {
  background: #fff3cd;
  padding: 15px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #ffeaa7;
}

.warning-banner {
  color: #856404;
  font-weight: bold;
}

.content {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
}
</style>
