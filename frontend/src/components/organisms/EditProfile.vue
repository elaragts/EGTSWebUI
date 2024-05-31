<script setup lang="ts">
    let checked = false;
</script>

<template>
    <div class="bg-cl5 rounded-xl m-1 h-full p-10 flex flex-col">
        <h1 class="text-4xl font-bold mb-10">Profile Options</h1>
        <form class="text-xl flex flex-col mx-2 space-y-2">
            <div class="flex flex-col border-b border-cl6 px-2 mb-4 w-full md:w-1/2">
                <label class="text-sm mb-1">Name:</label>
                <input
                    class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                    type="text" placeholder="Name">
            </div>
            <label class="inline-flex items-center cursor-pointer">
                <input type="checkbox" value="true" class="sr-only peer" v-model="checked">
                <div
                    class="relative w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>
                <span class="ms-3 text-sm font-medium text-cl1">Use Custom Title</span>
            </label>
            <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/2">
                <label for="title" class="text-sm mb-1">Title:</label>
                <input
                    id="title"
                    class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                    type="text" placeholder="Title">
            </div>
            <div class="mt-6 flex space-x-5">
                <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/4">
                    <label for="language" class="text-sm mb-1 mt-6">Language</label>
                    <select id="language"
                            class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
                        <option value="0">Japanese</option>
                        <option value="1">English</option>
                        <option value="2">Chinese (Traditional)</option>
                        <option value="3">Korean</option>
                        <option value="4">Chinese (Simplified)</option>
                    </select>
                </div>
                <div class="flex flex-col border-b border-cl6 px-2 w-full md:w-1/4">
                    <label for="titlePlate" class="text-sm mb-1 mt-6">Title Plate</label>
                    <select id="language"
                            class="w-full bg-transparent focus:outline-none focus:ring-0 focus:border-gray-200 h-10 text-base">
                        <option value="0">Wood</option>
                        <option value="1">Rainbow</option>
                        <option value="2">Gold</option>
                        <option value="3">Purple</option>
                        <option value="4">AI 1</option>
                        <option value="5">AI 2</option>
                        <option value="6">AI 3</option>
                        <option value="7">AI 4</option>
                    </select>
                </div>
            </div>
        </form>
    </div>

</template>

<style scoped>

</style>