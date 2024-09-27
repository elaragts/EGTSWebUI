<script setup lang="ts">
import {ref, watch} from "vue";
import {useDatatableStore} from "@/store/datatableStore";
import {useToast} from "vue-toastification";

const toast = useToast();
const datatableStore = useDatatableStore();

const songs = ref<Record<number, {englishName: string, japaneseName: string}>>(datatableStore.songMap);
// watch to ensure that we update song list once the data loads if it wasn't already loaded before
watch(() => datatableStore.initialized, (a) => songs.value = datatableStore.songMap);

// loading stats with a get request
const totalUsers = ref(0)
const activeUsers = ref(0)
const totalPlayCount = ref(0)
const fetchStats = async () => {
  try {
    const response = await fetch(`/api/stats`);
    if (response.ok) {
      const res = await response.json();
      totalUsers.value = res.totalUsers;
      activeUsers.value = res.activeUsers;
      totalPlayCount.value = res.totalPlayCount;
    } else {
      toast.error("Failed to fetch server stats");
      console.error('Unexpected server stats response status:', response.status);
    }
  } catch (error) {
    toast.error("Failed to fetch server stats");
    console.error('Failed to fetch server stats:', error);
  }
}
fetchStats()


const songSearch = ref("");

</script>

<template>
  <div class="font-sans py-5 text-gray-200 bg-cl2 rounded my-5 mx-5 px-10">
    <h1 class="text-center text-4xl font-bold mb-4 pb-2 border-b">Welcome to EGTSWebUI</h1>

    <div class="relative mb-1 flex flex-row w-full justify-around">
      <div class="flex flex-col rounded-lg border-2 border-cl14 bg-cl5 text-center text-cl1 w-1/3">
        <h1 class="text-2xl text-cl14 py-1 font-bold text-center border-b border-cl14 mx-1">Total Songs</h1>
        <div class="relative w-2/3 self-center">
          <img class="w-full" src="../assets/don.png" alt="Red Taiko Drum"/>
          <div class="absolute inset-0 flex flex-col justify-center items-center">
            <h2 class="text-white text-4xl font-bold">{{Object.entries(songs).length}}</h2>
            <h2 class="text-white text-lg font-bold">SONGS</h2>
          </div>
        </div>
        <h2 class="font-bold text-cl14 text-sm mb-1">WITH REGULAR UPDATES MONTHLY!</h2>
      </div>


      <div class="flex flex-col rounded-lg border-2 border-l-cl14 border-t-cl14 border-r-cl13 border-b-cl13 bg-cl5 text-center text-cl1 w-1/3">
        <h1 class="text-2xl text-purple-800 py-1 font-bold border-b border-b-purple-800 mx-1">Users</h1>

        <div class="relative w-2/3 self-center">
          <img class="w-full" src="../assets/donka.png" alt="Half-red half-blue Taiko drum">
          <div class="absolute inset-0 flex flex-col justify-center items-stretch">
            <div class="w-full flex flex-row justify-around">
              <h2 class="text-white w-1/2 text-4xl ml-14 font-bold">{{totalUsers}}</h2>
              <h2 class="text-white w-1/2 text-4xl mr-14 font-bold">{{activeUsers}}</h2>
            </div>
            <div class="w-full flex flex-row justify-around">
              <h2 class="text-white w-1/2 text-lg ml-14 font-bold">TOTAL</h2>
              <h2 class="text-white w-1/2 text-lg mr-14 font-bold">ACTIVE</h2>
            </div>
          </div>
        </div>

        <h2 class="font-bold text-purple-800 text-sm mb-1">FROM ALL OVER THE WORLD!</h2>
      </div>

      <div class="flex flex-col rounded-lg border-2 border-cl13 bg-cl5 text-center text-cl1 w-1/3">
        <h1 class="text-2xl text-cl13 py-1 font-bold text-center border-b border-cl13 mx-1">Total Play Count</h1>
        <div class="relative w-2/3 self-center">
          <img class="w-full" src="../assets/ka.png" alt="Blue Taiko Drum">
          <div class="absolute inset-0 flex flex-col justify-center items-stretch">
            <h2 class="text-white text-4xl font-bold">{{totalPlayCount}}</h2>
            <h2 class="text-white text-lg font-bold">SCORES SUBMITTED</h2>
          </div>
        </div>
        <h2 class="font-bold text-cl13 text-sm mb-1">AND COUNTING!</h2>
      </div>

      <img class="absolute left-1/4 top-1/3 w-1/6 scale-x-[-1]" src="../assets/Don-chan.webp" alt="Don-chan">
      <img class="absolute right-1/4 top-1/4 w-1/6" src="../assets/Katsu.webp" alt="Katsu-chan">
    </div>


    <div class="flex flex-row w-full justify-around">
      <div class="flex flex-col rounded-lg border-2 border-cl3 bg-cl7 text-cl9 w-1/3">
        <h1 class="text-2xl py-1 font-bold text-center border-b border-cl3 mx-1">Discord</h1>
        <img class="bg-cl1 rounded-3xl mt-5 mb-4 w-1/6 self-center" src="../assets/taiko.png" alt="Elara Community Icon">
        <h2 class="text-center text-gray-300 mb-1">You've been invited to join</h2>
        <h1 class="text-center font-bold text-2xl">Elara Community</h1>
        <div class="flex flex-row w-full justify-center gap-1 mt-1.5 items-center">
          <div class="h-2.5 w-2.5 rounded-full bg-green-600"></div>
          <h2 class="text-center text-gray-300 ">650+ Members</h2>
        </div>
        <div class="w-2/3 h-full flex self-center items-center justify-center">
          <a class="w-full text-center py-3 rounded bg-cl12 border-2 border-cl12" target="_blank" href="https://discord.egts.ca">Join <strong>Elara Community</strong></a>
        </div>
      </div>

      <div class="flex flex-col rounded-lg border-2 border-purple-800 bg-white text-cl1 w-1/3">
        <h1 class="text-2xl py-1 text-purple-800 font-bold text-center border-b border-purple-800 mx-1">Song List</h1>
        <v-card class="h-full flex flex-col text-center" title="Want to check if we have a certain song?" variant="elevated">
          <template v-slot:text>
            <v-text-field
                v-model="songSearch"
                label="Search"
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                hide-details
                single-line
                density="compact"
            ></v-text-field>
          </template>
          <v-data-table
              :items="Object.entries(songs).map(([id, song]) => ({ id: id, englishName: song.englishName }))"
              :headers="[
                        { title: 'ID', align: 'start', value: 'id', width: '10px' },
                        { title: 'Song', align: 'start', value: 'englishName' }
                        ]"
              item-key="id"
              item-value="englishName"
              no-data-text="No songs found"
              items-per-page="5"
              density="compact"
              :search="songSearch"
              :items-per-page-options="[{value: 5, title: '5'}]"
          >
          </v-data-table>
        </v-card>
      </div>

      <div class="flex flex-col rounded-lg border-2 border-white bg-cl10 text-white w-1/3">
        <h1 class="text-2xl py-1 font-bold text-center border-b border-white mx-1">Add to our song list!</h1>
        <a class="flex flex-col items-center justify-center" target="_blank" href="https://github.com/keitannunes/KeifunsDatatableEditor">
          <h2 class="text-xl text-center py-2">Check out <u><strong>KeifunsDatatableEditor (KDE)</strong></u></h2>
          <img class="w-1/6 mb-2" src="../assets/github.png" alt="GitHub Logo">
        </a>

        <div class="flex flex-col h-full w-full self-center items-center justify-center">
          <p class="text-center text-l border-white border-y py-1 mx-1">
            <strong>KeifunsDatatableEditor (KDE)</strong> is a replacement for TaikoSoundEditor (TSE),
            specifically designed for Nijiiro JPN39. KDE allows you to easily modify
            and manage the datatable for your Taiko games, with several additional
            features to streamline the process.
          </p>
          <h1 class="font-bold my-1 text-lg">Features:</h1>
          <ul class="w-full text-sm rounded mb-4 px-8 list-disc">
            <li>
              Edit Datatable: Modify the datatable with ease.
            </li>
            <li>
              Add/Remove Songs: Add or remove songs from the datatable.
            </li>
            <li>
              Auto-Generate Song Details: Automatically generate song details from TJA files.
            </li>
            <li>
              Create Fumen and Sound Files: Generate fumen and sound files directly from TJA.
            </li>
          </ul>
        </div>
      </div>
    </div>

  </div>
</template>