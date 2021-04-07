<template>
  <div class="idea">
   <div :key="idea.id" v-for="idea in ideas" >
        <p class="decription">{{ idea.description }}</p>
        <input type="hidden" name="id" id="id"  :ideaid="idea.id"/>
        <p class="votes">{{ votes }}</p>
        <input type="submit" v-on:click="Vote()" class="vote" value="vote" />
  </div>
</div>
</template>


<script>

export default {
  name: "Ideas",
  props: {
    ideas: Array
  },
  methods: {
    Votes() {
      const url = process.env.VUE_APP_API_URL + "/vote";
      fetch(url, {
        method: "POST",
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
          "Content-Type": "application/json",
        },
        redirect: "follow",
        referrerPolicy: "same-origin",
        body: JSON.stringify({
        id: this.id,
        description: this.description,
        votes: this.votes})
      }).then(data => console.log(data))
      .catch(err => {console.log(err)})
    },
  },
};
</script>


<style scoped>
.idea{
 margin: 0 auto;
    background-color:#35556be7;
    width:48rem;
    height:20rem;
    margin-top: 3rem;
    border-radius:1.5rem;
    box-shadow: 10px 5px 5px #35556b;
    transition: all .5s ease-in-out;
}
.idea .description{
    margin-top:30px;
    font-size:90px;
    font-family: 'Source Sans Pro', sans-serif;
    color:#d7e4e2;
    text-align: center;
}

vote{
    font-size:1.3rem;
    color: #d7e4e2;
    margin-left: 16rem;
    margin-top:5rem;
    background:#35556b;
    border-radius: 4px;
    border-color:#35556b;
    width:12rem;
    height:3rem;
}

</style>