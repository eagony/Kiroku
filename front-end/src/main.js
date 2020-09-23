import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import vuetify from './plugins/vuetify';
import 'roboto-fontface/css/roboto/roboto-fontface.css';
import '@mdi/font/css/materialdesignicons.css';

// axios
import axios from './http';
Vue.prototype.$axios = axios;

// mavon-editor
import mavonEditor from 'mavon-editor';
import 'mavon-editor/dist/css/index.css';
Vue.use(mavonEditor);

// highlight主题
import 'highlight.js/styles/atom-one-dark-reasonable.css';

// github-markdown样式
import 'github-markdown-css/github-markdown.css';

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app');
