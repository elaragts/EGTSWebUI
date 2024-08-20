<script setup lang="ts">

</script>

<template>
  <div class="w-full m-2 flex flex-col">
    <div class="bg-cl5 rounded-xl m-1 h-full p-8 flex flex-col">
      <h1 class="text-4xl text-center border-b border-cl6 pb-2 font-bold mb-5">Settings</h1>
      <div class="text-xl flex flex-row">
        <div class="flex flex-col w-full md:w-1/3">
          <form class="flex flex-col mx-2 space-y-2 border-2 rounded-xl w-full">
            <div class="flex flex-col border-b border-cl6 px-2 mx-2 mb-4">
              <h1 class="text-2xl text-center font-bold mb-1">Change Username</h1>
              <label for="username" class="text-sm mb-1">Username:</label>
              <input
                  class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                  type="text"
                  placeholder="Username"
                  id="username"
                  required
                  minlength="1"
                  maxlength="100"
              >
            </div>
          </form>

          <form class="flex flex-col mx-2 space-y-2 border-2 rounded-xl w-full">
            <div class="flex flex-col border-b border-cl6 px-2 mx-2 mb-4">
              <h1 class="text-2xl text-center font-bold mb-1">Change Password</h1>
              <label for="oldPassword" class="text-sm mb-1">Old Password:</label>
              <input
                  class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                  type="password"
                  placeholder="Old Password"
                  id="oldPassword"
                  required
                  minlength="1"
                  maxlength="100"
              >
            </div>
            <div class="flex flex-col border-b border-cl6 px-2 mx-2 mb-4">
              <label for="newPassword" class="text-sm mb-1">New Password:</label>
              <input
                  class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                  type="password"
                  placeholder="New Password"
                  id="newPassword"
                  required
                  minlength="1"
                  maxlength="100"
              >
            </div>
            <div class="flex flex-col border-b border-cl6 px-2 mx-2 mb-4">
              <label for="confirmNewPassword" class="text-sm mb-1">Confirm New Password:</label>
              <input
                  class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                  type="password"
                  placeholder="Confirm New Password"
                  id="confirmNewPassword"
                  required
                  minlength="1"
                  maxlength="100"
              >
            </div>
          </form>
        </div>

        <form class="flex flex-col mx-2 space-y-2 border-2 rounded-xl w-full md:w-2/3">
          <div class="flex flex-col border-b border-cl6 px-2 mx-2 mb-4">
            <h1 class="text-2xl text-center font-bold mb-1">Access Codes</h1>
            <label for="newAccessCode" class="text-sm mb-1">Add Access Code:</label>
            <input
                class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                type="text"
                placeholder="New Access Code"
                id="newAccessCode"
                required
                minlength="1"
                maxlength="100"
            >
          </div>
        </form>
      </div>
    </div>
  </div>

</template>

<style scoped>

</style>