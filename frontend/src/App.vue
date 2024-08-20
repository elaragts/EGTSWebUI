<script setup lang="ts">
import {RouterView} from 'vue-router'
import Navbar from "@/components/organisms/navbar.vue";
import {onMounted} from "vue";
import {useAuthStore} from "@/store/authStore";
import {useDatatableStore} from "@/store/datatableStore";

const authStore = useAuthStore();
const datatableStore = useDatatableStore();
onMounted(() => {
    authStore.initAuth();
    datatableStore.initDatatable();
});
</script>

<template>
    <div class="flex flex-col min-h-screen">
        <Navbar/>
        <RouterView/>
    </div>
</template>
