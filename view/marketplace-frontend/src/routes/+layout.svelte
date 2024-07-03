<script>
	import { source } from 'sveltekit-sse';
	import { MetaMaskStore } from '$lib/stores/metamaskStore';
	import { PUBLIC_TRUSTED_URL } from '$env/static/public';
	import { onMount } from 'svelte';
	const eventSourceGas = source(`${PUBLIC_TRUSTED_URL}/stream`).select('message');

	const { walletState, isMetaMaskPresent, connect, loaded, init } = MetaMaskStore();
	onMount(() => {
		init();
	});
	console.log(walletState);
</script>

<nav>
	<div class="container">
		<div class="logo">Logo</div>
		<div class="navbar">
			<a href="/">home</a>
			<!-- svelte-ignore a11y-invalid-attribute -->
			<a href="#">about</a>
			<!-- svelte-ignore a11y-invalid-attribute -->
			<button on:click={connect}>connect wallet</button>
		</div>
	</div>
</nav>

<slot />

<footer>
	<section>
		Gas: {$eventSourceGas}
	</section>
</footer>

<style>
	.container {
		position: fixed;
		top: 0;
		display: flex;
		height: 60px;
		width: 100%;
		/* border-style: dotted; */
		/* border-color: blueviolet; */
		justify-content: space-between;
		align-items: center;
		background-color: wheat;
	}
	.logo {
		font-size: x-large;
		font-weight: bold;
		/* border-style: dotted; */
		/* border-color: palevioletred; */
		height: 50%;
		margin-left: 5px;
	}

	.navbar a,
	button {
		font-size: 18px;
		font-weight: bold;
		height: 100%;
		justify-content: space-between;
		margin-right: 30px;
	}
	.navbar a,
	button:hover {
		color: brown;
	}

	.navbar button {
		background: none;
		border: none;
		padding: 0 !important;
		/*optional*/
		font-family: arial, sans-serif;
		/*input has OS specific font-family*/
		color: #069;
		text-decoration: underline;
		cursor: pointer;
	}

	footer {
		position: sticky;
		bottom: 0;
		background-color: wheat;
		width: 100%;
	}

	footer section {
		text-align: center;
	}
</style>
