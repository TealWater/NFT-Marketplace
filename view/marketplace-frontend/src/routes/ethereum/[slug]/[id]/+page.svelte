<script>
	/* @type {import('./$types').PageData} */
	import { page } from '$app/stores';
	import onboard from '$lib/stores/embeddedWalletStore';
	import { wallet_state } from '$lib/stores/store.js';
	import { onMount } from 'svelte';
	export let data;
	const { nft } = data;
	// console.log('nft ', nft);
	// console.log('params ', $page.params);
	const wallets = onboard.state.select('wallets');
	let hasProvider = false;
 	//connectedAccount;

	function buyNFT() {
		// console.log("**:",hasProvider);
		if (hasProvider) {
			alert(
				'ROFL! Hey bud, you need to connect your crypto wallet (metamask, phantom, etc...) to purchase this NFT.\n\nRemember... no shoes, no shirt, no service.\n\n\nIn this case its no wallet no service.'
			);
		} else {
			alert(
				'Hmm... it seems the elves have not fully implemented the buying functionality yet. Check back later!'
			);
		}
	}

	onMount(() => {
		const connectedAccount = $wallets?.[0]?.accounts?.[0];
		if(connectedAccount.address != null ){
			hasProvider = true;
		}
	})
</script>

<section>
	{#await nft}
		<p>loading...</p>
	{:then nft}
		<div class="img">
			<img src={nft.image_url} alt="" />
			<!-- <p>contract: {$page.params.slug}</p>
			<p>identifier: {$page.params.id}</p> -->
		</div>
		<div class="offer">
			<h2>{nft.collection} #{nft.identifier}</h2>
			<button on:click={buyNFT}>buy</button>
		</div>
		<div class="traits">
			<h2>Traits</h2>
			<table>
				<!-- <tr>
					<thead>Trait</thead>
					<thead>Floor</thead>
				</tr> -->
				{#each nft.traits as { trait_type, value }}
					<tr>
						{trait_type} : {value}
					</tr>
				{/each}
				<!-- <p>{nft.traits[0].trait_type}</p> -->
			</table>
		</div>
	{/await}
</section>
<section>
	<table>
		<tr>
			<thead>
				<th>Item</th>
				<th>Price</th>
				<th>From</th>
				<th>Order Type</th>
			</thead>
		</tr>
		<tr>
			<td>hi mom!</td>
		</tr>
	</table>
</section>

<style>
	section {
		display: flex;
		align-content: center;
		padding-top: 60px;
		margin: 10px;
		width: 100%;
	}
	.img {
		border-color: blue;
		border-style: dashed;
		max-width: 400px;
		max-height: 400px;
	}

	img {
		display: block;
		max-width: 400px;
		max-height: 400px;

		min-width: 200px;
		min-height: 200px;
	}

	.offer {
		border-color: red;
		border-style: dashed;
		font-weight: bold;
		flex-direction: column;
		justify-content: end;
	}
	.offer h2,
	button {
		display: block;
	}

	.traits {
		border-color: black;
		border-style: groove;
	}
	.traits tr {
		font-size: large;
	}
	.traits tr:nth-of-type(odd) {
		background-color: #eee;
	}
</style>
