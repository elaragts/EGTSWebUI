<script setup lang="ts">
import {useRouter} from 'vue-router';
import {ref} from 'vue';
import {useToast} from 'vue-toastification';

const toast = useToast();
const accessCode = ref('');
const username = ref('');
const password = ref('');
const cPassword = ref('');
const errorMessage = ref('');
const router = useRouter();

const submitForm = async () => {
    if (!username.value || !password.value || !accessCode.value) {
        toast.error('Please fill in all fields');
        return;
    }
    if (password.value !== cPassword.value) {
        toast.error('Passwords do not match');
        return;
    }
    const formData = new URLSearchParams();
    formData.append('accessCode', accessCode.value);
    formData.append('username', username.value);
    formData.append('password', password.value);

    const response = await fetch('/auth/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: formData
    });
    if (!response.ok) {
        toast.error(await response.text());
        return;
    }
    toast.success('Account created successfully');
    await router.push({name: 'login'});
};

</script>

<template>
    <div v-if="errorMessage" class="text-red-500">{{ errorMessage }}</div>
    <form class="space-y-6 font-[Roboto]" @submit.prevent="submitForm">
        <div>
            <label for="accessCode" class="block mb-2 text-xl text-cl5">Access Code</label>
            <input type="text" name="accessCode" v-model="accessCode"
                   class="bg-gray-50 border border-gray-300  sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                   placeholder="Access Code"/>
        </div>
        <div>
            <label for="username" class="block mb-2 text-xl text-cl5">Username</label>
            <input type="text" name="username" v-model="username"
                   class="bg-gray-50 border border-gray-300  sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                   placeholder="username"/>
        </div>
        <div>
            <label for="password" class="block mb-2 text-xl text-cl5 ">Password</label>
            <input type="password" name="password" v-model="password"
                   class="bg-gray-50 border border-gray-300 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                   placeholder="password"/>
        </div>
        <div>
            <label for="cPassword" class="block mb-2 text-xl text-cl5 ">Confirm Password</label>
            <input type="password" name="cPassword" v-model="cPassword"
                   class="bg-gray-50 border border-gray-300 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                   placeholder="password"/>
        </div>
        <button
            class="w-full text-white bg-cl6 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center">
            Register
        </button>
    </form>
</template>
