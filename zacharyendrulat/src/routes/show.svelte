<script>
	import Apis, {actualapi, getResult} from "./getapi.svelte";
	let actualtt = getResult("http://localhost:8081/get");
	let ff = [];
	var promise = Promise.resolve(actualtt).then(function(val) { 
			console.log(val); 
		ff = val;
	});  
</script>

{#each ff as {ID, Title, Slug}, i }
 <ul>
 <li>{i}: {ID[0]} {Title[0]} {Slug[0]}</li>
 </ul>

 {/each}

