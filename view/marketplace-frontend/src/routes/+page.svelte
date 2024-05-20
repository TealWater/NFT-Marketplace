<script>
	import { collection_count } from '$lib/stores/store';
	import NftCard from './nft_card.svelte';
	export let data;
	const { opensea } = data;
	let collection = 'persona';
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<section>
	<!-- {#await opensea}
		<p>loading...</p>
	{:then opensea}
		{#each opensea as { collection }}
		<div>
			<NftCard {collection}></NftCard>
		</div>
		{/each}
	{/await} -->

	<NftCard {collection}></NftCard>
</section>

<style>
	section{
		display: flex;
		flex-wrap: wrap;
		align-content: space-between;
		margin-left: 20px;
	}
	div{
		margin: 10px;
	}
</style>
