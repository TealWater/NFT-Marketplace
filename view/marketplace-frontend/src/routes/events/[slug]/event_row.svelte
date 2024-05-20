<script>

    export let event = '#';
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
    export let payment = {};
    export let event_timestamp = '#';

    if(event == '#'){
        event = order_type;
    }
    if(collection == '#'){
        // @ts-ignore
        collection = asset.collection;
    }

    if(price == '#'){
        // @ts-ignore
        price = payment.quantity;
    }

    if(timestamp = '#'){
        timestamp = event_timestamp;
    }

</script>

<tr>
    <td class="event">{event}</td>
    <td class="item">{collection}</td>
    <td class="price">{price}</td>
    <td class="rarity">#</td>
    <td class="quantity">{quantity}</td>
    <td class="from">{maker}</td>
    <td class="to">{taker}</td>
    <td class="time">{timestamp}</td>
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