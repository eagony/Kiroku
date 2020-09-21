<template>
  <nav>
    <!-- 消息条 -->
    <v-snackbar v-model="snackbar" :timeout="4000" top color="success">
      <span>Wow! You added a new project.</span>
      <v-btn text color="white" @click="snackbar = false">Close</v-btn>
    </v-snackbar>

    <!-- 应用栏 -->
    <v-app-bar app elevation="2" color="white">
      <v-app-bar-nav-icon class="grey--text" @click="drawer = !drawer">
      </v-app-bar-nav-icon>

      <v-toolbar-title class="grey--text" @click="toIndex">
        <span class="font-weight-dark">R</span>
        <span class="font-weight-light">interest</span>
      </v-toolbar-title>

      <v-spacer></v-spacer>

      <v-badge color="green" overlap>
        <span slot="badge">
          1
        </span>
        <v-icon medium>notifications</v-icon>
      </v-badge>
    </v-app-bar>

    <!-- 导航抽屉 -->
    <v-navigation-drawer v-model="drawer" app dark color="teal">
      <v-list dense nav>
        <v-list-item two-line link>
          <v-list-item-avatar>
            <img src="https://randomuser.me/api/portraits/men/81.jpg" />
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title>{{
              this.$store.state.user.username
            }}</v-list-item-title>
            <v-list-item-subtitle>{{
              this.$store.state.user.signature
            }}</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-divider></v-divider>

        <v-list-item
          v-for="item in items"
          :key="item.title"
          link
          style="margin-top: 5px;"
          :to="item.route"
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <template v-slot:append>
        <div class="pa-2">
          <v-btn block @click="doLogout">注销登录</v-btn>
        </div>
      </template>
    </v-navigation-drawer>
  </nav>
</template>

<script>
export default {
  name: 'Navbar',
  data() {
    return {
      messages: 10,
      show: true,
      snackbar: false,
      drawer: false,
      items: [
        { title: '代办', icon: 'dashboard', route: '/' },
        { title: '日记', icon: 'account_box', route: '/diary' },
        { title: '博客', icon: 'mdi-image', route: '/blog' },
        { title: '文件', icon: 'mdi-file', route: '/file' },
        { title: 'Admin', icon: 'gavel' },
        { icon: 'settings', title: 'Settings', route: '/settings' },
        { title: 'About', icon: 'mdi-help-box' }
      ]
    };
  },
  methods: {
    toIndex() {
      if (this.$route.path !== '/') {
        this.$router.push('/');
      }
    },
    doLogout() {
      this.$store.commit('logout');
      window.localStorage.removeItem('r-token');
      this.$router.push('/login');
    }
  }
};
</script>
