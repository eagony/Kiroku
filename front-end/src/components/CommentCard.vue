<template>
  <v-card class="">
    <!-- 标题 -->
    <v-card-title class="d-flex justify-start">
      <v-list>
        <v-list-item two-line>
          <v-list-item-avatar>
            <img
              :src="
                userAvatar ||
                  'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
              "
            />
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title class="text-body-1">{{
              username
            }}</v-list-item-title>
            <v-list-item-subtitle
              ><small> {{ createdAt }}</small></v-list-item-subtitle
            >
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-card-title>
    <!-- 摘要 -->
    <v-card-text>
      <div class="text-justify text-body-1 ml-16 pl-2">
        {{ content }}
      </div>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-list class="">
        <v-btn icon color="grey">
          <v-icon>mdi-thumb-up</v-icon>
        </v-btn>
        <v-btn icon color="grey" class="ml-3">
          <v-icon>mdi-thumb-down</v-icon>
        </v-btn>
        <!-- <v-btn text large color="teal" class="ml-2" >评论
            </v-btn> -->
      </v-list>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'CommentCard',
  props: ['username', 'userAvatar', 'createdAt', 'content']
};
</script>
