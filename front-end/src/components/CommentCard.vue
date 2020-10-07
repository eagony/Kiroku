<template>
  <v-card class="">
    <!-- 标题 -->
    <v-card-title class="d-flex justify-start">
      <v-list>
        <v-list-item two-line>
          <v-list-item-avatar>
            <img
              :src="
                this.$store.state.user.avatar ||
                  'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
              "
            />
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title class="text-body-1">重邮谢霆锋</v-list-item-title>
            <v-list-item-subtitle><small> 9 分钟之前</small></v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-card-title>
    <!-- 摘要 -->
    <v-card-text>
      <div class="text-justify text-body-1 ml-16 pl-2">
        写得太好了
      </div>
    </v-card-text>
    <v-card-actions>
      <div class="ml-16 pl-2">
            <v-btn icon color="deep-purple">
              <v-icon>mdi-thumb-up</v-icon>
            </v-btn>
            <v-btn icon color="deep-purple" class="ml-2">
              <v-icon>mdi-thumb-down</v-icon>
            </v-btn>
      </div>
      <v-spacer></v-spacer>
      <div class="mr-5">
        <v-btn large text color="purple">回复</v-btn>
      </div>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'CommentCard'
};
</script>
