<script setup lang="ts">
import "@/index.css"
</script>

<template>

    <!--    Someone make this work please! thanks // -->
    <div class="sm:hidden">
        <label for="tabs" class="sr-only">Profile</label>
        <select id="tabs"
                class="bg-cl2 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5">
            <option>Profile</option>
            <option>Costume</option>
            <option>Song Options</option>
        </select>
    </div>
    <ul class="hidden text-sm font-medium text-center rounded-lg shadow sm:flex m-1">
        <li class="w-full focus-within:z-10">
            <RouterLink to="/dashboard/edit/profile"
               class="inline-block w-full p-4 bg-white hover:text-gray-700 border-r border-gray-200 rounded-s-lg active focus:outline-none">Profile</RouterLink>
        </li>
        <li class="w-full focus-within:z-10">
            <RouterLink to="/dashboard/edit/costume"
               class="inline-block w-full p-4 bg-white border-r border-gray-200 hover:text-gray-700 hover:bg-gray-50 focus:outline-none">Costume</RouterLink>
        </li>
        <li class="w-full focus-within:z-10">
            <RouterLink to="/dashboard/edit/song-options"
               class="inline-block w-full p-4 bg-white border-s-0 border-gray-200 rounded-e-lg hover:text-gray-700 hover:bg-gray-50">Song Options</RouterLink>
        </li>
    </ul>

</template>

<style scoped>
.router-link-active {
    @apply text-gray-900 bg-gray-100
}
</style>