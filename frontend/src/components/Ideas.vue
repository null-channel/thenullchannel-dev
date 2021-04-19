<template>
  <div>
    <div :key="idea.id" v-for="idea in ideas" class="idea">
      <p class="description">{{ idea.description }}</p>
      <p class="votes">{{ idea.votes }}</p>
      <input type="submit" v-on:click="Vote(idea.id,idea)" class="vote-btn" value="vote" />
    </div>
  </div>
</template>


<script>
import axios from "axios";

export default {
  name: "Ideas",
  data() {
    return {
      ideas: [],
    };
  },
  mounted() {
    fetch(process.env.VUE_APP_API_URL + "/ideas")
      .then((response) => response.json())
      .then((data) => this.ideas = data)
      .catch((err) => console.log(err.message));
  },
  methods: {
    Vote(id,idea) {
      axios
        .post(process.env.VUE_APP_API_URL + "/vote", {
          id: id,
          description:null,
          votes:null,
        })
        .then(
          (response) => {
           console.log(response.status)
           idea.votes +=1
          },
          (error) => {
            console.log(error);
          }
        );
    },
  },
};
</script>


<style scoped>
.idea {
  margin: 0 auto;
  background-color: #35556be7;
  width: 48rem;
  height: 20rem;
  margin-top: 3rem;
  border-radius: 1.5rem;
  box-shadow: 10px 5px 5px #35556b;
  transition: all 0.5s ease-in-out;
}
.idea .description {
  margin-top: 30px;
  font-size: 2.5rem;
  font-family: "Source Sans Pro", sans-serif;
  color: #d7e4e2;
  text-align: center;
}

.idea .votes {
  font-size: 2rem;
  margin-top: 4rem;
  margin-left: 22rem;
  color: #d7e4e2;
  font-family: "Source Sans Pro", sans-serif;
}
.idea .vote-btn {
  font-size: 1.3rem;
  color: #d7e4e2;
  margin-left: 16rem;
  margin-top: 5rem;
  background: #35556b;
  border-radius: 4px;
  border-color: #35556b;
  width: 12rem;
  height: 3rem;
}
</style>