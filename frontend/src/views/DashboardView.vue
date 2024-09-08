<script setup lang="ts">
import { useAuthStore } from "@/store/authStore";
import DashboardSidebar from "@/components/organisms/DashboardSidebar.vue";
import {watch} from "vue";
import router from "@/router";

const authStore = useAuthStore();

// Make sure user is logged in
if (authStore.initialized && !authStore.isAuthenticated) {
  router.push({name: 'login'})
} else {
  watch(() => authStore.initialized, () => {
    if (!authStore.isAuthenticated) {
      router.push({name: 'login'})
    }
  });
}

</script>

<template>
    <div class="flex">
        <DashboardSidebar/>
        <router-view/>
    </div>
</template>

<style scoped>

</style>