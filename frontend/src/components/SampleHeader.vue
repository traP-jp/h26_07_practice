<script setup lang="ts">
defineProps<{
  msg: string
}>()
</script>

<template>
    <main>
        <h1 class="atama">{{ msg }}</h1>
    </main>
</template>

<style>
.atama {
    display: block;
    height: 10svh;
    background-color: navy;
    color: var(--vt-c-white);
    font-size: 40pt;
    text-align: center;
    position: sticky;
}
</style>