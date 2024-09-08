<script setup lang="ts">
import { useDatatableStore } from '@/store/datatableStore';
import {ref, watch} from "vue";
import {useAuthStore} from "@/store/authStore";
import {useToast} from "vue-toastification";

const authStore = useAuthStore();
const datatableStore = useDatatableStore();
const toast = useToast();

const songs = ref<Record<number, {englishName: string, japaneseName: string}>>(datatableStore.songMap);
// watch to ensure that we update song list once the data loads if it wasn't already loaded before
watch(() => datatableStore.initialized, (a) => songs.value = datatableStore.songMap);

// loading favourited songs with a get request
const favouritedSongs = ref<number[]>([]);
const fetchFavouritedSongs = async () => {
  try {
    const response = await fetch(`/api/user/${authStore.baid}/songs`);
    if (response.ok) {
      favouritedSongs.value = await response.json();
    } else {
      toast.error("Failed to fetch favourited songs");
      console.error('Unexpected favourited songs data response status:', response.status);
    }
  } catch (error) {
    toast.error("Failed to fetch favourited songs");
    console.error('Failed to fetch favourited songs:', error);
  }
}
// calling the above function when baid loads
if (authStore.initialized) {
  fetchFavouritedSongs()
} else {
  watch(() => authStore.initialized, (value) => {
    fetchFavouritedSongs();
  });
}

async function addFavouritedSong(songId: number): Promise<boolean> {
  const data = {
    songId: songId
  }

  const response = await fetch(`/api/user/${authStore.baid}/songs`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });


  if (response.ok) {
    toast.success('Song Favourited Successfully');
    favouritedSongs.value.push(songId);
    return true;
  } else {
    toast.error(`Failed to favourite song: ${await response.text()}`);
    return false;
  }
}

async function deleteFavouritedSong(songId: number): Promise<boolean> {
  const data = {
    songId: songId
  }

  const response = await fetch(`/api/user/${authStore.baid}/songs`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });


  if (response.ok) {
    toast.success('Removed Favourited Song Successfully');
    favouritedSongs.value = favouritedSongs.value.filter(currentSong => currentSong != songId)
    return true;
  } else {
    toast.error(`Failed to remove favourited song: ${await response.text()}`);
    return false;
  }
}

const favSongsSearch = ref("");
const allSongsSearch = ref("");

</script>

<template>
  <div class="w-full m-2 flex flex-col">
    <div class="bg-cl5 rounded-xl m-1 h-full p-8 flex flex-col">
      <h1 class="text-4xl text-center border-b border-cl6 pb-2 font-bold mb-5">Favourited Songs</h1>
      <div class="text-xl flex flex-row mx-2 gap-1 h-full">
        <div class="h-full w-1/2 rounded-xl">
          <v-card class="h-full flex flex-col" title="Favourited Songs" variant="elevated">
            <template v-slot:text>
              <v-text-field
                  v-model="favSongsSearch"
                  label="Search"
                  prepend-inner-icon="mdi-magnify"
                  variant="outlined"
                  hide-details
                  single-line
              ></v-text-field>
            </template>
            <v-data-table
                :items="favouritedSongs.map((songID: number) => ({ id: songID, englishName: songs[songID]?.englishName || 'Unknown Song' }))"
                :headers="[
                        { title: 'ID', align: 'start', value: 'id', width: '10px' },
                        { title: 'Song', align: 'start', value: 'englishName' },
                        { title: '', align: 'end', value: 'actions', sortable: false }
                        ]"
                item-key="id"
                item-value="englishName"
                no-data-text="No songs found"
                items-per-page="7"
                density="comfortable"
                :search="favSongsSearch"
                :items-per-page-options="[{value: 7, title: '7'}]"
            >
              <template v-slot:item.actions="{ item }">
                <v-icon
                    size="small"
                    icon="mdi-delete"
                    @click="deleteFavouritedSong(item.id)"
                >
                </v-icon>
              </template>
            </v-data-table>
          </v-card>
        </div>

        <div class="h-full w-1/2 rounded-xl">
          <v-card class="h-full" title="All Songs" variant="elevated">
            <template v-slot:text>
              <v-text-field
                  v-model="allSongsSearch"
                  label="Search"
                  prepend-inner-icon="mdi-magnify"
                  variant="outlined"
                  hide-details
                  single-line
              ></v-text-field>
            </template>
            <v-data-table
                :items="Object.entries(songs).map(([id, song]) => ({ id: id, englishName: song.englishName }))"
                :headers="[
                            { title: 'ID', align: 'start', value: 'id', width: '10px' },
                            { title: 'Song', align: 'start', value: 'englishName' },
                            { title: '', align: 'end', value: 'actions', sortable: false }
                            ]"
                item-key="id"
                item-value="englishName"
                no-data-text="No songs found"
                items-per-page="7"
                density="comfortable"
                :search="allSongsSearch"
                :items-per-page-options="[{value: 7, title: '7'}]"
            >
              <template v-slot:item.actions="{ item }">
                <v-icon
                  v-if="!(favouritedSongs.includes(parseInt(item.id)))"
                  size="small"
                  icon="mdi-star-outline"
                  @click="addFavouritedSong(parseInt(item.id))"
                >
                </v-icon>
                <v-icon
                  v-else
                  size="small"
                  icon="mdi-star"
                  @click="deleteFavouritedSong(parseInt(item.id))"
                >
                </v-icon>
              </template>
            </v-data-table>
          </v-card>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>