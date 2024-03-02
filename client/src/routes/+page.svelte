<script lang="ts">
	import AboutMe from '$lib/sections/aboutMe.svelte';
	import ContactMe from '$lib/sections/contactMe.svelte';
	import Education from '$lib/sections/education.svelte';
	import Experience from '$lib/sections/experience.svelte';
	import Footer from '$lib/sections/footer.svelte';
	import Header from '$lib/sections/header.svelte';
	import Interests from '$lib/sections/interests.svelte';
	import Projects from '$lib/sections/projects.svelte';

	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<script
		async
		src="https://www.google.com/recaptcha/api.js?render=6LcXPIEpAAAAAG9e43MfkzzIHoxA6C0PehTNnI-w"
	>
	</script>
</svelte:head>

<Header />
<AboutMe />
<Interests />
<Education />
<Experience experience={data.experience} />
<Projects projects={data.projects} />
<ContactMe />
<Footer />

<style>
	@font-face {
		src: url('../Outfit.ttf');
		font-family: Outfit;
	}

	:global(:root) {
		--background-colour: #ffffff;
		--text-primary-colour: #0065ff;
		--text-secondary-colour: #000000;
		--text-tertiary-colour: #ffffff;
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
