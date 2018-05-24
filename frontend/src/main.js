import Vue from 'vue'
import App from './App.vue'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'

require('./assets/sass/main.scss');

Vue.config.productionTip = false
Vue.use(Vuetify)

new Vue({
  render: h => h(App)
}).$mount('#app')
