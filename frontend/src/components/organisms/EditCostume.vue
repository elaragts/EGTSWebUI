<script setup lang="ts">
import {useDatatableStore} from "@/store/datatableStore";
import {useProfileStore} from "@/store/profileStore";
import {type Ref, ref, watch} from "vue";
import type { Item } from '@/types';
import { VSwatches } from 'vue3-swatches'
import 'vue3-swatches/dist/style.css'

const colorOptions = [
  "#F84828", "#68C0C0", "#DC1500", "#F8F0E0", "#009687", "#00BF87", "#00FF9A", "#66FFC2",
  "#FFFFFF", "#690000", "#FF0000", "#FF6666", "#FFB3B3", "#00BCC2", "#00F7FF", "#66FAFF",
  "#B3FDFF", "#E4E4E4", "#993800", "#FF5E00", "#FF9E78", "#FFCFB3", "#005199", "#0088FF",
  "#66B8FF", "#B3DBFF", "#B9B9B9", "#B37700", "#FFAA00", "#FFCC66", "#FFE2B3", "#000C80",
  "#0019FF", "#6675FF", "#B3BAFF", "#858585", "#B39B00", "#FFDD00", "#FFFF00", "#FFFF71",
  "#2B0080", "#5500FF", "#9966FF", "#CCB3FF", "#505050", "#38A100", "#78C900", "#B3FF00",
  "#DCFF8A", "#610080", "#C400FF", "#DC66FF", "#EDB3FF", "#232323", "#006600", "#00B800",
  "#00FF00", "#8AFF9E", "#990059", "#FF0095", "#FF66BF", "#FFB3DF", "#000000"
];

const datatableStore = useDatatableStore();
const headOptions: Ref<Item[]> = ref(datatableStore.headList);
const bodyOptions: Ref<Item[]> = ref(datatableStore.bodyList);
const faceOptions: Ref<Item[]> = ref(datatableStore.faceList);
const kigurumiOptions: Ref<Item[]> = ref(datatableStore.kigurumiList);
const puchiOptions: Ref<Item[]> = ref(datatableStore.puchiList);

// watch to ensure that we update the lists once the data loads if it wasn't loaded above
watch(() => datatableStore.initialized, (a) => {
  headOptions.value = datatableStore.headList
  bodyOptions.value = datatableStore.bodyList
  faceOptions.value = datatableStore.faceList
  kigurumiOptions.value = datatableStore.kigurumiList
  puchiOptions.value = datatableStore.puchiList
});

const profileStore = useProfileStore();

// due to vue swatches component, we need to adjust some stuff to model it
// when we submit the form, we need to convert these back into id/numbers
const swatchColorBody = ref(colorOptions[profileStore.costumeOptions.colorBody]);
const swatchColorFace = ref(colorOptions[profileStore.costumeOptions.colorFace]);
const swatchColorLimb = ref(colorOptions[profileStore.costumeOptions.colorLimb]);

// update variables in case data wasn't loaded before
watch(() => profileStore.initialized, (a) => {
  swatchColorBody.value = colorOptions[profileStore.costumeOptions.colorBody];
  swatchColorFace.value = colorOptions[profileStore.costumeOptions.colorFace];
  swatchColorLimb.value = colorOptions[profileStore.costumeOptions.colorLimb];
});

watch(() => swatchColorBody.value, (val) => {
  profileStore.costumeOptions.colorBody = colorOptions.indexOf(val)
})

watch(() => swatchColorFace.value, (val) => {
  profileStore.costumeOptions.colorFace = colorOptions.indexOf(val)
})

watch(() => swatchColorLimb.value, (val) => {
  profileStore.costumeOptions.colorLimb = colorOptions.indexOf(val)
})

</script>

<template>
  <div class="bg-cl5 rounded-xl m-1 h-full p-8 flex flex-col">
    <h1 class="text-4xl font-bold mb-5">Costume Options</h1>
    <form class="text-xl flex flex-col mx-2 space-y-2" @submit.prevent="profileStore.uploadProfile()">
      <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/2">
        <label for="head" class="text- mt-0.5">Head</label>
        <select id="head"
                v-model="profileStore.costumeOptions.currentHead"
                class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
          <option v-for="(option) in headOptions" :key="option.id" :value="option.id">
            {{ option.id + " - " + option.englishName }}
          </option>
        </select>
      </div>
      <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/2">
        <label for="body" class="text-m mt-0.5">Body</label>
        <select id="body"
                v-model="profileStore.costumeOptions.currentBody"
                class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
          <option v-for="(option) in bodyOptions" :key="option.id" :value="option.id">
            {{ option.id + " - " + option.englishName }}
          </option>
        </select>
      </div>
      <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/2">
        <label for="face" class="text-m mt-0.5">Face</label>
        <select id="face"
                v-model="profileStore.costumeOptions.currentFace"
                class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
          <option v-for="(option) in faceOptions" :key="option.id" :value="option.id">
            {{ option.id + " - " + option.englishName }}
          </option>
        </select>
      </div>
      <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/2">
        <label for="kigurumi" class="text-m mt-0.5">Kigurumi</label>
        <select id="kigurumi"
                v-model="profileStore.costumeOptions.currentKigurumi"
                class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
          <option v-for="(option) in kigurumiOptions" :key="option.id" :value="option.id">
            {{ option.id + " - " + option.englishName }}
          </option>
        </select>
      </div>
      <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/2">
        <label for="puchi" class="text-m mt-0.5">Puchi</label>
        <select id="puchi"
                v-model="profileStore.costumeOptions.currentPuchi"
                class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
          <option v-for="(option) in puchiOptions" :key="option.id" :value="option.id">
            {{ option.id + " - " + option.englishName }}
          </option>
        </select>
      </div>

      <h3 class="font-bold">Color Settings</h3>
      <div class="flex">
        <div class="flex flex-col text-center w-full md:w-2/12">
          <label for="bodyColor" class="text-sm mb-1">Body</label>
          <VSwatches id="bodyColor"
                     v-model="swatchColorBody"
                     :swatches="colorOptions"
                     show-border
                     swatch-size="30"
                     :spacing-size=3
                     shapes="circles"
                     row-length=5
          />
        </div>
        <div class="flex flex-col text-center w-full md:w-2/12">
          <label for="faceColor" class="text-sm mb-1">Face</label>
          <VSwatches id="faceColor"
                     v-model="swatchColorFace"
                     :swatches="colorOptions"
                     show-border
                     swatch-size="30"
                     :spacing-size=3
                     shapes="circles"
                     row-length=5
          />
        </div>
        <div class="flex flex-col text-center w-full md:w-2/12">
          <label for="limbColor" class="text-sm mb-1">Limb</label>
          <VSwatches id="limbColor"
                     v-model="swatchColorLimb"
                     :swatches="colorOptions"
                     show-border
                     swatch-size="30"
                     :spacing-size=3
                     shapes="circles"
                     row-length=5
          />
        </div>
      </div>
    </form>
  </div>
</template>

<style scoped>

</style>