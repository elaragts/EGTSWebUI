<script setup lang="ts">

import NavItem from '../atoms/navitem.vue'
import {useAuthStore} from "@/store/authStore";

const authStore = useAuthStore();

</script>

<template>
    <nav
        class="flex flex-col h-16 md:flex-row text-center sm:text-left sm:justify-between py-4 px-6 bg-cl1 shadow sm:items-baseline w-full items-center md:items-center">
        <div class="mb-2 sm:mb-0 flex items-center justify-center md:justify-start w-full md:w-auto">
            <router-link to="/" class="flex items-center justify-center md:justify-start">
                <img src="../../assets/taiko.png" class="mr-3 h-10 hidden md:block" alt="Logo"/>
                <span class="text-2xl font-[Audiowide] whitespace-nowrap text-cl5 text-center md:text-left">Elara Global Taiko Server</span>
            </router-link>
        </div>
        <div
            class="text-xl flex flex-col md:flex-row p-4 mb:2 md:p-0 justify-center md:justify-end md:items-center space-x-0 md:space-x-4 w-full md:w-auto">
            <a href="https://legacy.egts.ca/" class="block py-2 px-3 text-cl5 rounded font-[Roboto] md:text-2xl sm:text-xl md:bg-transparent md:p-0 hover:text-cl6 mx-auto md:mx-0">LegacyWebUI</a>
            <NavItem title="Guide" to="/guide"/>
<!--            <NavItem title="Leaderboard" to="/leaderboard"/>-->
            <NavItem v-if="!authStore.isAuthenticated" title="Log In" to="/login"/>
            <NavItem v-if="authStore.isAuthenticated" title="Dashboard" to="/dashboard/profile"/>
        </div>
    </nav>
</template>
