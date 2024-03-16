<script setup lang="ts">
import { useAuthStore } from "@/store/authStore";
import router from "@/router";
const authStore = useAuthStore();
authStore.logout();
import {useToast} from 'vue-toastification';
const toast = useToast();
toast.success('Logged out successfully');
router.push({name: 'home'});

</script>
