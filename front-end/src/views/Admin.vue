<template>
  <v-container>
    <v-row dense>
      <v-col cols="12" sm="4" md="3">
        <v-card>
          <v-card-title>
            <v-row justify="center" align="center">
              <v-avatar class="mt-8" size="200" tile>
                <img
                  tile
                  :src="
                    this.$store.state.user.avatar ||
                      'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
                  "
                />
              </v-avatar>
            </v-row>
          </v-card-title>
          <v-card-title>
            <v-row justify="center" align="center">
              {{ this.$store.state.user.username }}
            </v-row>
          </v-card-title>
          <v-card-text>
            <v-list-item
              class="justify-center"
              v-for="item in items"
              :key="item.title"
              link
              style="margin-top: 5px;"
              :to="item.route"
            >
              <v-list-item-icon>
                <v-icon>{{ item.icon }}</v-icon>
              </v-list-item-icon>

              <v-list-item-content>
                <v-list-item-title>{{ item.title }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Admin',
  data() {
    return {
      items: [
        { title: 'User', icon: 'mdi-account', route: '/todo' },
        { title: 'ToDo', icon: 'mdi-format-list-bulleted', route: '/diary' },
        { title: 'Diary', icon: 'mdi-book-open-blank-variant', route: '/blog' },
        { title: 'Blog', icon: 'mdi-post-outline', route: '/settings' }
      ]
    };
  }
};
</script>

<style scoped></style>
