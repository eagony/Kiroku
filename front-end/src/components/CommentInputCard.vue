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
      <v-btn class="white--text" color="teal accent-4">
        完成
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'CommentInputCard',
  data() {
    return {
      content: '',
      isPublic: true
    };
  }
};
</script>
