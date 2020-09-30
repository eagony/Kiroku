<template>
  <v-container>
    <v-row dense>
      <v-col cols="12">
        <!-- 提示栏 -->
        <v-card>
          <v-card-title>
            <div class="text-justify text--black">
              共 {{ blogs.length }} 篇公开博客
            </div>
            <v-spacer></v-spacer>
            <div>
              <v-btn class="mx-2" fab dark color="teal">
                <v-icon dark>mdi-refresh</v-icon>
              </v-btn>
            </div>
          </v-card-title>
        </v-card>
      </v-col>
    </v-row>

    <!-- 博客列表 -->
    <v-row>
      <v-col cols="12" v-for="blog in blogs" :key="blog.ID">
        <PublicBlogPreviewCard
          :id="blog.ID"
          :title="blog.title"
          :summary="blog.summary"
          :tags="blog.tags"
          :views="blog.views"
          :likes="blog.likes"
        ></PublicBlogPreviewCard>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import PublicBlogPreviewCard from '../components/PublicBlogPreviewCard';
export default {
  name: 'Square',
  components: {
    PublicBlogPreviewCard
  },
  data() {
    return {
      blogs: []
    };
  },
  methods: {
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
          console.log(err);
        });
    }
  },
  mounted() {
    this.getPublicBlogList();
  }
};
</script>
