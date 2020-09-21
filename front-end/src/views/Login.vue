<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <!--        <v-card class="elevation-1 pa-3">-->
        <!--          <v-card-text>-->
        <!--            <div class="layout column align-center">-->
        <!--&lt;!&ndash;              <img src="static/logo.png" alt="Vue Material Admin" width="180" height="180">&ndash;&gt;-->
        <!--              <h1 class="flex my-4 primary&#45;&#45;text">登陆到Rinterest</h1>-->
        <!--            </div>-->
        <!--            <v-form>-->
        <!--              <v-text-field-->
        <!--                append-icon="person"-->
        <!--                name="login"-->
        <!--                label="Username"-->
        <!--                type="text"-->
        <!--                v-model="userEmail"-->
        <!--                :error="error"-->
        <!--                :rules="[rules.required]"/>-->
        <!--              <v-text-field-->
        <!--                :type="hidePassword ? 'password' : 'text'"-->
        <!--                :append-icon="hidePassword ? 'visibility_off' : 'visibility'"-->
        <!--                name="password"-->
        <!--                label="Password"-->
        <!--                id="password"-->
        <!--                :rules="[rules.required]"-->
        <!--                v-model="password"-->
        <!--                :error="error"-->
        <!--                @click:append="hidePassword = !hidePassword"/>-->
        <!--            </v-form>-->
        <!--          </v-card-text>-->
        <!--          <v-card-actions>-->
        <!--            <v-spacer></v-spacer>-->
        <!--            <v-btn block color="primary" @click="login" :loading="loading">登陆</v-btn>-->
        <!--          </v-card-actions>-->
        <!--        </v-card>-->

        <v-card class="elevation-12">
          <v-toolbar color="teal" dark flat>
            <v-toolbar-title>登陆到Rinterest</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-tooltip bottom>
              <template v-slot:activator="{ on }">
                <v-btn :href="source" icon large target="_blank" v-on="on">
                  <!--                  <v-icon>mdi-code-tags</v-icon>-->
                  <span>注册</span>
                </v-btn>
              </template>
              <span>Source</span>
            </v-tooltip>
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
            <v-spacer></v-spacer>
            <v-btn block color="teal" @click="doLogin">登陆</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
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
      const path = '/api/v1/users/login';
      this.$axios
        .post(path, {
          username: this.username,
          password: this.password
        })
        .then(res => {
          window.localStorage.setItem('r-token', res.data.data);
          let user = JSON.parse(atob(res.data.data.split('.')[1]));
          this.$store.commit('login', user);
          console.log(user);
          this.$router.push('/');
        })
        .catch(err => {
          alert(err);
        });
    }
  }
};
</script>
