<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
      <p class="lead-text">What would you like to see next ?</p>
    <Ideas/>
  </div>
</template>

<script>
import Idea from './components/Idea.vue'

export default {
  name: 'App',
  components: {
   Idea
  }
}
</script>

<style>
#app .lead-text{
    color:#8898a4;
    font-size:5rem;
    margin-top:7rem;
    text-align:center;
    font-family: 'Source Sans Pro', sans-serif;

}
</style>
