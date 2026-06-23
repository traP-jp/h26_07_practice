<script steup lang="ts">
import { computed, ref } from 'vue'
const time = ref(new Date())
const h = computed(() => time.value.getHours())
const m = computed(() => time.value.getMinutes())
const s = computed(() => time.value.getSeconds())
</script>

<template>
  <p>{{ h }} 時 {{ m }} 分 {{ s }} 秒</p>
</template>
