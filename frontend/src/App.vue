<script setup lang="ts">
// import { RouterLink, RouterView } from 'vue-router'
import SampleHeader from './components/SampleHeader.vue'
//import { ref } from 'vue'
const URL = ''
let cont = ''
const f = async () => {
  const response = await fetch(URL)
  cont = String(response)
}
</script>

<template>
  <header>
    <SampleHeader msg="Hello"></SampleHeader>
  </header>
  <input type="button" value="fetch" @click="f" />
  <div>{{ cont }}</div>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
  position: sticky;
  top: 0;
  z-index: 1000;
}
div {
  margin: 10px;
}

/*.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}*/
</style>
