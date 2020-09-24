<template>
  <v-container>
    <v-row dense>
      <v-col cols="12">
        <v-card>
          <v-card-title v-if="!adding">
            <div class="text-justify text--black">
              共 {{ diaries.length }} 篇日记
            </div>
            <v-spacer></v-spacer>
            <div>
              <v-btn class="mx-2" fab dark color="teal" @click="adding = true">
                <v-icon dark>mdi-pencil</v-icon>
              </v-btn>
            </div>
          </v-card-title>

          <v-card-text v-if="adding">
            <v-textarea
              v-model="content"
              auto-grow
              counter
              name="input-7-4"
              label="新增日记"
            ></v-textarea>
            <v-row class="ml-1">
              <v-chip-group
                column
                v-for="(chip, index) in chips"
                :key="chip.ID"
                multiple
                v-model="selectedChips"
                active-class="teal--text text--accent-4"
              >
                <v-chip :value="index"
                  ><v-icon left>{{ chip.icon || 'mdi-label' }}</v-icon
                  >{{ chip.text }}</v-chip
                >
              </v-chip-group>
            </v-row>
          </v-card-text>
          <v-card-actions v-if="adding">
            <v-spacer></v-spacer>
            <v-btn
              class="white--text"
              color="purple accent-4"
              @click="adding = false"
            >
              取消
            </v-btn>
            <v-btn class="white--text" color="teal accent-4" @click="addDiary">
              完成
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col v-for="diary in diaries" :key="diary.ID" cols="12">
        <v-card class="mx-auto">
          <v-card-title>
            <p class="text-justify">
              {{ diary.CreatedAt.split('T')[0] }}({{
                getDay(diary.CreatedAt.split('T')[0])
              }})
            </p>
            <v-spacer></v-spacer>
            <v-icon @click="deleteDiary(diary.ID)"
              >mdi-trash-can-outline</v-icon
            >
          </v-card-title>

          <v-card-text>
            <div class="text-justify text-body-1 text--black">
              {{ diary.content }}
            </div>
          </v-card-text>

          <v-divider class="mt-6 mx-4"></v-divider>

          <v-card-text class="d-flex align-center">
            <v-chip
              v-for="(tag, index) in diary.tags.split(';')"
              :key="index"
              class="ma-2"
              :color="tag.split(',')[2]"
              label
              text-color="white"
            >
              <v-icon left>{{ tag.split(',')[0] || 'mdi-label' }}</v-icon>
              {{ tag.split(',')[1] }}
            </v-chip>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import moment from '../plugins/moment';
import Toast from '../plugins/toast';

export default {
  name: 'Diary',
  data: () => ({
    date: moment().format('YYYY.MM.DD'),
    day: moment().format('dddd'),
    adding: false,
    tags: '',
    content: '',
    chips: [],
    diaries: [],
    selectedChips: []
  }),
  methods: {
    getDay(date) {
      return moment(date).format('dddd');
    },
    getTags() {
      this.$axios({
        method: 'get',
        url: `/tags`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          console.log(res.data);
          this.chips = res.data.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    getDiaryList() {
      this.$axios({
        method: 'get',
        url: `/users/${this.$store.state.user.id}/diaries`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          console.log(res.data);
          this.diaries = res.data.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    deleteDiary(id) {
      this.$axios({
        method: 'delete',
        url: `/diaries/${id}`,
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(res => {
          if (res.status == 200) {
            Toast.fire({
              icon: 'success',
              title: '删除成功!'
            });
            this.getDiaryList();
          }
        })
        .catch(err => {
          Toast.fire({
            icon: 'error',
            title: `删除失败，错误${err}`
          });
        });
    },
    addDiary() {
      for (var i = 0, len = this.selectedChips.length; i < len; i++) {
        this.tags +=
          this.chips[this.selectedChips[i]].icon +
          ',' +
          this.chips[this.selectedChips[i]].text +
          ',' +
          this.chips[this.selectedChips[i]].color +
          ';';
      }
      this.$axios({
        method: 'post',
        url: '/diaries',
        data: {
          tags: this.tags.slice(0, -1),
          content: this.content,
          user_id: this.$store.state.user.id
        },
        headers: {
          Authorization: 'Bearer ' + window.localStorage.getItem('r-token')
        }
      })
        .then(() => {
          console.log('新建日记成功');
        })
        .catch(err => {
          console.log(err);
        });
      (this.tags = ''), (this.adding = false);
      this.content = null;
      this.selectedChips = [];
      this.getDiaryList();
    }
  },
  mounted() {
    this.getDiaryList();
    this.getTags();
  }
};
</script>

<style scoped></style>
