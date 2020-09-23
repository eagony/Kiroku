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
    </v-card>
  </v-container>
</template>

<script>
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
    }
  },
  mounted() {
    this.getBlog();
    highlightCode();
  },
  updated() {
    highlightCode();
  }
};
</script>
