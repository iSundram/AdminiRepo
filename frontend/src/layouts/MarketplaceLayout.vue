
<template>
  <div class="marketplace-layout">
    <header class="marketplace-header">
      <div class="container">
        <h1>AdminiSoftware Marketplace</h1>
        <nav>
          <router-link to="/marketplace">Home</router-link>
          <router-link to="/marketplace/plugins">Plugins</router-link>
          <router-link to="/marketplace/themes">Themes</router-link>
          <router-link to="/marketplace/apps">Apps</router-link>
        </nav>
      </div>
    </header>
    <main class="marketplace-content">
      <div class="container">
        <router-view />
      </div>
    </main>
    <footer class="marketplace-footer">
      <div class="container">
        <p>&copy; 2024 AdminiSoftware Marketplace</p>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'MarketplaceLayout'
}
</script>

<style scoped>
.marketplace-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.marketplace-header {
  background: #2c3e50;
  color: white;
  padding: 20px 0;
}

.marketplace-header .container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.marketplace-header nav a {
  color: white;
  text-decoration: none;
  margin: 0 15px;
  padding: 10px 15px;
  border-radius: 5px;
  transition: background 0.3s;
}

.marketplace-header nav a:hover {
  background: #34495e;
}

.marketplace-content {
  flex: 1;
  padding: 40px 0;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.marketplace-footer {
  background: #ecf0f1;
  padding: 20px 0;
  text-align: center;
  color: #7f8c8d;
}
</style>
