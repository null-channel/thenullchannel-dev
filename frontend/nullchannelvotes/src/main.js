import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'

Vue.config.productionTip = false

window.axios = axios
axios.defaults.baseURL = process.env.API_URL
new Vue({
  render: h => h(App),
}).$mount('#app')
