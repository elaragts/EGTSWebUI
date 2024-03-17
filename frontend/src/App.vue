<script setup lang="ts">
import {RouterView} from 'vue-router'
import Navbar from "@/components/navbar.vue";
import {onMounted} from "vue";
import {useAuthStore} from "@/store/authStore";

const authStore = useAuthStore();
onMounted(() => {
    authStore.initAuth();
});
</script>

<template>
    <div class="flex flex-col min-h-screen">
        <Navbar/>
        <RouterView/>
    </div>
</template>
