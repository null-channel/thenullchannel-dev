
import Vue from 'vue'
import App from './App.vue'
require('dotenv').config()
import axios from 'axios'


window.axios = axios
axios.defaults.baseURL = process.env.API_URL
Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
