<template>
  <v-footer light padless>
    <v-card class="flex" flat tile>
      <v-card-title>
        <v-row
          justify="center"
          align="center"
          class="d-none d-sm-flex d-md-flex d-lg-flex d-xl-flex"
        >
          <div class="text-h6">
            Copyright © {{ new Date().getFullYear() }} Kiroku.cn 备案号:<a
              href="http://beian.miit.gov.cn"
              style="decoration: none; color:teal;"
              >渝ICP备18015002号-3</a
            >
          </div>
        </v-row>
        <v-row justify="center" align="center" class="d-flex d-sm-none">
          <v-col cols="12" class="d-flex justify-center">
            <div class="text-h6">
              Copyright © {{ new Date().getFullYear() }} Kiroku.cn
            </div>
          </v-col>
          <v-col cols="12" class="d-flex justify-center">
            <div class="text-h6">
              备案号:<a
                href="http://beian.miit.gov.cn"
                style="decoration: none; color:teal;"
                >渝ICP备18015002号-3</a
              >
            </div>
          </v-col>
        </v-row>
      </v-card-title>
    </v-card>
  </v-footer>
</template>

<script>
export default {
  name: 'Footer'
};
</script>
