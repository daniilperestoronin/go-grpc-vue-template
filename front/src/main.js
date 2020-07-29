import Vue from 'vue'
import VueRouter from 'vue-router'
import vuetify from './plugins/vuetify';

import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@mdi/font/css/materialdesignicons.css'

import App from './App.vue'
import router from './plugins/router'

Vue.config.productionTip = false
Vue.use(VueRouter);

import VueLayers from 'vuelayers'
import 'vuelayers/lib/style.css' // needs css-loader

Vue.use(VueLayers)

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
