import Vue from 'vue';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import Vuelidate from 'vuelidate';
// import VueAxios from 'vue-axios';
import axios from 'axios';
import App from './App.vue';
import router from './router';
import store from './store';

// Import Bootstrap and BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import './assets/scss/index.scss';

Vue.config.productionTip = false;

// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);

Vue.use(Vuelidate);

Vue.prototype.axios = axios;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
