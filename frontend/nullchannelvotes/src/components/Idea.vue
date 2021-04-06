<template>  
    <div v-for="idea in ideas :key" class="idea">
      <p class="description">{{ idea.description }}</p>
      <p class="votes">{{ idea.votes }}</p>
      <input type="button" value="vote" v-on:click="vote" class="vote-btn" />
    </div>

</template>

<script>
export default {
  name: "Idea",
  data() {
    return {
      ideas: [],
    };
  },
  created() {
    this.fetchideas();
  },
  methods: {
    fetchideas() {
      axios
        .get("/ideas", { crossDomain: true })
        .then((response) => (this.ideas = response.data))
        .catch((err) => console.log(err));
    },
  },
};
</script>