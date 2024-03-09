<script setup lang="ts">
import { useRouter } from 'vue-router';
import { ref } from 'vue';

const username = ref('');
const password = ref('');
const errorMessage = ref('');
const router = useRouter();

const submitForm = async () => {
  const formData = new URLSearchParams();
  formData.append('username', username.value);
  formData.append('password', password.value);

  const response = await fetch('/auth/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: formData,
  });

  if (response.ok) {
    await router.push('/');
  } else {
    const errorData = await response.text();
    errorMessage.value = errorData;
  }
};
</script>

<template>
    <form class="space-y-6 font-[Roboto] text-red-500" @submit.prevent="submitForm">
        <div v-if="errorMessage" class="text-red-500">{{ errorMessage }}</div>
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
        <button
            class="w-full text-white bg-cl6 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center ">
            Log In
        </button>
    </form>
</template>
