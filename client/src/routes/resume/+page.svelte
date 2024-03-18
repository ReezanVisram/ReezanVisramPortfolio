<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_BASE_URL } from '$env/static/public';
	import PdfViewer from '$lib/components/pdfViewer.svelte';
	import { theme } from '../../stores';

	const navigateHome = () => {
		goto('/');
	};

	let themeValue: string;
	theme.subscribe((value) => {
		themeValue = value;
	});
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
			<img
				class="button-icon back-icon"
				src={`/${themeValue}/back.svg`}
				alt="back arrow"
				width={45}
			/>
		</button>
		<a href={`https://api.${PUBLIC_BASE_URL}/resume`} download="Reezan_Visram_Resume.pdf">
			<button class="action-button">
				<img
					class="button-icon"
					src={`/${themeValue}/download.svg`}
					alt="download arrow"
					width={45}
				/>
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

	.heading-container {
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		align-items: center;
		margin-bottom: 2.5vh;
		min-height: 12.5vh;
	}

	.resume-heading {
		font-size: var(--subheading-font-size);
		color: var(--text-primary-colour);
		text-align: center;
	}

	.action-buttons-container {
		display: flex;
		width: calc(0.5 * (100% + 25vw));
		justify-content: space-between;
	}

	.action-button {
		background: var(--background-colour);
		outline: none;
		border: 3px solid var(--text-primary-colour);
		padding: 2px;
		border-radius: 10px;
		transition: transform 0.1s linear;
	}

	.action-button:hover {
		cursor: pointer;
		transform: scale(105%);
		border: 3px solid var(--text-secondary-colour);
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

	@media (max-width: 1600px) {
		.action-buttons-container {
			width: calc(0.75 * (100% + 25vw));
		}
	}

	@media (max-width: 1024px) {
		.action-buttons-container {
			width: calc(0.9 * (100% + 25vw));
		}
	}

	@media (max-width: 800px) {
		.action-buttons-container {
			width: calc(0.9 * (100% + 6vw));
		}
	}
</style>
