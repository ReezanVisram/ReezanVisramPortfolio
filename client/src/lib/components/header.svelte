<script lang="ts">
	import Hamburger from './hamburger.svelte';

	export let isMobileScreen: boolean;

	const scrollIntoView = (event: MouseEvent) => {
		const target = event.currentTarget as HTMLAnchorElement;

		const targetHref = target.getAttribute('href');
		if (!targetHref) {
			return;
		}
		const elementToScrollTo = document.querySelector(targetHref);

		if (!elementToScrollTo) {
			return;
		}

		elementToScrollTo.scrollIntoView({
			behavior: 'smooth'
		});
	};
</script>

<nav class="header">
	<h2 class="header-logo">Reezan Visram</h2>

	{#if !isMobileScreen}
		<div class="header-links-container">
			<a class="header-link" href="#education-section" on:click|preventDefault={scrollIntoView}
				>Education</a
			>
			<a class="header-link" href="#experience-section" on:click|preventDefault={scrollIntoView}
				>Experience</a
			>
			<a class="header-link" href="#projects-section" on:click|preventDefault={scrollIntoView}
				>Projects</a
			>
			<a class="header-link" href="#contact-me-section" on:click|preventDefault={scrollIntoView}
				>Contact Me</a
			>
			<a class="header-button" href="/resume">Resume</a>
		</div>
	{:else}
		<Hamburger />
	{/if}
</nav>

<style>
	a {
		text-decoration: none;
	}
	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		width: 100%;
		margin-top: 0.5vh;
	}

	.header-logo {
		font-size: 2rem;
		color: var(--text-primary-colour);
		flex-grow: 1;
	}

	.header-link {
		font-size: 1.5rem;
		color: var(--text-primary-colour);
	}

	.header-links-container {
		display: flex;
		flex-grow: 1;
		justify-content: space-between;
		align-items: center;
	}

	.header-button {
		border-radius: 10px;
		outline: none;
		border: none;
		font-size: 1.5rem;
		color: var(--button-text-colour);
		background-color: var(--text-primary-colour);
		padding: 0.5rem 2.5rem;
		text-align: center;
	}

	.header-button:hover {
		cursor: pointer;
		background-color: #0045ad;
	}
</style>
