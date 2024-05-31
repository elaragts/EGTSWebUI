<script setup lang="ts">

import EditTabs from "@/components/organisms/EditTabs.vue";
</script>

<template>
    <div class="w-full m-2 flex flex-col">
        <EditTabs/>
        <router-view/>
    </div>
</template>

<style scoped>

</style>