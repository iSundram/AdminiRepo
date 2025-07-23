
<template>
  <div class="auth-layout">
    <div class="auth-container">
      <div class="auth-header">
        <h1>AdminiSoftware</h1>
        <p>Next-Generation Control Panel</p>
      </div>
      <div class="auth-content">
        <router-view />
      </div>
      <div class="auth-footer">
        <p>&copy; 2024 AdminiSoftware. All rights reserved.</p>
        <p><a href="https://admini.tech">admini.tech</a></p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AuthLayout'
}
</script>

<style scoped>
.auth-layout {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.auth-container {
  background: white;
  border-radius: 10px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  overflow: hidden;
  min-width: 400px;
  max-width: 500px;
  width: 100%;
}

.auth-header {
  background: #2c3e50;
  color: white;
  padding: 30px;
  text-align: center;
}

.auth-header h1 {
  margin: 0 0 10px 0;
  font-size: 2em;
}

.auth-header p {
  margin: 0;
  opacity: 0.8;
}

.auth-content {
  padding: 30px;
}

.auth-footer {
  background: #ecf0f1;
  padding: 20px;
  text-align: center;
  font-size: 0.9em;
  color: #7f8c8d;
}

.auth-footer a {
  color: #3498db;
  text-decoration: none;
}
</style>
