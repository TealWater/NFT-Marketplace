<script>
	import { page } from '$app/stores';
    const goofy  = 'https://media1.tenor.com/m/LRGmSJAf_dYAAAAC/goofy-laugh.gif';
</script>

<!-- <h1>{$page.status}</h1> -->
<div>
	<img src={goofy}/>
	<script type="text/javascript" async src="https://tenor.com/embed.js"></script>
	{#if $page.status == 404}
		<h1>{$page.status}</h1>
		<h1>Hmm, you went somewhere you're not suppose to be, therefore you broke it.</h1>
		<a href="/">Go Home</a>
	{:else if $page.status == 500}
		<h1>My bad, it appears my elves and worker bees have dropped the ball!</h1>
		<h2>We're on it. In the mean time, why not <a href="/">Go Home?</a></h2>
	{/if}
</div>

<style>
	div {
		padding-top: 60px;
		text-align: center;
	}
	a {
		font-size: x-large;
	}
</style>
