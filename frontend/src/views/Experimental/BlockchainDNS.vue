
<template>
  <div class="blockchain-dns">
    <div class="header">
      <h1>🔗 Blockchain DNS</h1>
      <p>Experimental decentralized domain name system</p>
      <div class="status-badge experimental">
        <span>🧪 Experimental Feature</span>
      </div>
    </div>

    <div class="dns-grid">
      <div class="card">
        <h3>Register Domain</h3>
        <form @submit.prevent="registerDomain">
          <input 
            v-model="newDomain" 
            placeholder="mydomain.block" 
            type="text"
            class="domain-input"
          />
          <button type="submit" :disabled="isRegistering">
            {{ isRegistering ? 'Registering...' : 'Register on Blockchain' }}
          </button>
        </form>
      </div>

      <div class="card">
        <h3>My Blockchain Domains</h3>
        <div class="domain-list">
          <div v-for="domain in domains" :key="domain.id" class="domain-item">
            <span class="domain-name">{{ domain.name }}</span>
            <span class="domain-status" :class="domain.status">{{ domain.status }}</span>
            <button @click="viewTransaction(domain.txHash)">View TX</button>
          </div>
        </div>
      </div>

      <div class="card">
        <h3>Network Stats</h3>
        <div class="stats">
          <div class="stat">
            <span class="label">Total Domains</span>
            <span class="value">{{ networkStats.totalDomains }}</span>
          </div>
          <div class="stat">
            <span class="label">Pending Transactions</span>
            <span class="value">{{ networkStats.pendingTx }}</span>
          </div>
          <div class="stat">
            <span class="label">Network Fee</span>
            <span class="value">{{ networkStats.fee }} ETH</span>
          </div>
        </div>
      </div>

      <div class="card">
        <h3>Recent Transactions</h3>
        <div class="transactions">
          <div v-for="tx in recentTransactions" :key="tx.hash" class="transaction">
            <span class="tx-hash">{{ tx.hash.substring(0, 10) }}...</span>
            <span class="tx-domain">{{ tx.domain }}</span>
            <span class="tx-status" :class="tx.status">{{ tx.status }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BlockchainDNS',
  data() {
    return {
      newDomain: '',
      isRegistering: false,
      domains: [
        { id: 1, name: 'mysite.block', status: 'confirmed', txHash: '0x1234...' },
        { id: 2, name: 'blog.block', status: 'pending', txHash: '0x5678...' }
      ],
      networkStats: {
        totalDomains: 1337,
        pendingTx: 42,
        fee: 0.001
      },
      recentTransactions: [
        { hash: '0x1234567890', domain: 'test.block', status: 'confirmed' },
        { hash: '0x0987654321', domain: 'demo.block', status: 'pending' }
      ]
    }
  },
  methods: {
    async registerDomain() {
      if (!this.newDomain) return
      
      this.isRegistering = true
      try {
        // TODO: Implement blockchain domain registration
        await new Promise(resolve => setTimeout(resolve, 2000))
        
        this.domains.push({
          id: Date.now(),
          name: this.newDomain,
          status: 'pending',
          txHash: '0x' + Math.random().toString(16).substr(2, 8) + '...'
        })
        
        this.newDomain = ''
      } catch (error) {
        console.error('Registration failed:', error)
      } finally {
        this.isRegistering = false
      }
    },
    viewTransaction(txHash) {
      // TODO: Open blockchain explorer
      window.open(`https://etherscan.io/tx/${txHash}`, '_blank')
    }
  }
}
</script>

<style scoped>
.blockchain-dns {
  padding: 2rem;
}

.header {
  margin-bottom: 2rem;
  text-align: center;
}

.status-badge.experimental {
  display: inline-block;
  background: linear-gradient(45deg, #ff6b35, #f7931e);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-size: 0.9rem;
  margin-top: 1rem;
}

.dns-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.card {
  background: white;
  border-radius: 8px;
  padding: 1.5rem;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

.domain-input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  margin-bottom: 1rem;
}

.domain-item, .transaction {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem 0;
  border-bottom: 1px solid #eee;
}

.domain-status, .tx-status {
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
}

.domain-status.confirmed, .tx-status.confirmed {
  background: #d4edda;
  color: #155724;
}

.domain-status.pending, .tx-status.pending {
  background: #fff3cd;
  color: #856404;
}

.stats {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.stat {
  display: flex;
  justify-content: space-between;
}

.value {
  font-weight: bold;
  color: #667eea;
}
</style>
