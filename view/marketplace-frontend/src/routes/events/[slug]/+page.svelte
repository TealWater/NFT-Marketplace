<script>
    export let data;
    const {opensea} = data
</script>

<section>
	<h2>Collection events for NFT</h2>
</section>
<section>
	<div class="title">Events</div>
	<hr />

    {#await opensea}
			<p>loading...</p>
		{:then opensea}
			{#each opensea as {order_type}}
				<div>{order_type}</div>
			{/each}
		{:catch error}
			<p>{error.message}</p>
		{/await}
	<!-- <hr> -->
</section>

<style>
	.title {
		text-align: center;
		font-size: x-large;
	}
</style>
