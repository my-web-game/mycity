-- Active: 1753917762459@@127.0.0.1@3306
<script setup>
import { init, uninit } from '@/scripts/main';
import { onMounted, onUnmounted } from 'vue';

onMounted(() => {
    init()
});

onUnmounted(() => {
    uninit()
})
</script>

<template>
  <main></main>
</template>

<style>
html {
  margin: 0px 0px;
  padding: 0px 0px;
  overflow: hidden;
}
body {
  margin: 0px 0px;
  padding: 0px 0px;
  overflow: hidden;
}
#app {
  margin: 0px 0px;
  padding: 0px 0px;
  overflow: hidden;
}
</style>
