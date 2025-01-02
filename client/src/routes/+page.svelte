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
	import SideBar from '$lib/components/sideBar.svelte';
	import { expandSidebar } from '../stores';

	export let data: PageData;

	let isMobileScreen = false;

	let innerWidth: number;

	let expandSidebarValue: boolean;
	expandSidebar.subscribe((value) => {
		expandSidebarValue = value;
	});

	$: isMobileScreen = innerWidth <= 1200;
	$: if (expandSidebarValue && !isMobileScreen) {
		expandSidebar.update((value) => (value = false));
	}
</script>

<svelte:window bind:innerWidth />

<svelte:head>
	<script
		async
		src="https://www.google.com/recaptcha/api.js?render=6LcXPIEpAAAAAG9e43MfkzzIHoxA6C0PehTNnI-w"
	>
	</script>
</svelte:head>

{#if innerWidth}
	<Header {isMobileScreen} />
	{#if isMobileScreen && expandSidebarValue}
		<SideBar />
	{/if}

	<section id="about-me-section">
		<AboutMe />
	</section>
	<section id="education-section">
		<Education />
	</section>
	<section id="interests-section">
		<Interests />
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
{/if}

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
		--svg-filter: brightness(0) saturate(100%) invert(22%) sepia(97%) saturate(3338%)
			hue-rotate(212deg) brightness(96%) contrast(92%);
		transition: background 0.4s;
	}

	:global(.sideBarExpanded) {
		overflow-y: hidden;
	}

	:global([data-theme='light']) {
		--background-colour: #ffffff;
		--text-primary-colour: #0a62e6;
		--text-secondary-colour: #000000;
		--text-tertiary-colour: #ffffff;
		--button-text-colour: #ffffff;
		--backdrop-colour: rgba(0, 0, 0, 0.3);
	}

	:global([data-theme='dark']) {
		--background-colour: #18181d;
		--text-primary-colour: #0a62e6;
		--text-secondary-colour: #ffffff;
		--text-tertiary-colour: #18181d;
		--button-text-colour: #18181d;
		--backdrop-colour: rgba(255, 255, 255, 0.3);
	}

	@media (max-width: 800px) {
		:global(body) {
			--heading-font-size: 2.5rem;
			--subheading-font-size: 1.75rem;
			--body-font-size: 1rem;
			margin-left: 3vw;
			margin-right: 3vw;
		}
	}

	section {
		margin-bottom: 5vh;
	}
</style>
