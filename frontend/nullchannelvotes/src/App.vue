<template>
  <div id="app">
     <p class="lead-text">What would you like to see next ?</p> 
    <Ideas :ideas="ideas" />
  </div>
</template>

<script>
import Ideas from "./components/Ideas";
import axios from "axios";
require("dotenv").config();

export default {
  name: "App",
  components: {
    Ideas,
  },
  data() {
    return {
      ideas: [],
    };
  },
  methods: {
    getideas() {
      var data 
      axios
        .get(process.env.VUE_APP_API_URL + "/ideas", { crossDomain: true })
        .then((response) => (data = response.json))
        .catch((err) => console.log(err));
        return data
    },
  },
  created() {
    this.ideas = this.getideas();
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
