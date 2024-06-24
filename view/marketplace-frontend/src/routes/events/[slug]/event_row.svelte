<script>
	import { onMount } from 'svelte';
	export let event = '#';
	export let event_type = '#';
	export let collection = '#';
	export let rarity = '#';
	export let price = '#';
	export let quantity = '#';
	export let maker = '#';
	export let taker = '#';
	export let timestamp = '#';

	// for values from API call
	export let order_type = '#';
	export let asset = {};
	export let nft = {};
	export let payment = {};
	export let event_timestamp = '#';

	if (event == '#') {
		event = order_type;
	}
	
	if (collection == '#') {
		// @ts-ignore
		if(asset.name != ''){
			// @ts-ignore
			collection = asset.name;
		}else{
			// @ts-ignore
			collection = nft.name;
		}
	}

	if (price == '#') {
		// @ts-ignore
		price = payment.quantity;

		if(price == '' && event_type == 'cancel'){
			price = event + ' canceled';
		}
	}

	if (timestamp == '#') {
		timestamp = event_timestamp;
	}

	let parsedDate = new Date(Number.parseInt(timestamp) * 1000);
	// console.log('time ;', parsedDate);
	let elapsedTime = '';
	onMount(() => {
		const interval = setInterval(() => {
			const currentDate = new Date();
			const timeDiff = currentDate.getTime() - parsedDate.getTime();
			const seconds = Math.floor(timeDiff / 1000);
			const minutes = Math.floor(seconds / 60);
			const hours = Math.floor(minutes / 60);

			elapsedTime = `${hours} hours, ${minutes % 60} minutes, ${seconds % 60} seconds`;
		}, 1000);

		return () => {
			clearInterval(interval);
		};
	});
</script>

<tr>
	<td class="event">{event}</td>
	<td class="item">{collection}</td>
	<td class="price">{price}</td>
	<td class="rarity">#</td>
	<td class="quantity">{quantity}</td>
	<td class="from">{maker}</td>
	<td class="to">{taker}</td>
	<td class="time">{elapsedTime}</td>
</tr>

<style>
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
