<script>
	import { collection_count } from '$lib/stores/store';
	import { error } from '@sveltejs/kit';
	import NftCard from './nft_card.svelte';
	export let data;
	const { opensea } = data;
	let collection = 'persona';
</script>

<section>
	{#await opensea}
		<p>loading...</p>
	{:then opensea}
		{#each opensea as { collection, image_url }}
			<div>
				<NftCard {collection} {image_url}></NftCard>
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
