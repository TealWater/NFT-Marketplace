<script>
	import { onDestroy, onMount } from 'svelte';
	import { CreateNFTEventFromSocket } from '$lib/util/parse.js';
	import EventRow from './event_row.svelte';
	export let data;
	const { opensea, collection } = data;

	let startTime;
	let elapsedTime = 0;
	let timer;
	onMount(() => {
		startTime = Date.parse('2024-05-19T19:47:52.777562+00:00');
		const interval = setInterval(() => {
			elapsedTime = Date.now() - startTime;
		}, 1000);
		const time = new Date(interval);
		console.log('**: ', time.toLocaleTimeString());

		return () => {
			clearInterval(interval);
		};
	});

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
		messages = [...messages, CreateNFTEventFromSocket(event.data)];
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
			{#each messages as { collection, event, timestamp, quantity, maker }}
				<EventRow {collection} {event} {timestamp} {quantity} {maker}></EventRow>
			{/each}

			{#each opensea as { order_type, asset, payment, quantity, maker, taker, event_timestamp }}
				<EventRow {order_type} {asset} {payment} {quantity} {maker} {taker} {event_timestamp}
				></EventRow>
			{/each}
		{:catch error}
			<p>{error.message}</p>
		{/await}
	</table>
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
