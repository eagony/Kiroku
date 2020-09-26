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
                      src="https://firebasestorage.googleapis.com/v0/b/todoteam-3263d.appspot.com/o/images%2FC6ridYN4MZfCUIw8L6bDG6wX0U32?alt=media&token=423cba6a-cfc5-4241-859e-11341cc87e62"
                      class="userAvatar"
                      @click="selectFile"
                    />
                  </v-avatar>
                </v-row>
                <input
                  type="file"
                  ref="uploadInput"
                  id="files"
                  accept="image/*"
                  :multiple="false"
                  @change="detectFiles($event)"
                />
                <v-text-field
                  prepend-icon="person"
                  label="用户名"
                  v-model="user.name"
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
                <v-text-field
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
                ></v-text-field>
              </v-form>
            </v-row>
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              class="mb-3 mr-3"
              color="teal"
              large
              @click="updateProfile()"
              :loading="loading"
              style="min-width:150px;"
              >Save</v-btn
            >
            <v-spacer></v-spacer>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Settings',
  data() {
    return {
      user: {},
      newPassword: '',
      newPassword2: '',
      loading: false
    };
  },

  methods: {
    updateProfile() {
      this.loading = true;

      if (this.newPassword.length > 0 || this.newPassword2.length > 0) {
        if (this.newPassword != this.newPassword2) {
          this.loading = false;
          this.snackbarMessage = 'Passwords do not match';
          this.snackbar = true;
        }
      }
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
    upload(file) {
      console.log(file);
    }
  },
  created() {}
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
