import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    logged: !!window.localStorage.getItem('r-token'),
    user: {
      id: window.localStorage.getItem('r-token')
        ? JSON.parse(atob(window.localStorage.getItem('r-token').split('.')[1]))
            .id
        : 0,
      role: window.localStorage.getItem('r-token')
        ? JSON.parse(atob(window.localStorage.getItem('r-token').split('.')[1]))
            .role
        : 'user',
      email: window.localStorage.getItem('r-token')
        ? JSON.parse(atob(window.localStorage.getItem('r-token').split('.')[1]))
            .email
        : '',
      phone: window.localStorage.getItem('r-token')
        ? JSON.parse(atob(window.localStorage.getItem('r-token').split('.')[1]))
            .phone
        : '',
      avatar: window.localStorage.getItem('r-token')
        ? JSON.parse(atob(window.localStorage.getItem('r-token').split('.')[1]))
            .avatar
        : 0,
      username: window.localStorage.getItem('r-token')
        ? JSON.parse(atob(window.localStorage.getItem('r-token').split('.')[1]))
            .username
        : '',
      signature: window.localStorage.getItem('r-token')
        ? JSON.parse(atob(window.localStorage.getItem('r-token').split('.')[1]))
            .signature
        : ''
    }
  },
  getters: {
    isLoggedIn: state => state.logged,
    isAdmin: state => state.user.role === 'admin'
  },
  mutations: {
    login(state, user) {
      state.logged = true;
      state.user.id = user.id;
      state.user.role = user.role;
      state.user.phone = user.phone;
      state.user.email = user.email;
      state.user.avatar = user.avatar;
      state.user.username = user.username;
      state.user.signature = user.signature;
    },
    logout(state) {
      state.logged = false;
      state.user.id = '';
      state.user.role = '';
      state.user.phone = '';
      state.user.email = '';
      state.user.avatar = '';
      state.user.username = '';
      state.user.signature = '';
    }
  },
  actions: {},
  modules: {}
});
