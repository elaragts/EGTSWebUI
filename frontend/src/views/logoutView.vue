<script setup lang="ts">
import { useAuthStore } from "@/store/authStore";
import router from "@/router";
const authStore = useAuthStore();
authStore.logout();
router.push({name: 'home'});
</script>
