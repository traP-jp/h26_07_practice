<script setup lang="ts">
import { ref } from 'vue'
const Id = ref<string>('HokubuRailway')
const changeImage = () => {
  Id.value = 'HokubuRailway'
}
</script>

<template>
  <div>
    <input type="text" v-model="Id" />
    <input type="button" value="Go" @click="changeImage" />
  </div>
  <div>
    <UserIcon :id="Id"></UserIcon>
  </div>
</template>
