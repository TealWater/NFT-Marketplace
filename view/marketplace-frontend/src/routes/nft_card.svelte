<script>
	import { goto } from '$app/navigation';
	export let collection = '';
	let src =
		'https://images.pexels.com/photos/23279639/pexels-photo-23279639/free-photo-of-a-solar-eclipse-is-seen-in-the-sky.jpeg';

	let nft = 'https://i.seadn.io/s/raw/files/c8dba01a030a3824ace3cc20521785fa.png';

	function handleClick() {
		// console.log('clicked');
		goto(`/events/${collection}`);
	}
</script>

<article>
	<section>
		<span>
			<img src={nft} alt="solar eclipse" />
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
			<button on:click={handleClick}>See Collection Events</button>
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
	}
	img {
		max-width: 300px;
		max-height: 300px;
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
</style>
