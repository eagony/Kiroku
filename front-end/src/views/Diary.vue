<template>
  <v-container>
    <v-row dense>
      <v-col cols="12">
        <v-card>
          <v-card-title v-if="!adding">
            <div class="text-justify text--black">共100篇日记</div>
            <v-spacer></v-spacer>
            <div>
              <v-btn class="mx-2" fab dark color="teal" @click="adding = true">
                <v-icon dark>mdi-pencil</v-icon>
              </v-btn>
            </div>
          </v-card-title>

          <v-card-text v-if="adding">
            <v-textarea
              auto-grow
              counter
              name="input-7-4"
              label="新增日记"
            ></v-textarea>
            <v-row class="ml-1">
              <v-chip-group
                column
                v-for="chip in chips"
                :key="chip.name"
                multiple
                v-model="selections"
                active-class="teal--text text--accent-4"
              >
                <v-chip :value="chip.name">{{ chip.name }}</v-chip>
              </v-chip-group>
              {{ selections }}
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
            <v-btn
              class="white--text"
              color="teal accent-4"
              @click="adding = false"
            >
              完成
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col v-for="(item, i) in 4" :key="i" cols="12">
        <v-card class="mx-auto">
          <v-card-title>
            <p class="text-justify">{{ date }}({{ day }})</p>
          </v-card-title>

          <v-card-text>
            <div class="text-justify text-body-1 text--black">
              很真实地欣赏着自己的心 偶然间我与你面对面坐下来 而你只是说出两个字
              明天 现在我们俩只能在静寂中煎熬 悲哀于瞬息之间 其实人生并不重要
              和遥远的童年一样 在那恬美爱情的时候 我们梦想着每一种不同质的东西
              面对一个梦萧萧的秋季 看见有生命力的枯叶 或在沉默之中开口
              为你描述故事的主角 只是每次要等到经历最困难的时候
              我们才能学会如何去歪曲生命的定义 岁月是没有承诺的
              我们靠各种努力才能建立起 毫无价值的精神财富 在透明的景色里
              一层一层将自己包装起来
            </div>
          </v-card-text>

          <v-divider class="mt-6 mx-4"></v-divider>

          <v-card-text>
            <v-chip
              v-for="chip in chips"
              :key="chip.name"
              class="ma-2"
              :color="chip.color"
              label
              text-color="white"
            >
              <v-icon left>mdi-label</v-icon>
              {{ chip.name }}
            </v-chip>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import moment from '../plugins/moment';

export default {
  name: 'Diary',
  data: () => ({
    date: moment().format('YYYY.MM.DD'),
    day: moment().format('dddd'),
    adding: false,
    chips: [
      {
        name: '开心',
        color: 'red',
        icon: ''
      },
      {
        name: '失落',
        color: 'grey',
        icon: ''
      },
      {
        name: '晴朗',
        color: 'blue',
        icon: ''
      },
      {
        name: '大雨',
        color: 'green',
        icon: ''
      },
      {
        name: '电影',
        color: 'orange',
        icon: ''
      }
    ],
    selections: []
  }),
  methods: {}
};
</script>

<style scoped></style>
