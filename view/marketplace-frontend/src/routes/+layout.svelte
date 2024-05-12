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

<style>
	.container {
		display: flex;
		height: 60px;
		/* border-style: dotted; */
		/* border-color: blueviolet; */
        justify-content: space-between;
        align-items: center;
	}
    .logo{
        font-size:x-large;
        font-weight: bold;
        /* border-style: dotted; */
		/* border-color: palevioletred; */
        height: 50%;
    }

    .navbar a{
        font-size: 18px;
		font-weight: bold;
        height: 100%;
        justify-content: space-between;
        margin-left: 20px;
    }
    .navbar a:hover{
        color: brown;
    }
</style>
