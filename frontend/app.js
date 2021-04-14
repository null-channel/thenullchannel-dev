
const endpoint = 'http://localhost:8080/'
var app = new Vue({
  el: "#app",
  data () {
   return {
	   ideas: [] 
   }
  },
  mounted(){
	axios.get(endpoint+"/ideas",{crossDomain:true})
	.then(response =>  (this.ideas = response.data))
	.catch((err) => console.log(err));		
},
methods: {
  vote(){
    axios.post(endpoint+"/vote",{
      id: this.ideas.id,
      description: this.ideas.description,
      votes: this.ideas.votes
    })
    .then(response => (console.log(response.data)))
    .catch((err) => console.log(err));	
  }
}
});

