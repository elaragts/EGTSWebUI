<script setup lang="ts">
import {useDatatableStore} from "@/store/datatableStore";
import {useProfileStore} from "@/store/profileStore";
import {type Ref, ref, watch} from "vue";
import type { Item } from '@/types';

const datatableStore = useDatatableStore();
const titles: Ref<Item[]> = ref(datatableStore.titleList);
// watch to ensure that we update titles once the data loads if it wasn't already loaded above
watch(() => datatableStore.initialized, (a) => titles.value = datatableStore.titleList);

const profileStore = useProfileStore();

const search = ref("");
const selectTitleMenu = ref(false);
const selectedTitle = ref(profileStore.profileOptions.title);
watch(() => selectedTitle.value, (value) => profileStore.profileOptions.title = value[0]);

const languageOptions = ["Japanese", "English", "Chinese (Traditional)", "Korean", "Chinese (Simplified)"];

const titlePlateOptions = ["Wood", "Rainbow", "Gold", "Purple", "AI 1", "AI 2", "AI 3", "AI 4"];

const achievementDisplayOptions = ["None", "Easy", "Normal", "Hard", "Oni", "UraOni"];

const difficultyCourseOptions = ["None", "Set up each time", "Easy", "Normal", "Hard", "Oni", "Ura Oni"];

const difficultyStarOptions = ["None", "Set up each time", "1", "2", "3",
                                "4", "5", "6", "7", "8", "9", "10"];

const difficultySortOptions = ["None", "Set up each time", "Default", "Not Cleared",
                                "Not Full Combo", "Not Donderful Combo"];
</script>

<template>
    <div class="bg-cl5 rounded-xl m-1 h-full p-8 flex flex-col">
      <h1 class="text-4xl font-bold mb-5">Profile Options</h1>
      <form class="text-xl flex flex-col mx-2 space-y-2" @submit.prevent="profileStore.uploadProfile()">
          <div class="flex flex-col border-b border-cl6 px-2 mb-4 w-full md:w-1/2">
              <label for="name" class="text-sm mb-1">Name:</label>
              <input
                  class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                  type="text"
                  placeholder="Name"
                  id="name"
                  v-model="profileStore.profileOptions.myDonName"
                  required
                  minlength="1"
                  maxlength="100"
                  @keydown.enter.prevent
              >
          </div>

          <div class="flex flex-row gap-4 mb-4 w-full md:w-1/2">
            <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-3/4">
              <label for="title" class="text-sm mb-1">Title:</label>
              <input
                v-if="profileStore.profileOptions.customTitleOn"
                id="title"
                class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                type="text"
                placeholder="Title"
                v-model="profileStore.profileOptions.title"
                maxlength="200"
              >

              <input
                v-else
                class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                type="text"
                placeholder="Title"
                v-model="profileStore.profileOptions.title"
                maxlength="200"
                disabled
              >

              <v-dialog max-width="1000" class="flex justify-center" v-model="selectTitleMenu">
                <v-card title="Player Titles" variant="elevated">
                  <template v-slot:subtitle>
                    <p class="text-lg font-medium text-cl6">
                      Current Title: {{ profileStore.profileOptions.title }}
                    </p>
                  </template>
                  <template v-slot:text>
                    <v-text-field
                        v-model="search"
                        label="Search"
                        prepend-inner-icon="mdi-magnify"
                        variant="outlined"
                        hide-details
                        single-line
                    ></v-text-field>
                  </template>
                  <v-data-table
                      class="border-y-2 border-gray-200"
                      :search="search"
                      :items="titles.filter(item => item.englishName.length > 0)"
                      :headers="[
                      { title: 'ID', align: 'start', key: 'id', width: '10px' },
                      { title: 'Title', align: 'start', key: 'englishName' }
                      ]"
                      item-key="id"
                      item-value="englishName"
                      no-data-text="No results found"
                      show-select
                      items-per-page="5"
                      :items-per-page-options="[{value: 5, title: '5'}]"
                      elevation="4"
                      select-strategy="single"
                      v-model="selectedTitle"
                  >
                  </v-data-table>
                  <v-card-actions class="ml-auto mr-2 my-2">
                    <button
                            class="px-4 py-2 text-base font-bold text-cl6 rounded-lg border-2 shadow hover:bg-cl6 hover:text-cl5 hover:transition-colors"
                            @click.prevent="selectTitleMenu = false"
                            type="button"
                    >
                      OKAY
                    </button>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </div>

            <label class="inline-flex items-center cursor-pointer m-auto">
              <input type="checkbox" value="true" class="sr-only peer" v-model="profileStore.profileOptions.customTitleOn">
              <div
                  class="relative w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600">
              </div>
              <span class="text-center text-sm font-medium text-cl1">Use Custom Title</span>
            </label>
          </div>

          <div>
            <button v-if="profileStore.profileOptions.customTitleOn"
                    class="p-2 text-base font-bold text-cl3 rounded-lg border-2"
                    type="button"
                    disabled
            >
              SELECT A TITLE
            </button>
            <button v-else
                    class="p-2 text-base font-bold text-cl6 rounded-lg border-2 shadow hover:bg-cl6 hover:text-cl5 hover:transition-colors"
                    type="button"
                    @click.prevent="selectTitleMenu = true"
            >
              SELECT A TITLE
            </button>
          </div>

          <div class="mt-2 flex space-x-5">
              <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/4">
                  <label for="language" class="text-sm mb-1 mt-2">Language</label>
                  <select id="language"
                          v-model="profileStore.profileOptions.language"
                          class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
                    <option v-for="(option, index) in languageOptions" :key="index" :value="index">
                      {{option}}
                    </option>
                  </select>
              </div>
              <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/4">
                  <label for="titlePlate" class="text-sm mb-1 mt-2">Title Plate</label>
                  <select id="titlePlate"
                          v-model="profileStore.profileOptions.titlePlateId"
                          class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
                    <option v-for="(option, index) in titlePlateOptions" :key="index" :value="index">
                      {{option}}
                    </option>
                  </select>
              </div>
          </div>
          <div class="mt-6 flex flex-row space-x-5">
            <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/4">
              <label for="achievementDifficulty" class="text-sm mb-1 mt-6">Achievement Panel Difficulty</label>
              <select id="achievementDifficulty"
                      v-model="profileStore.profileOptions.achievementDisplayDifficulty"
                      class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
                <option v-for="(option, index) in achievementDisplayOptions" :key="index" :value="index">
                  {{option}}
                </option>
              </select>
            </div>
            <div class="flex flex-col px-2 w-full md:w-1/4 gap-2.5 justify-end pb-1">
              <label class="inline-flex items-center cursor-pointer">
                <input type="checkbox" value="true" class="sr-only peer" v-model="profileStore.profileOptions.displayAchievement">
                <div
                    class="relative w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600">
                </div>
                <span class="ms-3 text-sm font-medium text-cl1">Display Achievement Panel</span>
              </label>
              <label class="inline-flex items-center cursor-pointer">
                <input type="checkbox" value="true" class="sr-only peer" v-model="profileStore.profileOptions.displayDan">
                <div
                    class="relative w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600">
                </div>
                <span class="ms-3 text-sm font-medium text-cl1">Display Dan Rank on Name Plate</span>
              </label>
            </div>
          </div>
          <h3 class="font-bold">Difficulty Settings</h3>
          <div class="mt-6 flex space-x-2">
            <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-2/12">
              <label for="courseDifficulty" class="text-sm my-1">Course</label>
              <select id="courseDifficulty"
                      v-model="profileStore.profileOptions.difficultySettingCourse"
                      class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
                <option v-for="(option, index) in difficultyCourseOptions" :key="index" :value="index">
                  {{option}}
                </option>
              </select>
            </div>
            <div class="flex flex-col border-b border-cl6 px-0 w-full md:w-2/12">
              <label for="starDifficulty" class="text-sm my-1">Star</label>
              <select id="starDifficulty"
                      v-model="profileStore.profileOptions.difficultySettingStar"
                      class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
                <option v-for="(option, index) in difficultyStarOptions" :key="index" :value="index">
                  {{option}}
                </option>
              </select>
            </div>
            <div class="flex flex-col border-b border-cl6 pt-0 px-2 w-full md:w-2/12">
              <label for="sortDifficulty" class="text-sm my-1">Sort</label>
              <select id="sortDifficulty"
                      v-model="profileStore.profileOptions.difficultySettingSort"
                      class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
                <option v-for="(option, index) in difficultySortOptions" :key="index" :value="index">
                  {{option}}
                </option>
              </select>
            </div>
          </div>
      </form>
    </div>

</template>

<style scoped>

</style>