<template>
  <v-card class="">
    <!-- 标题 -->
    <v-card-title class="d-flex justify-center">
      <span class="title font-weight-regular">{{ title }}</span>
    </v-card-title>
    <!-- 摘要 -->
    <v-card-text>
      <router-link
        :to="{ name: 'BlogDetail', params: { id: id } }"
        style="text-decoration: none;"
      >
        <div class="ml-1 mr-1">
          <h5 class="text-h6 font-weight-regular" style="color: black;">
            {{ summary }}
          </h5>
        </div>
      </router-link>
    </v-card-text>

    <v-divider></v-divider>

    <!-- 标签 -->
    <v-card-text>
      <v-row justify="center" align="center">
        <v-col v-if="tags.length > 0" cols="12" md="9" xl="10">
          <v-chip
            v-for="chip in tags"
            :key="chip.ID"
            class="ma-2"
            :color="chip.color"
            label
            text-color="white"
          >
            <v-icon left>{{ chip.icon || 'mdi-label' }}</v-icon>
            {{ chip.text }}
          </v-chip>
        </v-col>
        <v-col cols="12" md="3" xl="2" class="d-flex justify-end">
          <!-- 浏览量 -->
          <v-icon class="mr-2">mdi-eye-outline</v-icon>
          <span class="subheading">{{ views }}</span>
          <v-divider vertical class="ml-5 mr-5"></v-divider>
          <!-- 点赞数 -->
          <v-icon class="mr-2">mdi-heart-outline</v-icon>
          <span class="subheading">{{ likes }}</span>
          <v-divider vertical class="ml-5 mr-5"></v-divider>
          <!-- 评论数 -->
          <v-icon class="mr-2">mdi-comment-outline</v-icon>
          <span class="subheading">{{ comments }}</span>
          <div class="mr-3"></div>
        </v-col>
      </v-row>
    </v-card-text>
    <!-- 数据和操作 -->
  </v-card>
</template>

<script>
export default {
  name: 'PublicBlogPreviewCard',
  props: ['id', 'title', 'summary', 'tags', 'views', 'likes', 'comments']
};
</script>
