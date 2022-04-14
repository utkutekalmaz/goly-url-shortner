<script setup>
import { RouterLink, RouterView } from 'vue-router'
import './index.css'
</script>

<template>
  <!-- <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header> -->
<div id="app-main" class="w-screen h-screen" >
  <RouterView />

</div>
</template>

<style>

</style>
