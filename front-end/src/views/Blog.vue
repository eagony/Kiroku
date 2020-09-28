<template>
  <v-container>
    <v-row dense>
      <v-col cols="12">
        <v-card>
          <v-card-title v-if="!adding">
            <div class="text-justify text--black">共 10 篇博客</div>
            <v-spacer></v-spacer>
            <div>
              <v-btn class="mx-2" fab dark color="teal" @click="adding = true">
                <v-icon dark>mdi-pencil</v-icon>
              </v-btn>
            </div>
          </v-card-title>

          <v-card-text v-if="adding">
            <v-row dense>
              <v-col cols="12">
                <div class="d-flex justify-center">
                  <h2 class="">标题</h2>
                </div>
              </v-col>
              <v-col cols="12">
                <v-text-field solo v-model="title" label="博客标题...">
                </v-text-field>
              </v-col>
              <v-col cols="12">
                <div class="d-flex justify-center">
                  <h2 class="">摘要</h2>
                </div>
              </v-col>
              <v-col cols="12">
                <v-textarea
                  class="ml-2"
                  v-model="summary"
                  auto-grow
                  counter
                  name="input-7-4"
                  label="博客摘要..."
                ></v-textarea>
              </v-col>
              <v-col cols="12">
                <div class="d-flex justify-center">
                  <h2 class="mb-1">正文</h2>
                </div>
              </v-col>
              <v-col cols="12">
                <mavon-editor v-model="content"></mavon-editor>
              </v-col>
            </v-row>
            <v-row class="ml-1 mt-5">
              <span class="text-body-2 mr-5 mt-3">添加标签</span>

              <v-chip-group
                column
                v-for="chip in chips"
                :key="chip.ID"
                multiple
                v-model="tags"
                active-class="teal--text text--accent-4"
              >
                <v-chip :value="chip"
                  ><v-icon left>{{ chip.icon || 'mdi-label' }}</v-icon
                  >{{ chip.text }}</v-chip
                >
              </v-chip-group>
            </v-row>
          </v-card-text>
          <v-card-actions v-if="adding">
            <v-spacer></v-spacer>
            <v-btn
              class="white--text"
              color="purple accent-4"
              @click="adding = false"
            >
              取消
            </v-btn>
            <v-btn class="white--text" color="teal accent-4" @click="addBlog">
              完成
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" md="4" v-for="blog in blogs" :key="blog.ID">
        <v-card class="" min-height="400">
          <v-card-title class="d-flex justify-center">
            <span class="title font-weight-regular">{{ blog.title }}</span>
          </v-card-title>

          <v-card-text>
            <router-link
              :to="{ name: 'BlogDetail', params: { id: blog.ID } }"
              style="text-decoration: none;"
            >
              <div class="ml-1 mr-1" style="min-height: 150px;">
                <h5 class="text-h6 font-weight-regular" style="color: black;">
                  {{ blog.summary }}
                </h5>
              </div>
            </router-link>
          </v-card-text>
          <v-divider></v-divider>

          <v-card-actions>
            <v-row>
              <v-row justify="center">
                <v-chip
                  v-for="(chip, index) in chips"
                  :key="index"
                  class="ma-2"
                  :color="chip.color"
                  label
                  text-color="white"
                >
                  <v-icon left>{{ chip.icon || 'mdi-label' }}</v-icon>
                  {{ chip.text }}
                </v-chip>
              </v-row>

              <v-row class="mt-3 pt-3 pb-3" justify="center">
                <v-icon class="mr-2">mdi-eye-outline</v-icon>
                <span class="subheading">256</span>
                <v-divider vertical class="ml-5 mr-5"></v-divider>
                <v-icon class="mr-2">mdi-heart-outline</v-icon>
                <span class="subheading">45</span>
                <v-divider vertical class="ml-5 mr-5"></v-divider>
                <v-icon class="mr-2">mdi-comment-outline</v-icon>
                <span class="subheading">45</span>
                <v-divider vertical class="ml-5 mr-5"></v-divider>
                <v-icon>mdi-trash-can-outline</v-icon>
              </v-row>
            </v-row>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Toast from '../plugins/toast';

export default {
  name: 'Blog',
  data() {
    return {
      blogs: [],
      adding: false,
      title: '',
      summary: '',
      content: '',
      tags: [],
      // temp
      chips: []
    };
  },
  methods: {
    // temp
    getTags() {
      this.$axios({
        method: 'get',
        url: `/tags?use_for=blog`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          console.log(res.data);
          this.chips = res.data.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    getBlogList() {
      this.$axios({
        method: 'get',
        url: `/users/${this.$store.state.user.id}/blogs`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          this.blogs = res.data.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    addBlog() {
      this.$axios({
        method: 'post',
        url: '/blogs',
        data: {
          title: this.title,
          summary: this.summary,
          content: this.content,
          user_id: this.$store.state.user.id
        },
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(() => {
          console.log('新建博客成功');
          Toast.fire({
            icon: 'success',
            title: '成功！',
            text: '成功新建一篇博客。'
          });
        })
        .catch(err => {
          console.log(err);
        });
      this.adding = false;
      this.title = '';
      this.summary = '';
      this.content = '';
      this.getBlogList();
    }
  },
  mounted() {
    this.getBlogList();
    this.getTags();
  }
};
</script>

<style scoped>
a {
  text-decoraction: none;
}
.router-link-active {
  text-decoration: none;
}
</style>
