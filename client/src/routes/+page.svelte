<script>
	import AboutMe from '$lib/sections/aboutMe.svelte';
	import Education from '$lib/sections/education.svelte';
	import Header from '$lib/sections/header.svelte';
	import Interests from '$lib/sections/interests.svelte';
</script>

<Header />
<AboutMe />
<Interests />
<Education />

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
		--heading-font-size: 4rem;
		--subheading-font-size: 2.5rem;
		--body-font-size: 1.5rem;
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