<script>
	import AboutMe from '$lib/aboutMe.svelte';
	import Header from '$lib/header.svelte';
</script>

<Header />
<AboutMe />

<style>
	@font-face {
		src: url('./Outfit.ttf');
		font-family: Outfit;
	}

	:global(:root) {
		--background-colour: #ffffff;
		--text-primary-colour: #0065ff;
		--text-secondary-colour: #000000;
		--button-text-colour: #ffffff;
	}

	:global(*) {
		margin: 0;
		padding: 0;
		font-family: Outfit, 'sans-serif';
	}

	:global(body) {
		margin-left: 12.5vw;
		margin-right: 12.5vw;
	}
</style>
