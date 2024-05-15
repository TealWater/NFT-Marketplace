<script>
	export let data;
	const { opensea } = data;

	let messages = [];

	// Create WebSocket connection.
	const socket = new WebSocket('ws://localhost:8080/opensea');

	// Connection opened
	socket.addEventListener('open', (event) => {
		console.log('socket is open');
	});

	// Listen for messages
	socket.addEventListener('message', (event) => {
		messages = [...messages, event.data];
		messages.reverse();
	});
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
		{#each messages as msg}
			<div>{msg}</div>
		{/each}

		{#each opensea as { order_type }}
			<div>{order_type}</div>
		{/each}
	{:catch error}
		<p>{error.message}</p>
	{/await}
	<!-- <hr> -->

	<!-- {#each messages as msg}
		<div>{msg}</div>
	{/each} -->
</section>

<style>
	.title {
		text-align: center;
		font-size: x-large;
	}
</style>
