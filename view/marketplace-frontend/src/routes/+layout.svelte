<script>
	import { source } from 'sveltekit-sse';
	import { PUBLIC_TRUSTED_URL } from '$env/static/public';
	const eventSourceGas = source(`${PUBLIC_TRUSTED_URL}/stream`).select('message');
	console.log(eventSourceGas);
</script>

<nav>
	<div class="container">
		<div class="logo">Logo</div>
		<div class="navbar">
			<a href="/">home</a>
			<!-- svelte-ignore a11y-invalid-attribute -->
			<a href="#">about</a>
			<!-- svelte-ignore a11y-invalid-attribute -->
			<a href="#">sign in</a>
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

	.navbar a {
		font-size: 18px;
		font-weight: bold;
		height: 100%;
		justify-content: space-between;
		margin-right: 30px;

	}
	.navbar a:hover {
		color: brown;
	}

	footer{
		position: sticky;
		bottom: 0;
		background-color: wheat;
		width: 100%;
	}

	footer section{
		text-align: center;
	}
</style>
