<template>
  <v-card class="">
    <!-- 标题 -->

    <router-link
      :to="{ name: 'BlogDetail', params: { id: id } }"
      style="text-decoration: none; color: black;"
    >
      <v-card-title class="d-flex justify-center">
        <span class="title font-weight-regular">{{ title }}</span>
      </v-card-title>
    </router-link>

    <!-- 摘要 -->
    <router-link
      :to="{ name: 'BlogDetail', params: { id: id } }"
      style="text-decoration: none; color: black;"
    >
      <v-card-text>
        <div class="ml-1 mr-1">
          <h5 class="text-h6 font-weight-regular" style="color: black;">
            <span class="text-h5" style="color: grey;">
              摘要:
            </span>
            {{ summary }}
          </h5>
        </div>
      </v-card-text>
    </router-link>

    <v-divider></v-divider>

    <!-- 标签 -->
    <v-card-text>
      <v-row align="center">
        <v-col v-if="tags.length > 0" cols="12" md="8">
          <v-row justify="center" justify-md="start" class="mx-0">
            <v-chip
              v-for="chip in tags"
              :key="chip.ID"
              class="mx-2 my-1"
              :color="chip.color"
              label
              text-color="white"
            >
              <v-icon left>{{ chip.icon || 'mdi-label' }}</v-icon>
              {{ chip.text }}
            </v-chip>
          </v-row>
        </v-col>
        <v-col cols="12" md="4">
          <v-row justify="center" justify-md="end" class="mx-0">
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
          </v-row>
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
