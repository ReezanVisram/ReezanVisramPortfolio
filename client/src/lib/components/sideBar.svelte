<script lang="ts">
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	import { expandSidebar } from '../../stores';

	onMount(() => {
		document.documentElement.style.overflow = 'hidden';
		document.documentElement.style.touchAction = 'none';

		return () => {
			document.documentElement.style.overflow = 'scroll';
			document.documentElement.style.touchAction = 'auto';
		};
	});

	const handleOverlayClicked = () => {
		expandSidebar.update((value) => (value = false));
	};

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

		expandSidebar.update((value) => (value = false));
	};
</script>

<div class="container" in:fade out:fade>
	<button class="body-overlay" on:click={handleOverlayClicked}></button>
	<div class="sidebar-container">
		<nav>
			<div class="sidebar-links-container">
				<a
					class="header-link"
					href="#education-section"
					on:click|preventDefault={scrollIntoView}
					in:fade={{ delay: 400 }}>Education</a
				>
				<a
					class="header-link"
					href="#experience-section"
					on:click|preventDefault={scrollIntoView}
					in:fade={{ delay: 450 }}>Experience</a
				>
				<a
					class="header-link"
					href="#projects-section"
					on:click|preventDefault={scrollIntoView}
					in:fade={{ delay: 500 }}>Projects</a
				>
				<a
					class="header-link"
					href="#contact-me-section"
					on:click|preventDefault={scrollIntoView}
					in:fade={{ delay: 550 }}>Contact Me</a
				>
				<a class="header-button" href="/resume" in:fade={{ delay: 600 }}>Resume</a>
			</div>
		</nav>
	</div>
</div>

<style>
	.container {
		position: absolute;
		height: 100vh;
		width: 100vw;
		inset: 0;
		z-index: 2;
	}

	.body-overlay {
		border: none;
		outline: none;
		background-color: var(--backdrop-colour);
		height: 100%;
		width: 50vw;
		left: 0;
		top: 0;
		overflow-y: hidden;
		touch-action: none;
	}

	.sidebar-container {
		position: absolute;
		right: 0;
		top: 0;
		height: 100%;
		width: 50vw;
		background-color: var(--background-colour);
		z-index: 3;
	}

	nav {
		height: 100%;
		width: 100%;
	}

	.sidebar-links-container {
		height: 90%;
		width: 100%;
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		align-items: center;
	}

	.header-link {
		font-size: 1.5rem;
		color: var(--text-primary-colour);
		text-decoration: none;
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
		text-decoration: none;
	}

	.header-button:hover {
		cursor: pointer;
		background-color: #0045ad;
	}
</style>
