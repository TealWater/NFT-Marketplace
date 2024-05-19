<script>
	import { onDestroy } from 'svelte';
	import {CreateNFTEvent} from '$lib/util/parse.js';
	export let data;
	const { opensea, collection } = data;

	/**
	 * @type {any[]}
	 */
	let messages = [];

	// Create WebSocket connection.
	const socket = new WebSocket(`ws://localhost:8080/opensea`);

	// Connection opened
	socket.addEventListener('open', (event) => {
		console.log('socket is open');
		socket.send(`${collection.toString().trim()}`);
	});

	// Listen for messages
	socket.addEventListener('message', (event) => {
		// messages = [...messages, JSON.parse(event.data)];
		messages = [...messages, CreateNFTEvent(event.data)];
		messages.reverse();
	});

	onDestroy(() => socket.close());
</script>

<section>
	<h2>Collection events for NFT</h2>
</section>
<section>
	<div class="title">Events</div>
	<hr />

	<table class="table">
		<thead>
			<tr>
				<th class="event"></th>
				<th class="item">Item</th>
				<th class="price">Price</th>
				<th class="rarity">Rarity</th>
				<th class="quantity">Quantity</th>
				<th class="from">From</th>
				<th class="to">To</th>
				<th class="time">Time</th>
			</tr>
		</thead>

		{#await opensea}
			<p>loading...</p>
		{:then opensea}
			{#each messages as {collection, event, chain, timestamp, quantity}}
				<!-- <div>{event}</div> -->
				<tr>
					<td class="event">{event}</td>
					<td class="item">{collection}</td>
					<td class="price">#</td>
					<td class="rarity">#</td>
					<td class="quantity">{quantity}</td>
					<td class="from">#</td>
					<td class="to">#</td>
					<td class="time">{timestamp}</td>
				</tr>
			{/each}

			{#each opensea as { order_type, asset, payment, quantity, maker, taker, event_timestamp }}
				<tr>
					<td class="event">{order_type}</td>
					<td class="item">{asset.collection}</td>
					<td class="price">{payment.quantity}</td>
					<td class="rarity">#</td>
					<td class="quantity">{quantity}</td>
					<td class="from">{maker}</td>
					<td class="to">{taker}</td>
					<td class="time">{event_timestamp}</td>
				</tr>
			{/each}
		{:catch error}
			<p>{error.message}</p>
		{/await}
	</table>

	<!-- {#await opensea}
		<p>loading...</p>
	{:then opensea} -->
	<!-- {#each messages as msg}
			<div>{msg}</div>
		{/each} -->

	<!-- {#each opensea as { order_type }}
			<div>{order_type}</div>
		{/each}
	{:catch error}
		<p>{error.message}</p>
	{/await} -->
	<!-- <hr> -->
</section>

<style>
	.title {
		text-align: center;
		font-size: x-large;
	}
	.table {
		width: 100%;
	}

	.item {
		width: 15%;
		text-align: left;
	}
	.price {
		width: 5%;
		text-align: right;
	}
	.rarity {
		width: 5%;
		text-align: right;
		color: pink;
	}

	.quantity {
		width: 5%;
		text-align: right;
		color: red;
	}

	.from {
		width: 30%;
		text-align: right;
		color: orange;
	}

	.to {
		width: 20%;
		text-align: right;
		color: thistle;
	}

	.time {
		width: 10%;
		text-align: right;
		color: purple;
	}
</style>
