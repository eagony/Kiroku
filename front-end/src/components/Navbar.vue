<template>
  <nav>
    <!-- 应用栏 -->
    <v-app-bar app elevation="2" color="white">
      <v-app-bar-nav-icon class="grey--text" @click="drawer = !drawer">
      </v-app-bar-nav-icon>

      <v-toolbar-title @click="toIndex">
        <!-- <span class="font-weight-dark">R</span>
        <span class="font-weight-light">interest</span> -->
        <!-- <span class="text-h5 font-weight-dark">Kiroku</span> -->
        <a class="text-decoration-none" href="/">
          <span class="text-h4 font-weight-dark" style="color:teal;"
            >Kiroku</span
          ></a
        >
      </v-toolbar-title>

      <v-spacer></v-spacer>
      <v-toolbar-title v-if="this.$route.path != '/explore'">
        <a class="text-decoration-none mr-5" href="/explore">
          <span class=" text-h6 font-weight-dark" style="color:black;"
            >发现</span
          >
        </a>
      </v-toolbar-title>
      <v-toolbar-title>
        <a
          class="text-decoration-none"
          href="https://github.com/eagony/Rinterest"
        >
          <span class="font-weight-dark" style="color:black;">Github</span></a
        >
      </v-toolbar-title>
    </v-app-bar>

    <!-- 导航抽屉 -->
    <v-navigation-drawer v-model="drawer" app dark color="teal">
      <v-list dense nav>
        <v-list-item two-line>
          <v-list-item-avatar>
            <img
              :src="
                this.$store.state.user.avatar ||
                  'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
              "
            />
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title>{{
              this.$store.state.user.username || '未登录'
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

      <template v-slot:append v-if="this.$store.getters.isLoggedIn">
        <div class="pa-2 mb-5">
          <v-btn block @click="doLogout">注销登录</v-btn>
        </div>
      </template>
      <template v-slot:append v-else>
        <div class="pa-2 mb-5">
          <v-btn block @click="goLogin">前往登录</v-btn>
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
      drawer: false,
      items: [
        { title: '代办', icon: 'mdi-format-list-bulleted', route: '/todo' },
        { title: '日记', icon: 'mdi-book-open-blank-variant', route: '/diary' },
        { title: '博客', icon: 'mdi-post-outline', route: '/blog' },
        // { title: '文件', icon: 'mdi-file', route: '/file' },
        // { title: '管理', icon: 'mdi-chart-donut', route: '/admin' },
        { title: '设置', icon: 'mdi-cog-outline', route: '/settings' },
        { title: '关于', icon: 'mdi-help-box', route: '/about' }
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
    },
    goLogin() {
      this.$router.push('/login');
    }
  }
};
</script>
