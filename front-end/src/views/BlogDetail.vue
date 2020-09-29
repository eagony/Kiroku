<template>
  <v-container>
    <v-card>
      <v-card-title class="d-flex justify-center mb-5">
        <h4>{{ blog.title }}</h4>
      </v-card-title>
      <v-card-text>
        <vue-markdown
          :source="blog.content"
          :toc-first-level="1"
          :toc-last-level="3"
          :toc="true"
          toc-id="toc"
          class="markdown-body"
        >
        </vue-markdown>
      </v-card-text>
      <v-card-actions>
        <v-row justify="center" align="center">
          <v-btn
            fab
            dark
            small
            color="pink"
            @click="blogLikesPlus"
          >
            <v-icon dark>
              mdi-heart
            </v-icon>
          </v-btn>
        </v-row>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script>
import Toast from '../plugins/toast';
import VueMarkdown from 'vue-markdown';
import hljs from 'highlight.js';
const highlightCode = () => {
  let blocks = document.querySelectorAll('pre code');
  blocks.forEach(block => {
    hljs.highlightBlock(block);
  });
};

export default {
  name: 'BlogDetail',
  components: { VueMarkdown },
  data() {
    return {
      blog: {},
      subfield: false,
      defaultOpen: 'preview',
      editable: false,
      toolbarsFlag: false,
      scrollStyle: true
    };
  },
  methods: {
    getBlog() {
      this.$axios({
        method: 'get',
        url: `/blogs/${this.$route.params.id}`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          console.log(res.data);
          this.blog = res.data.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    blogViewsPlus() {
      this.$axios.get(`/statistic/blogs/${this.$route.params.id}/views`).catch(err => {
        console.log(err)
      })
    },
    blogLikesPlus() {
            this.$axios.get(`/statistic/blogs/${this.$route.params.id}/likes`).then(() => {
              Toast.fire({
                icon: 'success',
                title: '作者已收到你的赞。'
              })
            }).catch(err => {
        console.log(err)
      })
    }
  },
  created() {
    this.getBlog();
    this.blogViewsPlus();
    highlightCode();
  },
  updated() {
    highlightCode();
  }
};
</script>
