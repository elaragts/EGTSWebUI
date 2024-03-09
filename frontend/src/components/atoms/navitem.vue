

<template>
  <router-link :to="to" class="block py-2 px-3 text-cl5 rounded font-[Roboto] md:text-2xl sm:text-xl md:bg-transparent md:p-0 hover:text-cl6">
    {{ title }}
  </router-link>
</template>

<script lang="ts">
export default {
  props: {
    to: {
      type: String,
      required: true
    },
    title: {
      type: String,
      required: true
    }
  }
}
</script>
