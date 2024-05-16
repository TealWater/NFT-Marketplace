<script>
	import NftCard from './nft_card.svelte';
	let collection = 'persona';
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<NftCard {collection}></NftCard>
