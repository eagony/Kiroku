<template>
  <v-container>
    <v-card>
      <v-card-title class="d-flex justify-center mb-5">
        <h4>{{ blog.title }}</h4>
      </v-card-title>
      <v-card-text>
        <vue-markdown
          :source="blog.content"
          :toc-first-level="1"
          :toc-last-level="3"
          :toc="true"
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
                  ><small> {{ blog.CreatedAt }}</small></v-list-item-subtitle
                >
              </v-list-item-content>
            </v-list-item>
          </v-list>
          <v-spacer></v-spacer>
          <v-list class="">
            <v-btn icon color="grey" class="ml-3" @click="blogLikesPlus">
              <v-icon>mdi-thumb-up</v-icon>
            </v-btn>
            <v-btn icon color="grey" class="ml-3">
              <v-icon>mdi-thumb-down</v-icon>
            </v-btn>
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
  </v-container>
</template>

<script>
import Toast from '../plugins/toast';
import VueMarkdown from 'vue-markdown';
import hljs from 'highlight.js';

import CommentInputCard from '../components/CommentInputCard';
import CommentCard from '../components/CommentCard';

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
      subfield: false,
      defaultOpen: 'preview',
      editable: false,
      toolbarsFlag: false,
      scrollStyle: true
    };
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
          console.log(err);
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
          console.log(err);
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
          console.log(err);
        });
    },
    blogViewsPlus() {
      this.$axios
        .get(`/statistic/blogs/${this.$route.params.id}/views`)
        .catch(err => {
          console.log(err);
        });
    },
    blogLikesPlus() {
      this.$axios
        .get(`/statistic/blogs/${this.$route.params.id}/likes`)
        .then(() => {
          Toast.fire({
            icon: 'success',
            title: '作者已收到你的赞。'
          });
        })
        .catch(err => {
          console.log(err);
        });
    }
  },
  created() {
    this.getBlog();
    this.blogViewsPlus();
    highlightCode();
  },
  updated() {
    highlightCode();
  }
};
</script>
