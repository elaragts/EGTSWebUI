<script setup lang="ts">

import NavItem from './atoms/navitem.vue'
import {useAuthStore} from "@/store/authStore";

const authStore = useAuthStore();

// // Watch for changes in authentication state
// watch(() => authStore.isAuthenticated, (newVal) => {
//     if (newVal) {
//         loginTitle.value = authStore.username;
//     }
// }, {immediate: true});

</script>

<template>
    <nav
        class="flex flex-col md:flex-row text-center sm:text-left sm:justify-between py-4 px-6 bg-cl1 shadow sm:items-baseline w-full items-center md:items-center">
        <div class="mb-2 sm:mb-0 flex items-center justify-center md:justify-start w-full md:w-auto">
            <router-link to="/" class="flex items-center justify-center md:justify-start">
                <img src="../assets/taiko.png" class="mr-3 h-10 hidden md:block" alt="Logo"/>
                <span class="text-2xl font-semibold whitespace-nowrap text-cl5 font-[Taiko] text-center md:text-left">Taiko Public Server</span>
            </router-link>
        </div>
        <div
            class="text-xl flex flex-col md:flex-row p-4 mb:2 md:p-0 justify-center md:justify-end md:items-center space-x-0 md:space-x-4 w-full md:w-auto">
            <NavItem title="Guide" to="/"/>
            <NavItem title="Leaderboard" to="/logout"/>
            <NavItem v-if="!authStore.isAuthenticated" title="Log In" to="/login"/>
            <NavItem v-if="authStore.isAuthenticated" title="Dashboard" to="/profile"/>
        </div>
    </nav>
</template>
