<template>
  <v-container>
    <v-row dense>
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

                <v-icon class="mr-1">mdi-heart-outline</v-icon>
                <span class="subheading mr-2">45</span>
                <v-icon class="mr-1">mdi-comment-outline</v-icon>
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
  name: 'Square',
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
