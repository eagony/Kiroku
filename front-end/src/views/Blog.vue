<template>
  <v-container>
    <v-row class="justify-center" dense>
      <!-- 编辑栏 -->
      <v-col cols="12" md="9" xl="7">
        <!-- 提示栏 -->
        <v-card>
          <!-- 不输入时 -->
          <v-card-title v-if="!adding">
            <div class="text-justify text--black">
              共 {{ blogs.length }} 篇博客
            </div>
            <v-spacer></v-spacer>
            <div>
              <v-btn class="mx-2" fab dark color="teal" @click="adding = true">
                <v-icon dark>mdi-pencil</v-icon>
              </v-btn>
            </div>
          </v-card-title>
          <!-- 添加时 -->
          <v-card-text v-if="adding">
            <v-row dense>
              <!-- 标题输入 -->
              <v-col cols="12">
                <div class="d-flex justify-center">
                  <h2 class="">标题</h2>
                </div>
              </v-col>
              <v-col cols="12">
                <v-text-field solo v-model="title" label="博客标题...">
                </v-text-field>
              </v-col>
              <!-- 摘要 -->
              <v-col cols="12">
                <div class="d-flex justify-center">
                  <h2 class="">摘要</h2>
                </div>
              </v-col>
              <v-col cols="12">
                <v-textarea
                  class="ml-2"
                  v-model="summary"
                  auto-grow
                  counter
                  name="input-7-4"
                  label="博客摘要..."
                ></v-textarea>
              </v-col>
              <!-- 正文 -->
              <v-col cols="12">
                <div class="d-flex justify-center">
                  <h2 class="mb-1">正文</h2>
                </div>
              </v-col>
              <v-col cols="12">
                <div class="d-flex justify-center">
                  <mavon-editor
                    ref="md"
                    :boxShadow="false"
                    v-model="content"
                    @imgAdd="$imgAdd"
                    @imgDel="$imgDel"
                  ></mavon-editor>
                </div>
              </v-col>
            </v-row>
            <!-- 标签 -->
            <v-row class="ml-1 mt-5">
              <span class="text-body-2 mr-5 mt-3">添加标签</span>
              <v-chip-group
                column
                v-for="chip in chips"
                :key="chip.ID"
                multiple
                v-model="tags"
                active-class="teal--text text--accent-4"
              >
                <v-chip :value="chip"
                  ><v-icon left>{{ chip.icon || 'mdi-label' }}</v-icon
                  >{{ chip.text }}</v-chip
                >
              </v-chip-group>
            </v-row>
          </v-card-text>
          <!-- 底部按钮 -->
          <v-card-actions v-if="adding">
            <!-- 是否公开， 默认是 -->
            <v-switch
              class="ml-3"
              v-model="isPublic"
              inset
              label="公开"
            ></v-switch>
            <v-spacer></v-spacer>
            <v-btn
              class="white--text"
              color="purple accent-4"
              @click="adding = false"
            >
              取消
            </v-btn>
            <v-btn class="white--text" color="teal accent-4" @click="addBlog">
              完成
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>

      <!-- 博客列表 -->
      <v-col cols="12" md="9" xl="7" v-for="blog in blogs" :key="blog.ID">
        <v-card class="">
          <!-- 标题 -->
          <router-link
            :to="{ name: 'BlogDetail', params: { id: blog.ID } }"
            style="text-decoration: none; color: black;"
          >
            <v-card-title class="d-flex justify-center">
              <span class="title font-weight-regular">{{ blog.title }}</span>
            </v-card-title>
          </router-link>

          <!-- 摘要 -->
          <router-link
            :to="{ name: 'BlogDetail', params: { id: blog.ID } }"
            style="text-decoration: none;"
          >
            <v-card-text>
              <div class="ml-1 mr-1">
                <h5 class="text-h6 font-weight-regular" style="color: black;">
                  {{ blog.summary }}
                </h5>
              </div>
            </v-card-text>
          </router-link>

          <v-divider></v-divider>

          <!-- 标签 -->
          <v-card-text v-if="blog.tags.length > 0">
            <v-row justify="center" align="center">
              <v-chip
                v-for="(chip, index) in blog.tags"
                :key="index"
                class="ma-2"
                :color="chip.color"
                label
                text-color="white"
              >
                <v-icon left>{{ chip.icon || 'mdi-label' }}</v-icon>
                {{ chip.text }}
              </v-chip>
            </v-row>
          </v-card-text>
          <!-- 数据和操作 -->
          <v-card-actions>
            <v-row class="pt-3 pb-3" justify="center" align="center">
              <!-- 浏览量 -->
              <v-icon class="mr-2">mdi-eye-outline</v-icon>
              <span class="subheading">{{ blog.views }}</span>

              <v-divider vertical class="ml-5 mr-5"></v-divider>

              <!-- 点赞数 -->
              <v-icon class="mr-2">mdi-heart-outline</v-icon>
              <span class="subheading">{{ blog.likes }}</span>

              <v-divider vertical class="ml-5 mr-5"></v-divider>

              <!-- 评论数 -->
              <v-icon class="mr-2">mdi-comment-outline</v-icon>
              <span class="subheading">{{
                blog.comments ? blog.comments.length : 0
              }}</span>

              <v-divider vertical class="ml-5 mr-5"></v-divider>

              <!-- 编辑 -->
              <v-icon @click="editBlog(blog)">mdi-circle-edit-outline</v-icon>

              <v-divider vertical class="ml-5 mr-5"></v-divider>

              <!-- 删除 -->
              <v-icon @click="deleteBlog(blog.ID)"
                >mdi-trash-can-outline</v-icon
              >
            </v-row>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Swal from 'sweetalert2';
import Toast from '../plugins/toast';

export default {
  name: 'Blog',
  data() {
    return {
      tags: [],
      chips: [],
      blogs: [],
      title: '',
      summary: '',
      content: '',
      adding: false,
      isPublic: true
    };
  },
  methods: {
    // 获取博客可选的标签
    getTags() {
      this.$axios({
        method: 'get',
        url: `/tags?use_for=blog`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          this.chips = res.data.data;
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    },
    // 获取用户博客列表
    getBlogList() {
      this.$axios({
        method: 'get',
        url: `/users/${this.$store.state.user.id}/blogs`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          this.blogs = res.data.data;
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    },
    // 添加博客
    addBlog() {
      if (this.content.length < 1) {
        alert('必须写点什么吧。');
        return;
      }
      this.$axios({
        method: 'post',
        url: '/blogs',
        data: {
          title: this.title,
          summary: this.summary,
          content: this.content,
          tags: this.tags,
          invisibility: this.isPublic ? 'public' : 'private',
          user_id: this.$store.state.user.id
        },
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(() => {
          Toast.fire({
            icon: 'success',
            title: '成功新建一篇博客。'
          });
          this.adding = false;
          this.title = '';
          this.summary = '';
          this.content = '';
          this.tags = [];
          this.getBlogList();
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    },
    // 编辑博客
    editBlog(blog) {
      console.log(blog);
    },
    // 删除博客
    deleteBlog(id) {
      Swal.fire({
        title: '确定要删除这篇博客吗？',
        text: '此操作不可逆，删除的博客无法找回哦!',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: '删除',
        cancelButtonText: '取消'
      }).then(result => {
        // 确认删除
        if (result.isConfirmed) {
          this.$axios({
            method: 'delete',
            url: `/blogs/${id}`,
            headers: {
              Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
            }
          })
            .then(() => {
              Swal.fire('删除成功!', '你刚刚删除了一篇博客。', 'success');
              this.getBlogList();
            })
            .catch(err => {
              Swal.fire({
                icon: 'error',
                title: '出错了',
                text: `${err.response.data.error}`
              });
            });
        } else if (
          // 取消删除
          result.dismiss === Swal.DismissReason.cancel
        ) {
          Swal.fire('已取消', '你的博客还健在 :)', 'error');
        }
      });
    },
    // TODO: 图片上传有问题，点击不弹出
    // 绑定@imgAdd event
    $imgAdd(pos, $file) {
      console.log('clicked');
      // 第一步.将图片上传到服务器.
      var formdata = new FormData();
      formdata.append('image', $file);
      this.$axios({
        url: '/images',
        method: 'post',
        data: formdata,
        headers: {
          'Content-Type': 'multipart/form-data',
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      }).then(res => {
        // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
        // $vm.$img2Url 详情见本页末尾
        this.$refs.md.$img2Url(
          pos,
          this.$axios.defaults.baseURL + '/' + res.data.url
        );
      });
    },
    $imgDel() {}
  },
  mounted() {
    this.getBlogList();
    this.getTags();
  }
};
</script>

<style scoped>
a {
  text-decoration: none;
}
.router-link-active {
  text-decoration: none;
}
</style>
