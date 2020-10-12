<template>
  <v-container>
    <v-row class="justify-center">
      <v-col cols="12" md="9" xl="8">
        <v-card>
          <v-card-title class="d-flex justify-center mb-5">
            <h4>{{ blog.title }}</h4>
          </v-card-title>
          <v-card-text class="pl-8 pr-5">
            <!-- vue-markdown 开始解析markdown，它是子组件，通过 props 给它传值
            要指定TOC的级数哦，如果要修改TOC的样式，要在toc-rendered指定的函数中操作，因为要等它把TOC给创建出来
             -->
            <vue-markdown
              :source="blog.content"
              :toc="showToc"
              :toc-first-level="1"
              :toc-last-level="3"
              toc-id="toc"
              class="markdown-body"
            >
            </vue-markdown>
          </v-card-text>
          <v-card-actions>
            <v-row align="center" class="pl-5 pr-5 mt-5">
              <v-list>
                <v-list-item two-line>
                  <v-list-item-avatar>
                    <img
                      :src="
                        author.avatar ||
                          'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
                      "
                    />
                  </v-list-item-avatar>

                  <v-list-item-content>
                    <v-list-item-title class="text-body-1">{{
                      author.username
                    }}</v-list-item-title>
                    <v-list-item-subtitle
                      ><small>
                        {{ this.blogCreatedTime }}</small
                      ></v-list-item-subtitle
                    >
                  </v-list-item-content>
                </v-list-item>
              </v-list>
              <v-spacer></v-spacer>
              <v-list class="mr-5">
                <v-btn
                  icon
                  large
                  :color="thumbUpColor"
                  class="ml-3"
                  @click="blogLikesPlus"
                >
                  <v-icon>mdi-thumb-up</v-icon>
                </v-btn>
                <!-- <v-btn icon color="grey" class="ml-3">
                  <v-icon>mdi-thumb-down</v-icon>
                </v-btn> -->
                <!-- <v-btn text large color="teal" class="ml-2" >评论
            </v-btn> -->
              </v-list>
            </v-row>
          </v-card-actions>
        </v-card>

        <!-- 发表评论 -->
        <div class="mt-5 mb-5 mx-auto">
          <CommentInputCard :blogID="blog.ID"></CommentInputCard>
        </div>

        <!-- 评论列表 -->
        <div class="">
          <CommentCard
            class="mb-3"
            v-for="comment in comments"
            :key="comment.ID"
            :username="comment.username"
            :userAvatar="comment.user_avatar"
            :createdAt="comment.CreatedAt"
            :content="comment.content"
          ></CommentCard>
        </div>
      </v-col>
      <v-col cols="12" md="3" xl="2">
        <div id="sidebar">
          <v-card class="pb-5">
            <v-card-text>
              <div class="d-flex text-justfify justify-center mt-2">
                <h2>目录</h2>
              </div>
              <div id="toc" class="toc mt-5"></div>
            </v-card-text>
          </v-card>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Swal from 'sweetalert2';
import hljs from 'highlight.js';
import Toast from '../plugins/toast';
import moment from '../plugins/moment';

import VueMarkdown from 'vue-markdown';
import CommentCard from '../components/CommentCard';
import CommentInputCard from '../components/CommentInputCard';

const highlightCode = () => {
  let blocks = document.querySelectorAll('pre code');
  blocks.forEach(block => {
    hljs.highlightBlock(block);
  });
};

export default {
  name: 'BlogDetail',
  components: {
    VueMarkdown,
    CommentInputCard,
    CommentCard
  },
  data() {
    return {
      blog: {},
      author: {},
      comments: [],
      showToc: true,
      subfield: false,
      thumbUpColor: 'grey',
      defaultOpen: 'preview',
      editable: false,
      toolbarsFlag: false,
      scrollStyle: true
    };
  },
  computed: {
    blogCreatedTime: function() {
      return moment(this.blog.CreatedAt).format('llll');
    }
  },
  methods: {
    getBlog() {
      this.$axios({
        method: 'get',
        url: `/blogs/${this.$route.params.id}`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          this.blog = res.data.data;
          this.getAuthor();
          this.getComments();
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    },
    getAuthor() {
      this.$axios({
        method: 'get',
        url: `/users/${this.blog.user_id}`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          this.author = res.data.data;
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    },
    getComments() {
      this.$axios({
        method: 'get',
        url: `/blogs/${this.blog.ID}/comments`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          this.comments = res.data.data;
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    },
    blogViewsPlus() {
      this.$axios
        .get(`/statistic/blogs/${this.$route.params.id}/views`)
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    },
    blogLikesPlus() {
      if (this.thumbUpColor == 'pink') {
        Swal.fire({
          icon: 'error',
          title: '',
          text: '请勿重复点赞。'
        });
        return;
      }
      this.$axios
        .get(`/statistic/blogs/${this.$route.params.id}/likes`)
        .then(() => {
          this.thumbUpColor = 'pink';
          Toast.fire({
            icon: 'success',
            title: '作者已收到你的赞。'
          });
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    }
  },
  created() {
    this.getBlog();
    this.blogViewsPlus();
  },
  mounted() {
    highlightCode();
  },
  updated() {
    highlightCode();
  }
};
</script>

<style scoped>
#sidebar {
  position: sticky;
  top: 82px;
}
</style>
