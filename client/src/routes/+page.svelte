<script lang="ts">
	import AboutMe from '$lib/sections/aboutMe.svelte';
	import ContactMe from '$lib/sections/contactMe.svelte';
	import Education from '$lib/sections/education.svelte';
	import Experience from '$lib/sections/experience.svelte';
	import Footer from '$lib/components/footer.svelte';
	import Header from '$lib/components/header.svelte';
	import Interests from '$lib/sections/interests.svelte';
	import Projects from '$lib/sections/projects.svelte';

	import type { PageData } from './$types';
	import { onMount } from 'svelte';

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

<section id="about-me-section">
	<AboutMe />
</section>

<section id="interests-section">
	<Interests />
</section>
<section id="education-section">
	<Education />
</section>
<section id="experience-section">
	<Experience experience={data.experience} />
</section>
<section id="projects-section">
	<Projects projects={data.projects} />
</section>
<section id="contact-me-section">
	<ContactMe />
</section>

<Footer />

<style>
	@font-face {
		src: url('../Outfit.ttf');
		font-family: Outfit;
	}

	:global(*) {
		margin: 0;
		padding: 0;
		font-family: Outfit, 'sans-serif';
	}

	:global(body) {
		margin-left: 12.5vw;
		margin-right: 12.5vw;
		background-color: var(--background-colour);
		--heading-font-size: 4rem;
		--subheading-font-size: 2.5rem;
		--body-font-size: 1.5rem;
		--svg-filter: brightness(0) saturate(100%) invert(18%) sepia(39%) saturate(5861%)
			hue-rotate(211deg) brightness(91%) contrast(94%);
		transition: background 0.4s;
	}

	:global([data-theme='light']) {
		--background-colour: #ffffff;
		--text-primary-colour: #0a62e6;
		--text-secondary-colour: #000000;
		--text-tertiary-colour: #ffffff;
		--button-text-colour: #ffffff;
	}

	:global([data-theme='dark']) {
		--background-colour: #000000;
		--text-primary-colour: #0a62e6;
		--text-secondary-colour: #ffffff;
		--text-tertiary-colour: #000000;
		--button-text-colour: #000000;
	}
</style>
