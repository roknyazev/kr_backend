import Vue from 'vue'
import vuetify from '@/plugins/vuetify' // path to vuetify export
import VueMobileDetection from 'vue-mobile-detection'
import loader from "vue-ui-preloader"

import App from './App.vue'

Vue.config.productionTip = false

Vue.use(VueMobileDetection)
Vue.use(loader)

new Vue({
  vuetify,
  render: h => h(App),
  loader: loader
}).$mount('#app')
