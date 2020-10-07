<template>
  <v-container>
    <v-row dense>
      <v-col cols="12">
        <v-card>
          <v-card-text>
            <v-row justify="center">
              <v-form style="min-width: 450px;">
                <v-row justify="center" class="mb-8">
                  <v-avatar size="200" class="avatar">
                    <img
                      :src="user.avatar"
                      class="userAvatar"
                      @click="selectFile"
                    />
                  </v-avatar>
                </v-row>
                <input
                  type="file"
                  ref="uploadInput"
                  name="image"
                  id="files"
                  accept="image/*"
                  :multiple="false"
                  @change="detectFiles($event)"
                />
                <v-text-field
                  prepend-icon="person"
                  label="用户名"
                  v-model="user.username"
                  color="primary"
                ></v-text-field>
                <v-textarea
                  class="ml-7"
                  v-model="user.signature"
                  auto-grow
                  counter
                  label="个性签名"
                ></v-textarea>
                <v-text-field
                  prepend-icon="email"
                  label="邮箱"
                  v-model="user.email"
                  color="primary"
                ></v-text-field>
                <v-text-field
                  prepend-icon="phone"
                  label="手机"
                  v-model="user.phone"
                  color="primary"
                ></v-text-field>
                <!-- <v-text-field
                  prepend-icon="lock"
                  type="password"
                  label="新密码"
                  v-model="newPassword"
                  color="primary"
                ></v-text-field>
                <v-text-field
                  prepend-icon="lock"
                  type="password"
                  label="确认新密码"
                  v-model="newPassword2"
                  color="primary"
                ></v-text-field> -->
              </v-form>
            </v-row>
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              class="mb-3 mr-3"
              color="teal"
              large
              @click="updateProfile('资料修改成功。')"
              :loading="loading"
              style="min-width:150px;"
              >更新</v-btn
            >
            <v-spacer></v-spacer>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Toast from '../plugins/toast';

export default {
  name: 'Settings',
  data() {
    return {
      user: {},
      // newPassword: '',
      // newPassword2: '',
      loading: false
    };
  },

  methods: {
    // 获取用户个人信息
    getProfile() {
      this.$axios({
        method: 'get',
        url: `/users/${this.$store.state.user.id}`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          this.user = res.data.data;
        })
        .catch(err => {
          console.log('error on Settings.getProfile: ', err);
        });
    },
    // 更新用户信息
    updateProfile(msg) {
      this.loading = true;
      this.$axios({
        method: 'put',
        url: `/users/${this.$store.state.user.id}`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        },
        data: this.user
      })
        .then(() => {
          Toast.fire({
            icon: 'success',
            title: msg
          });
          this.getProfile();
          this.loading = false;
        })
        .catch(err => {
          console.log(err);
        });
    },

    selectFile() {
      this.$refs.uploadInput.click();
    },

    detectFiles(e) {
      this.loading = true;
      let fileList = e.target.files || e.dataTransfer.files;
      Array.from(Array(fileList.length).keys()).map(x => {
        this.upload(fileList[x]);
      });
    },
    // 上传头像，上传完后立即更新一次用户信息
    upload(file) {
      const formData = new FormData();
      formData.append('image', file);
      this.$axios({
        method: 'post',
        url: '/images',
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data',
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          this.user.avatar = this.$axios.defaults.baseURL + '/' + res.data.uri;
          this.updateProfile('头像更新成功，重新登陆后刷新。');
        })
        .catch(err => {
          console.log('error on Settings.uoloadAvatar: ', err);
        });
      this.loading = false;
    }
  },
  created() {
    this.getProfile();
  }
};
</script>

<style>
input[type='file'] {
  position: absolute;
  clip: rect(0, 0, 0, 0);
}

.avatar {
  border: 1px solid white;
  background-color: black;
}

.userAvatar {
  transition: opacity 0.2s ease-in-out;
}

.userAvatar:hover {
  opacity: 0.5;
  filter: alpha(opacity=50);
  cursor: pointer;
}
</style>
