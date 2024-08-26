<script setup lang="ts">
import {useAuthStore} from "@/store/authStore";
import {useToast} from "vue-toastification";
import {ref, watch} from "vue";

const authStore = useAuthStore();
const toast = useToast();

const username = ref(authStore.username);
watch(() => authStore.initialized, (value) => username.value = authStore.username);

const currentPassword = ref('');
const newPassword = ref('');
const confirmNewPassword = ref('');

// loading access codes with a get request
const accessCodes = ref<string[]>([]);
const fetchAccessCodes = async () => {
  try {
    const response = await fetch(`/api/user/${authStore.baid}/access_codes`);
    if (response.ok) {
      accessCodes.value = await response.json();
    } else {
      toast.error("Failed to fetch access codes");
      console.error('Unexpected access code data response status:', response.status);
    }
  } catch (error) {
    toast.error("Failed to fetch access codes");
    console.error('Failed to fetch access codes:', error);
  }
}
// calling the above function when baid loads
if (authStore.initialized) {
  fetchAccessCodes()
} else {
  watch(() => authStore.initialized, (value) => {
    fetchAccessCodes();
  });
}

const newAccessCode = ref('');


async function changeUsername(): Promise<boolean> {
  const data = {
    username: username.value
  }

  const response = await fetch(`/auth/user/${authStore.baid}/username`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });


  if (response.ok) {
    toast.success('Username Changed Successfully');
    authStore.username = username.value;
    return true;
  } else {
    toast.error(`Failed to change username: ${await response.text()}`);
    return false;
  }
}


async function changePassword(): Promise<boolean> {
  const data = {
    currentPassword: currentPassword.value,
    newPassword: newPassword.value
  }

  if (newPassword.value !== confirmNewPassword.value) {
    toast.error('Passwords do not match');
    return false;
  }

  const response = await fetch(`/auth/user/${authStore.baid}/password`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });


  if (response.ok) {
    toast.success('Password Changed Successfully');
    currentPassword.value = '';
    newPassword.value = '';
    confirmNewPassword.value = '';
    return true;
  } else {
    toast.error(`Failed to change password: ${await response.text()}`);
    return false;
  }
}


async function addAccessCode(): Promise<boolean> {
  const data = {
    accessCode: newAccessCode.value
  }

  const response = await fetch(`/api/user/${authStore.baid}/access_codes`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });


  if (response.ok) {
    toast.success('Access Code Added Successfully');
    accessCodes.value.push(newAccessCode.value);
    newAccessCode.value = '';
    return true;
  } else {
    toast.error(`Failed to add access code: ${await response.text()}`);
    return false;
  }
}

async function deleteAccessCode(accessCode: string): Promise<boolean> {
  const data = {
    accessCode: accessCode
  }

  const response = await fetch(`/api/user/${authStore.baid}/access_codes`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });


  if (response.ok) {
    toast.success('Access Code Deleted Successfully');
    accessCodes.value = accessCodes.value.filter(currentCode => currentCode != accessCode)
    return true;
  } else {
    toast.error(`Failed to delete access code: ${await response.text()}`);
    return false;
  }
}

</script>

<template>
  <div class="w-full m-2 flex flex-col">
    <div class="bg-cl5 rounded-xl m-1 h-full p-8 flex flex-col">
      <h1 class="text-4xl text-center border-b border-cl6 pb-2 font-bold mb-5">Settings</h1>
      <div class="text-xl flex flex-row space-x-4">
        <div class="flex flex-col space-y-2 w-full md:w-1/3">
          <form @submit.prevent="changeUsername" class="flex flex-col mx-2 space-y-2 border-2 pb-4 rounded-xl w-full">
            <div class="flex flex-col border-b border-cl6 pt-2 px-2 mx-2 mb-2">
              <h1 class="text-2xl text-center font-bold mb-2">Change Username</h1>
              <label for="username" class="text-sm mb-1">Username:</label>
              <input
                  class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                  type="text"
                  placeholder="Username"
                  id="username"
                  required
                  minlength="1"
                  maxlength="20"
                  v-model="username"
                  @keydown.enter.prevent
              >
            </div>

            <button
                    class="w-1/2 self-center p-2 text-base font-bold border-cl6 bg-cl6 text-cl5 rounded-lg border-2 shadow hover:bg-opacity-80 hover:transition-all"
                    type="submit"
            >
              Change Username
            </button>
          </form>

          <form @submit.prevent="changePassword" class="flex flex-col mx-2 space-y-4 border-2 rounded-xl pb-4 w-full">
            <div class="flex flex-col border-b border-cl6 pt-2 px-2 mx-2">
              <h1 class="text-2xl text-center font-bold mb-2">Change Password</h1>
              <label for="currentPassword" class="text-sm">Current Password:</label>
              <input
                  class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                  type="password"
                  placeholder="Current Password"
                  id="currentPassword"
                  required
                  minlength="1"
                  v-model="currentPassword"
                  @keydown.enter.prevent
              >
            </div>
            <div class="flex flex-col border-b border-cl6 px-2 mx-2">
              <label for="newPassword" class="text-sm">New Password:</label>
              <input
                  class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                  type="password"
                  placeholder="New Password"
                  id="newPassword"
                  required
                  minlength="8"
                  maxlength="100"
                  v-model="newPassword"
                  @keydown.enter.prevent
              >
            </div>
            <div class="flex flex-col border-b border-cl6 px-2 mx-2">
              <label for="confirmNewPassword" class="text-sm mb-1">Confirm New Password:</label>
              <input
                  class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                  type="password"
                  placeholder="Confirm New Password"
                  id="confirmNewPassword"
                  required
                  minlength="8"
                  maxlength="100"
                  v-model="confirmNewPassword"
                  @keydown.enter.prevent
              >
            </div>

            <button
                class="w-1/2 self-center p-2 text-base font-bold border-cl6 bg-cl6 text-cl5 rounded-lg border-2 shadow hover:bg-opacity-80 hover:transition-all"
                type="submit"
            >
              Change Password
            </button>
          </form>
        </div>

        <div class="flex flex-col mx-2 space-y-6 border-2 rounded-xl w-full md:w-2/3">
          <form @submit.prevent="addAccessCode" class="flex flex-col space-y-2 w-full">
            <div class="flex flex-col border-b border-cl6 pt-2 px-2 mx-2 mb-2">
              <h1 class="text-2xl text-center font-bold mb-2">Manage Access Codes</h1>
              <label for="newAccessCode" class="text-sm mb-1">Add Access Code:</label>
              <input
                  class="appearance-none bg-transparent border-none text-gray-700 focus:outline-none w-full h-10"
                  type="text"
                  placeholder="New Access Code"
                  id="newAccessCode"
                  required
                  minlength="1"
                  maxlength="100"
                  v-model="newAccessCode"
                  @keydown.enter.prevent
              >
            </div>
            <button
                class="w-1/4 self-center p-2 text-base font-bold border-cl6 bg-cl6 text-cl5 rounded-lg border-2 shadow hover:bg-opacity-80 hover:transition-all"
                type="submit"
            >
              Add
            </button>
          </form>
          <v-data-table
              class="h-full"
              :items="accessCodes.map((code, index) => ({ id: index + 1, code }))"
              :headers="[
                      { title: 'Active Codes', align: 'start', value: 'code', sortable: false },
                      { title: '', align: 'end', value: 'actions', sortable: false }
                      ]"
              item-key="id"
              item-value="code"
              no-data-text="No access codes found"
              items-per-page="5"
              density="comfortable"
              :items-per-page-options="[{value: 5, title: '5'}]"
          >
            <template v-slot:item.actions="{ item }">
              <v-icon
                  size="small"
                  icon="mdi-delete"
                  @click="deleteAccessCode(item.code)"
              >
              </v-icon>
            </template>
          </v-data-table>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>