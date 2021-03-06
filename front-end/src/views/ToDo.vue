<template>
  <v-container>
    <v-row class="justify-center">
      <v-col cols="12" md="9" xl="7">
        <v-card>
          <v-card-title>
            <v-row class="my-1" align="center">
              <strong class="mx-4 success--text text--darken-2">
                全部任务: {{ tasks.length }}
              </strong>
              <v-divider vertical></v-divider>
              <strong class="mx-4 info--text text--darken-2">
                剩余: {{ remainingTasks }}
              </strong>
              <v-spacer></v-spacer>
              <v-progress-circular
                size="40"
                width="5"
                :value="progress"
                color="pink"
                class="mr-4"
              ></v-progress-circular>
            </v-row>
          </v-card-title>
        </v-card>
        <v-text-field
          solo
          v-model="task"
          label="点击新建，回车完成..."
          @keypress.enter="addTask"
        >
          <template v-slot:append>
            <v-icon v-if="task" @click="addTask">
              mdi-plus
            </v-icon>
          </template>
        </v-text-field>

        <v-card v-for="(task, i) in tasks" :key="`${i}-divider`">
          <v-card-title>
            <v-spacer></v-spacer>
            <div :class="(task.done && 'grey--text') || 'success--text'">
              {{ task.text }}
              <v-spacer></v-spacer>
            </div>
            <v-spacer></v-spacer>
            <div class="d-flex align-center">
              <v-checkbox
                v-model="task.done"
                label=""
                :class="(task.done && 'success--text') || 'primary--text'"
                @click="updateTask(task)"
              ></v-checkbox>
              <v-icon @click="deleteTask(task.ID)">mdi-delete</v-icon>
            </div>
          </v-card-title>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Swal from 'sweetalert2';
import Toast from '../plugins/toast';

export default {
  data: () => ({
    tasks: [],
    task: null
  }),

  computed: {
    completedTasks() {
      return this.tasks.filter(task => task.done).length;
    },
    progress() {
      return (this.completedTasks / this.tasks.length) * 100;
    },
    remainingTasks() {
      return this.tasks.length - this.completedTasks;
    }
  },

  methods: {
    getToDoList() {
      this.$axios({
        method: 'get',
        url: `/users/${this.$store.state.user.id}/todos`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          this.tasks = res.data.data;
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    },
    addTask() {
      this.$axios({
        method: 'post',
        url: '/todos',
        data: {
          done: false,
          text: this.task,
          user_id: this.$store.state.user.id
        },
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(() => {
          Toast.fire({
            icon: 'success',
            title: '新建成功，记得完成哦。'
          });
          this.task = null;
          this.getToDoList();
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
          this.task = null;
        });
    },
    deleteTask(id) {
      this.$axios({
        method: 'delete',
        url: `/todos/${id}`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(() => {
          Toast.fire({
            icon: 'success',
            title: `删除成功。`
          });
          this.getToDoList();
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    },
    updateTask(task) {
      const path = `/todos/${task.ID}`;
      this.$axios({
        method: 'put',
        url: path,
        data: {
          text: task.text,
          done: task.done,
          user_id: this.$store.state.user.id
        },
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(() => {
          Toast.fire({
            icon: 'success',
            title: task.done
              ? '真胖，完成了一个任务！'
              : '真胖，又多出来了一个任务！'
          });
        })
        .catch(err => {
          Swal.fire({
            icon: 'error',
            title: '出错了',
            text: `${err.response.data.error}`
          });
        });
    }
  },
  mounted() {
    this.getToDoList();
  }
};
</script>
