<template>
  <v-app class="grey lighten-4">
    <Navbar
      v-if="!(this.$route.path == '/login' || this.$route.path == '/register')"
    ></Navbar>
    <v-main>
      <router-view></router-view>
    </v-main>
    <Footer
      v-if="!(this.$route.path == '/login' || this.$route.path == '/register')"
    ></Footer>
  </v-app>
</template>

<script>
import Navbar from '@/components/Navbar';
import Footer from '@/components/Footer';

export default {
  name: 'App',
  components: { Navbar, Footer },
  data: () => ({
    //
  })
};
</script>
