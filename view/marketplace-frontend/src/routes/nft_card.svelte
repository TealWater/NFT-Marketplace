<script>
	import { goto } from '$app/navigation';
	export let collection = '';
	export let image_url = '';

	function handleClickCollectionEvent() {
		// console.log('clicked');
		goto(`/events/${collection}`);
	}

	function handleClickViewCollection() {
		// console.log('clicked');
		goto(`/collection/${collection}`);
	}
</script>

<article>
	<section>
		<span>
			<img src={image_url} alt="solar eclipse" />
		</span>
	</section>

	<section class="info">
		<span class="stats">
			<!-- <h2>Hi mom!</h2> -->
		</span>
		<span class="stats">
			<h3 style="display: inline;">Collection: {collection}</h3>
		</span>

		<span class="stats">
			<!-- <h3>Floor Price:</h3> -->
		</span>

		<span class="stats">
			<h3>Chain: Ethereum</h3>
		</span>

		<span class="stats">
			<!-- <h3>Total supply:</h3> -->
		</span>

		<span id="button">
			<button on:click={handleClickCollectionEvent}>See Collection Events</button>
			<button on:click={handleClickViewCollection}>See Collection</button>
		</span>
	</section>
</article>

<style>
	article {
		display: flex;
		flex-direction: column;
		width: 300px;
		border: 1px solid #aaa;
		border-radius: 2px;
		border-color: blueviolet;
		box-shadow: 2px 2px 8px rgba(0, 0, 0, 0.1);
		padding: 1em;
		/* align-items: center; */
		min-height: 425px;
	}
	img {
		width: 300px;
		height: 300px;
		/* min-width: 300px; */
		/* min-height: 300px; */
		margin-left: -4px;
		border-color: orange;
		border-style: dotted;
	}
	article .info {
		align-items: left;
	}
	#button {
		display: flex;
		justify-content: center;
	}
	button{
		margin-left: 5px;
	}
</style>
