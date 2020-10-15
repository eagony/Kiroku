<template>
  <v-container>
    <v-row class="justify-center" dense>
      <v-col cols="12" md="10" xl="8">
        <!-- 提示栏 -->
        <v-card>
          <v-card-title>
            <div class="text-justify text--black ml-1">
              <span class="text-h4">
                {{ blogs.length }}
              </span>
              篇公开博客
            </div>
            <v-spacer></v-spacer>
            <div>
              <v-btn class="mx-2" fab dark color="teal" @click="refreshPage">
                <v-icon dark>mdi-refresh</v-icon>
              </v-btn>
            </div>
          </v-card-title>
        </v-card>
      </v-col>
    </v-row>

    <!-- 博客列表 -->
    <v-row class="justify-center">
      <v-col cols="12" md="10" xl="8" v-for="blog in blogs" :key="blog.ID">
        <PublicBlogPreviewCard
          :id="blog.ID"
          :title="blog.title"
          :summary="blog.summary"
          :tags="blog.tags"
          :views="blog.views"
          :likes="blog.likes"
          :comments="blog.comments.length"
        ></PublicBlogPreviewCard>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Swal from 'sweetalert2';
import PublicBlogPreviewCard from '../components/PublicBlogPreviewCard';

export default {
  name: 'Explore',
  components: {
    PublicBlogPreviewCard
  },
  data() {
    return {
      blogs: []
    };
  },
  methods: {
    refreshPage() {
      this.getPublicBlogList();
    },
    getPublicBlogList() {
      this.$axios({
        method: 'get',
        url: `/publicblogs`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          this.blogs = res.data.data;
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
  mounted() {
    this.getPublicBlogList();
  }
};
</script>
