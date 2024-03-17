<script setup lang="ts">

import DashboardSidebar from "@/components/organisms/DashboardSidebar.vue";
</script>

<template>
    <div class="flex">
        <DashboardSidebar/>
        <router-view></router-view>
    </div>
</template>

<style scoped>

</style>