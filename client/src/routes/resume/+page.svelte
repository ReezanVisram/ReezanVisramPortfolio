<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_BASE_URL } from '$env/static/public';
	import PdfViewer from '$lib/components/pdfViewer.svelte';

	const navigateHome = () => {
		goto('/');
	};
</script>

<svelte:head>
	<link
		rel="stylesheet"
		href="https://cdnjs.cloudflare.com/ajax/libs/pdf.js/4.0.379/pdf_viewer.css"
	/>
</svelte:head>
<div class="heading-container">
	<h1 class="resume-heading">My Resume</h1>
	<div class="action-buttons-container">
		<button class="action-button" on:click={navigateHome}>
			<img class="button-icon back-icon" src="/cardIcons/back.svg" alt="back arrow" width={45} />
		</button>
		<a href={`https://api.${PUBLIC_BASE_URL}/resume`} download="Reezan_Visram_Resume.pdf">
			<button class="action-button">
				<img class="button-icon" src="/cardIcons/download.svg" alt="download arrow" width={45} />
			</button>
		</a>
	</div>
</div>

<PdfViewer />

<style>
	@font-face {
		src: url('../../Outfit.ttf');
		font-family: Outfit;
	}

	:global(:root) {
		--background-colour: #ffffff;
		--text-primary-colour: #0848a7;
		--text-secondary-colour: #000000;
		--text-tertiary-colour: #ffffff;
		--button-text-colour: #ffffff;
		--heading-font-size: 4rem;
		--subheading-font-size: 2.5rem;
		--body-font-size: 1.5rem;
		--svg-filter: invert(18%) sepia(39%) saturate(5861%) hue-rotate(211deg) brightness(91%)
			contrast(94%);
	}

	:global(*) {
		margin: 0;
		padding: 0;
		font-family: Outfit;
	}

	:global(body) {
		margin-left: 12.5vw;
		margin-right: 12.5vw;
		height: 100vh;
	}

	.heading-container {
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		margin-bottom: 2.5vh;
		min-height: 12.5vh;
	}

	.resume-heading {
		font-size: var(--subheading-font-size);
		color: var(--text-primary-colour);
		text-align: center;
	}

	.action-buttons-container {
		text-align: center;
	}

	.action-button {
		background: var(--background-colour);
		outline: none;
		/* border: 3px solid var(--text-primary-colour); */
		border: none;
		padding: 2px;
		border-radius: 10px;
		transition: transform 0.1s linear;
		box-shadow: rgba(99, 99, 99, 0.2) 0px 2px 8px 0px;
	}

	.action-button:hover {
		cursor: pointer;
		transform: scale(105%);
		/* border: 3px solid black; */
	}

	.button-icon {
		filter: var(--svg-filter);
	}

	.action-button:hover .button-icon {
		filter: none;
	}

	.back-icon {
		transform: rotate(180deg);
	}
</style>
