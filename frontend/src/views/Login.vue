
<template>
  <div class="login">
    <div class="container">
      <div class="login-form">
        <h2>Login to AdminiSoftware</h2>
        <form @submit.prevent="handleLogin">
          <div class="form-group">
            <label for="email">Email:</label>
            <input 
              type="email" 
              id="email" 
              v-model="form.email" 
              required
              placeholder="support@admini.tech"
            />
          </div>
          
          <div class="form-group">
            <label for="password">Password:</label>
            <input 
              type="password" 
              id="password" 
              v-model="form.password" 
              required
              placeholder="Enter your password"
            />
          </div>
          
          <button type="submit" class="login-btn">Login</button>
        </form>
        
        <div class="login-footer">
          <p>Demo credentials: admin@admini.tech / admin123</p>
          <p>Visit: <a href="https://admini.tech" target="_blank">admini.tech</a></p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

export default {
  name: 'Login',
  setup() {
    const store = useStore()
    const router = useRouter()
    
    const form = ref({
      email: '',
      password: ''
    })
    
    const handleLogin = () => {
      // Demo login logic
      if (form.value.email && form.value.password) {
        store.dispatch('login', {
          email: form.value.email,
          name: 'Admin User'
        })
        router.push('/')
      }
    }
    
    return {
      form,
      handleLogin
    }
  }
}
</script>

<style scoped>
.login {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 60vh;
}

.login-form {
  background: white;
  padding: 2rem;
  border-radius: 10px;
  box-shadow: 0 5px 15px rgba(0,0,0,0.1);
  width: 100%;
  max-width: 400px;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
}

.form-group input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 5px;
  font-size: 1rem;
}

.login-btn {
  width: 100%;
  padding: 0.75rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 5px;
  font-size: 1rem;
  cursor: pointer;
  transition: opacity 0.3s;
}

.login-btn:hover {
  opacity: 0.9;
}

.login-footer {
  margin-top: 1.5rem;
  text-align: center;
  font-size: 0.9rem;
  color: #666;
}

.login-footer a {
  color: #667eea;
  text-decoration: none;
}
</style>
