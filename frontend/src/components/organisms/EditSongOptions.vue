<script setup lang="ts">
import {useProfileStore} from "@/store/profileStore";

const profileStore = useProfileStore();

const speedOptions = ["1.0", "1.1", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7",
                    "1.8", "1.9", "2.0", "2.5", "3.0", "3.5", "4.0"];

const randomOptions = ["None", "Whimsical", "Messy"];

const toneOptions = ["Taiko", "Festival", "Dogs & Cats", "Deluxe", "Drumset", "Tambourine", "Don Wada",
                    "Clapping", "Conga", "8-Bit", "Heave-ho", "Mecha", "Bujain", "Rap", "Hosogai",
                    "Akemi", "Synth Drum", "Shuriken", "Bubble Pop", "Electric Guitar"];

const notesPositionOptions = ["-5", "-4", "-3", "-2", "-1", "0", "+1", "+2", "+3", "+4", "+5"];

</script>


<template>
  <div class="bg-cl5 rounded-xl m-1 h-full p-8 flex flex-col">
    <h1 class="text-4xl font-bold mb-5">Song Options</h1>
    <form class="text-xl flex flex-row mx-2 h-full" @submit.prevent="profileStore.uploadProfile()">
      <div class="flex flex-col w-1/2 justify-between">
        <div class="flex flex-row w-full">
          <div class="flex flex-row w-full">
            <label class="inline-flex items-center cursor-pointer">
              <input type="checkbox" value="true" class="sr-only peer" v-model="profileStore.songOptions.isVanishOn">
              <div
                  class="relative w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600">
              </div>
              <span class="ms-3 font-normal text-cl1">Vanish</span>
            </label>
          </div>
          <div class="flex flex-row w-full">
            <label class="inline-flex items-center cursor-pointer">
              <input type="checkbox" value="true" class="sr-only peer" v-model="profileStore.songOptions.isInverseOn">
              <div
                  class="relative w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600">
              </div>
              <span class="ms-3 font-normal text-cl1">Inverse</span>
            </label>
          </div>
        </div>

        <div class="flex flex-row w-full">
          <div class="flex flex-row w-full">
            <label class="inline-flex items-center cursor-pointer">
              <input type="checkbox" value="true" class="sr-only peer" v-model="profileStore.songOptions.isSkipOn">
              <div
                  class="relative w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600">
              </div>
              <span class="ms-3 font-normal text-cl1">Give Up</span>
            </label>
          </div>
          <div class="flex flex-row w-full">
            <label class="inline-flex items-center cursor-pointer">
              <input type="checkbox" value="true" class="sr-only peer" v-model="profileStore.songOptions.isVoiceOn">
              <div
                  class="relative w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600">
              </div>
              <span class="ms-3 font-normal text-cl1">Voice</span>
            </label>
          </div>
        </div>

        <div class="flex flex-col border-b border-cl6 px-2 w-full">
          <label for="speed" class="mt-3">Speed</label>
          <select id="speed"
                  v-model="profileStore.songOptions.speedId"
                  class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
            <option v-for="(option, index) in speedOptions" :key="index" :value="index">
              {{ option }}
            </option>
          </select>
        </div>
        <div class="flex flex-col border-b border-cl6 px-2 w-full">
          <label for="random" class="mt-3">Random</label>
          <select id="random"
                  v-model="profileStore.songOptions.randomId"
                  class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
            <option v-for="(option, index) in randomOptions" :key="index" :value="index">
              {{option}}
            </option>
          </select>
        </div>
        <div class="flex flex-col border-b border-cl6 px-2 w-full">
          <label for="tone" class="mt-3">Tone</label>
          <select id="tone"
                  v-model="profileStore.songOptions.selectedToneId"
                  class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
            <option v-for="(option, index) in toneOptions" :key="index" :value="index">
              {{ option }}
            </option>
          </select>
        </div>
        <div class="flex flex-col w-full">
          <label for="position" class="my-1">Notes Position</label>
          <v-slider
              id="position"
              v-model="profileStore.songOptions.notesPosition"
              step="1"
              show-ticks="always"
              :thumb-size="15"
              tick-size="4"
              track-color="#3185FC"
              track-fill-color="#3185FC"
              thumb-color="#374151"
              track-size="5"
              :min="-5"
              :max="5"
              :ticks="notesPositionOptions">
          </v-slider>
        </div>
      </div>
      <div class="flex flex-col w-1/2 h-full justify-center">
        <button
            class="w-1/2 self-center p-2 text-base font-bold border-cl6 bg-cl6 text-cl5 rounded-lg border-2 shadow hover:bg-opacity-80 hover:transition-all"
            type="submit"
        >
          Save
        </button>
      </div>
    </form>
  </div>
</template>

<style scoped>

</style>