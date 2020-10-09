<template>
  <v-card>
    <v-card-text>
      <v-textarea
        v-model="content"
        auto-grow
        counter
        name="input-7-4"
        label="发表评论"
      ></v-textarea>
    </v-card-text>
    <v-card-actions>
      <!-- 是否公开， 默认否 -->
      <v-switch class="ml-3" v-model="isPublic" inset label="公开"></v-switch>
      <v-spacer></v-spacer>
      <v-btn class="white--text" color="purple accent-4">
        取消
      </v-btn>
      <v-btn
        class="white--text mr-3"
        color="teal accent-4"
        @click="addCommentOnBlog"
      >
        完成
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import Toast from '../plugins/toast';

export default {
  name: 'CommentInputCard',
  props: ['blogID'],
  data() {
    return {
      content: '',
      isPublic: true
    };
  },
  methods: {
    addCommentOnBlog() {
      if (!this.$store.getters.isLoggedIn) {
        Toast.fire({
          icon: 'error',
          title: '请先登录！'
        });
      this.$router.push('/login')
      return
      }
      if (this.content.length < 1) {
        alert('总得写点什么吧。');
        return;
      }
      this.$axios({
        method: 'post',
        url: '/comments',
        data: {
          content: this.content,
          user_id: this.$store.state.user.id,
          blog_id: this.blogID,
          username: this.$store.state.user.username,
          user_avatar: this.$store.state.user.avatar,
          invisibility: this.isPublic ? 'public' : 'private'
        },
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(() => {
          Toast.fire({
            icon: 'success',
            title: '添加评论成功。'
          });
          this.content = '';
          this.isPublic = true;
          this.$parent.getComments();
        })
        .catch(err => {
          console.log('error on Blog.addBlog: ', err);
        });
    }
  }
};
</script>
