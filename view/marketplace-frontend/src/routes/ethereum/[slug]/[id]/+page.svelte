<script>
	/* @type {import('./$types').PageData} */
	import { page } from '$app/stores';
	import { MetaMaskStore } from '$lib/stores/metamaskStore';
	export let data;
	const { nft } = data;
	const { walletState, isMetaMaskPresent, connect, loaded, init } = MetaMaskStore
	// const image = nft[image_url];
	// console.log("image: ", image);
	console.log('nft ', nft);
	console.log('params ', $page.params);

	function buyNFT(){
		console.log("loaded22:", loaded);
		// console.log("wallet state:", $walletState);

		if(!loaded){
			alert("ROFL! Hey bud, you need to connect your crypto wallet (metamask, phantom, etc...) to purchase this NFT.\n\nRemember... no shoes, no shirt, no service.\n\n\nIn this case its no wallet no service.");
		}else{
			alert("Hmm... it seems the elves have not fully implemented the buying functionality yet. Check back later!");
		}
	}

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
			<h2>{nft.collection} #{nft.identifier} </h2>
			<button on:click={buyNFT}>buy</button>
		</div>
	{/await}
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
		width: 30%;
		height: 30%;
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
		justify-content:end;
	}
	.offer h2, button{
		display: block;
	}
</style>
