<template>
  <v-card class="" min-height="400">
    <v-card-title class="d-flex justify-center">
      <span class="title font-weight-regular">{{ blog.title }}</span>
    </v-card-title>

    <v-card-text>
      <div class="text-h6 font-weight-regular ml-1 mr-1">
        {{ blog.summary }}
      </div>
    </v-card-text>

    <v-card-actions>
      <v-list-item class="grow">
        <v-row align="center" justify="start">
          <v-list-item-avatar color="grey darken-3">
            <v-img
              class="elevation-6"
              src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
            ></v-img>
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title>Evan You</v-list-item-title>
          </v-list-item-content>
        </v-row>
        <v-row align="center" justify="center">
          <v-chip
            v-for="(chip, index) in chips"
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

        <v-row align="center" justify="end">
          <v-icon class="mr-1">mdi-eye-outline</v-icon>
          <span class="subheading mr-2">256</span>
          <span class="mr-1">·</span>
          <v-icon class="mr-1">mdi-heart-outline</v-icon>
          <span class="subheading">45</span>
        </v-row>
      </v-list-item>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'PrivateBlogPreviewCard',
  props: ['blog']
};
</script>
