<script setup lang="ts">
import { computed } from "vue"
const props = defineProps<{id: string}>();
const url = computed(() => `https://q.trap.jp/api/v3/public/icon/${props.id}`)
</script>

<template>
    <img :src="url">
</template>

<style>
img {
    width: 125px;
    height: 125px;
}
</style>