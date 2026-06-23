<script setup lang="ts">
defineProps<{
    keta: string
}>();
</script>

<template>
    <p>Lorem Ipsum Dolor {{ keta }}0</p>
    <p>Lorem Ipsum Dolor {{ keta }}1</p>
    <p>Lorem Ipsum Dolor {{ keta }}2</p>
    <p>Lorem Ipsum Dolor {{ keta }}3</p>
    <p>Lorem Ipsum Dolor {{ keta }}4</p>
    <p>Lorem Ipsum Dolor {{ keta }}5</p>
    <p>Lorem Ipsum Dolor {{ keta }}6</p>
    <p>Lorem Ipsum Dolor {{ keta }}7</p>
    <p>Lorem Ipsum Dolor {{ keta }}8</p>
    <p>Lorem Ipsum Dolor {{ keta }}9</p>
</template>

<style>
p {
    margin-left: 8px;
    margin-right: 8px;
}
</style>