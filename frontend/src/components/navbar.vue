<script setup lang="ts">
import NavItem from './atoms/navitem.vue'
</script>

<template>
  <nav
      class="flex flex-col text-center sm:flex-row sm:text-left sm:justify-between py-4 px-6 bg-cl1 shadow sm:items-baseline w-full">
    <div class="mb-2 sm:mb-0">
      <router-link to="/" class="flex items-center">
        <img src="../assets/taiko.png" class="mr-3 h-10" alt="Logo"/>
        <span class="text-2xl font-semibold whitespace-nowrap text-cl5 font-[Taiko] self-center">Taiko Public Server</span>
      </router-link>
    </div>
    <div class="text-xl flex flex-row p-4 mb:2 md:p-0 self-center space-x-4">
      <NavItem title="Leaderboard" to="/"/>
      <NavItem title="Dashboard" to="/register"/>
      <NavItem title="Profile" to="/login"/>
    </div>
  </nav>
</template>