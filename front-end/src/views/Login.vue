<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <v-card class="elevation-12">
          <v-img src="../assets/cover1.jpg" max-height="250"></v-img>
          <v-toolbar class="pt-3" color="white" light flat>
            <v-toolbar-title>登陆到 <strong>Kiroku</strong></v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn text large @click="goRegister">注册</v-btn>
          </v-toolbar>
          <v-card-text @keydown.enter="doLogin">
            <v-form>
              <v-text-field
                label="Login"
                name="login"
                prepend-icon="mdi-account"
                type="text"
                v-model="username"
              ></v-text-field>

              <v-text-field
                id="password"
                label="Password"
                name="password"
                prepend-icon="mdi-lock"
                type="password"
                v-model="password"
              ></v-text-field>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn block dark color="teal" @click="doLogin">登陆</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Toast from '../plugins/toast';

export default {
  name: 'Login',
  props: {
    source: String
  },
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    doLogin() {
      const path = '/users/login';
      this.$axios
        .post(path, {
          username: this.username,
          password: this.password
        })
        .then(res => {
          window.localStorage.setItem('r-token', res.data.data);
          let user = JSON.parse(atob(res.data.data.split('.')[1]));
          this.$store.commit('login', user);
          this.$router.push('/todo');
          Toast.fire({
            icon: 'success',
            title: '登录成功，开始美好一天。'
          });
        })
        .catch(err => {
          alert(err);
        });
    },
    goRegister() {
      this.$router.push('/register');
    }
  }
};
</script>
