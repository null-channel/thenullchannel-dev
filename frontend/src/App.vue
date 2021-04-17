<template>
  <div id="app">
     <p class="lead-text">What would you like to see next ?</p> 
    <Ideas/>
  </div>
</template>

<script>
import Ideas from "./components/Ideas";
export default {
  name: "App",
  components: {
    Ideas,
  },
};
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
