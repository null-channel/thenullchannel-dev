const endpoint = 'http://localhost:8080/ideas'
var app = new Vue({
  el: "#app",
  data () {
   return {
	   ideas: [] 
   }
  },
  mounted(){
	axios.get(endpoint,{crossDomain:true})
	.then(response =>  (this.ideas = response.data))
	.catch((err) => console.log(err));
		
}
});

