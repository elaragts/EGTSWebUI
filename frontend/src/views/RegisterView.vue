<script setup lang="ts">
import RegisterForm from '../components/organisms/RegisterForm.vue'
import {useAuthStore} from "@/store/authStore";
import router from "@/router";
const authStore = useAuthStore();
if (authStore.isAuthenticated) {
    router.push({name: 'home'});
}

</script>

<template>
    <div class="flex flex-col items-center justify-center px-6 py-5 mx-auto md:my-16 lg:py-0">
        <div class="flex flex-col items-center space-y-5 mb-6 text-4xl">
            <img class="w-20 h-20 mr-2" src="../assets/taiko.png" alt="logo"/>
            <h1 class="font-semibold font-[Audiowide] text-center">Elara Global Taiko Server</h1>
        </div>
        <div class="w-full bg-cl2 rounded-lg shadow md:mt-0 sm:max-w-md xl:p-0">
            <div class="p-6 space-y-4 md:space-y-6 sm:p-8 ">
                <h1 class="text-3xl font-bold center-self leading-tight tracking-tight text-cl5 text-bold">
                    Register
                </h1>
                <div class="space-y-4">
                    <RegisterForm/>
                    <p class="text-sm font-light text-cl5">
                        Already have an account?
                        <router-link to="/login" class="font-medium text-primary-600 hover:underline">Log in here
                        </router-link>
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>

</style>