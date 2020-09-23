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
                <v-text-field solo v-model="title"> </v-text-field>
              </v-col>
              <v-col cols="12">
                <div class="d-flex justify-center">
                  <h2 class="mb-1">正文</h2>
                </div>
              </v-col>
              <v-col cols="12">
                <mavon-editor
                  v-model="content"
                  @change="handleChange"
                ></mavon-editor>
              </v-col>
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
      <v-col cols="12" v-for="blog in blogs" :key="blog.ID">
        <v-card class="">
          <v-card-title class="d-flex justify-center">
            <span class="title font-weight-regular">{{ blog.title }}</span>
          </v-card-title>

          <v-card-text class="headline font-weight-normal">
            一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要一些摘要
          </v-card-text>

          <v-card-actions>
            <v-list-item class="grow">
              <v-list-item-avatar color="grey darken-3">
                <v-img
                  class="elevation-6"
                  src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
                ></v-img>
              </v-list-item-avatar>

              <v-list-item-content>
                <v-list-item-title>Evan You</v-list-item-title>
              </v-list-item-content>

              <v-row align="center" justify="end">
                <v-icon class="mr-1">mdi-eye-outline</v-icon>
                <span class="subheading mr-2">256</span>
                <span class="mr-1">·</span>
                <v-icon class="mr-1">mdi-heart-outline</v-icon>
                <span class="subheading">45</span>
              </v-row>
            </v-list-item>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Blog',
  data() {
    return {
      blogs: [],
      adding: false,
      content: '',
      title: '',
      mark: '',
      html: ''
    };
  },
  methods: {
    handleChange(markdown, html) {
      this.mark = markdown;
      this.html = html;
    },
    getBlogList() {
      alert( `/users/${this.$store.state.user.id}/blogs`)
      this.$axios({
        method: 'get',
        url: `/users/${this.$store.state.user.id}/blogs`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      }).then(res => {
        this.blogs = res.data.data;
      }).catch(err => {
        console.log(err)
      })
    },
    addBlog() {
      this.$axios({
        method: 'post',
        url: '/blogs',
        data: {
          title: this.title,
          content: this.content,
          user_id: this.$store.state.user.id
        },
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(() => {
          console.log('新建博客成功');
        })
        .catch(err => {
          console.log(err);
        });
    }
  },
  mounted() {
    this.getBlogList()
  }
};
</script>
