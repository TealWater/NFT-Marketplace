<script>
	/** @type {import('./$types').PageData} */
	import NftListing from './nft_listing.svelte';
	export let data;
	const { opensea } = data;
</script>

<section>
	{#await opensea}
		<p>loading...</p>
	{:then opensea}
		{#each opensea as { collection, name, image_url, price, currency, identifier }}
			<div>
				<NftListing {name} {identifier} {image_url} {price} {currency}></NftListing>
			</div>
		{/each}
	{:catch error}
		<p>{error.message}</p>
	{/await}
</section>

<style>
	section {
		display: flex;
		flex-wrap: wrap;
		align-content: space-between;
		margin-left: 20px;
		padding-top: 60px;
	}
	div {
		margin: 10px;
	}
</style>
