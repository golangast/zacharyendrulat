<script>
	import Apis, {actualapi, getResult} from "./getapi.svelte";
	let actualtt = getResult("http://1299991b.ngrok.io/get");
	let ff = [];
	var promise = Promise.resolve(actualtt).then(function(val) { 
			console.log(val); 
		ff = val;
	});  
</script>

{#each ff as {Dates, Title, Slug, Html}, i }
 <ul>
 <li>{i}: {Dates} {Title} {Slug} {Html}</li>
 </ul>
 {/each}
