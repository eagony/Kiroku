<template>
  <v-app class="grey lighten-4">
    <Navbar
      v-if="!(this.$route.path == '/login' || this.$route.path == '/register')"
    ></Navbar>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
import Navbar from '@/components/Navbar';
export default {
  name: 'App',
  components: { Navbar },
  data: () => ({
    //
  })
};
</script>
