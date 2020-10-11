<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" md="4" xl="3">
        <v-card class="elevation-12">
          <v-img src="../assets/cv1.png" max-height="250"></v-img>
          <v-toolbar class="pt-3" color="white" light flat>
            <v-toolbar-title>账号注册</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn text large @click="goLogin">登录</v-btn>
          </v-toolbar>
          <v-card-text>
            <v-form
              ref="form"
              v-model="valid"
              @keydown.enter="doLogin"
              lazy-validation
            >
              <v-text-field
                type="text"
                :counter="16"
                name="username"
                v-model="username"
                :rules="usernameRules"
                prepend-icon="mdi-account"
                label="输入用户名(英文加数字)"
                required
              ></v-text-field>

              <v-text-field
                :counter="16"
                id="password"
                type="password"
                label="输入密码"
                name="password"
                v-model="password"
                :rules="passwordRules"
                prepend-icon="mdi-lock"
                required
              ></v-text-field>

              <v-text-field
                :counter="16"
                type="password"
                label="重复密码"
                id="password-repeat"
                name="password-repeat"
                prepend-icon="mdi-lock"
                v-model="passwordRepeat"
                :rules="passwordRepeatRules"
                required
              ></v-text-field>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn block dark color="deep-purple" @click="doRegister"
              >注册</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Toast from '../plugins/toast';

export default {
  name: 'Register',
  data() {
    return {
      valid: false,
      snackbar: false,

      username: '',
      usernameRules: [
        v => !!v || '用户名不能为空。',
        v => (v && v.length >= 4) || '用户名不能小于四位。',
        v => (v && v.length <= 16) || '用户名太长啦，你记得住嘛？'
      ],
      password: '',
      passwordRules: [
        v => !!v || '密码不能为空。',
        v => (v && v.length >= 5) || '密码不能小于六位。',
        v => (v && v.length <= 16) || '密码太长啦，你记得住嘛？'
      ],
      passwordRepeat: '',
      passwordRepeatRules: [v => v == this.password || '前后密码不一致。']
    };
  },
  methods: {
    goLogin() {
      this.$router.push('/login');
    },
    doRegister() {
      if (this.valid === false) {
        this.snackbar = true;
        return;
      }
      this.$axios({
        method: 'post',
        url: '/users/register',
        data: {
          username: this.username,
          password: this.password
        }
      })
        .then(res => {
          console.log(res);
          Toast.fire({
            icon: 'success',
            title: '注册成功，可以登录啦。'
          });
        })
        .catch(err => {
          Toast.fire({
            icon: 'success',
            title: `注册失败, 错误: ${err}`
          });
        });
      this.goLogin();
    }
  }
};
</script>
