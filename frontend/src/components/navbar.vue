<script setup lang="ts">

import NavItem from './atoms/navitem.vue'
import {useAuthStore} from "@/store/authStore";
import {ref, watch} from 'vue';

const authStore = useAuthStore();
const loginTitle = ref('');

// Watch for changes in authentication state
watch(() => authStore.isAuthenticated, (newVal) => {
    if (newVal) {
        loginTitle.value = authStore.username;
    } else {
        loginTitle.value = 'Login';
    }
}, {immediate: true});

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
            <NavItem title="Dashboard" to="/register"/>
            <NavItem v-if="!authStore.isAuthenticated" :title="loginTitle" to="/login"/>
            <button @click="toggleDropdown"
                    class="text-white bg-cl6 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                    type="button">
                Dropdown button
                <svg class="w-2.5 h-2.5 ms-3" fill="none" viewBox="0 0 10 6">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                          d="m1 1 4 4 4-4"/>
                </svg>
            </button>
        </div>
    </nav>
    <div v-show="dropdownVisible"
         class="z-10 bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700 float-right m-2">
        <ul class="py-2 text-sm text-gray-700 dark:text-gray-200">
            <li>
                <a href="#"
                   class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Dashboard</a>
            </li>
            <li>
                <a href="#"
                   class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Settings</a>
            </li>
            <li>
                <a href="#"
                   class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Earnings</a>
            </li>
            <li>
                <a href="#"
                   class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Sign
                    out</a>
            </li>
        </ul>
    </div>
</template>


<script lang="ts">
export default {
    data() {
        return {
            dropdownVisible: false
        };
    },
    methods: {
        toggleDropdown() {
            this.dropdownVisible = !this.dropdownVisible;
        }
    }
};
</script>