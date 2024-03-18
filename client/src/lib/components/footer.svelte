<script lang="ts">
	import Slider from '$lib/components/slider.svelte';
	import { theme } from '../../stores';

	let themeValue: string;
	theme.subscribe((value) => {
		themeValue = value;
	});
</script>

<footer class="footer">
	<div class="links-container">
		<h2>Reezan Visram</h2>
		<h3>Links:</h3>
		<a href="https://github.com/reezanvisram" target="_blank" class="icon-link">
			<img src={`/${themeValue}/github.svg`} alt="Github Icon" />
			<p>Github</p>
		</a>
		<a href="https://linkedin.com/in/reezan-visram" target="_blank" class="icon-link">
			<img src={`/${themeValue}/linkedin.svg`} alt="LinkedIn Icon" />
			<p>LinkedIn</p>
		</a>
	</div>
	<div class="icons-and-slider-container">
		<p class="powered-by-text">Powered By</p>
		<div class="technology-icons-container">
			<img src="/cardIcons/go.svg" alt="Go Icon" class="large" />
			<img src="/cardIcons/go-chi.svg" alt="Go-Chi Icon" class="large" />
			<img src="/cardIcons/python.svg" alt="Python Icon" />
			<img src="/cardIcons/mongodb.svg" alt="MongoDB Icon" />
			<img src="/cardIcons/gcp.svg" alt="GCP Icon" />
			<img src="/cardIcons/svelte.svg" alt="Svelte Icon" />
			<img src="/cardIcons/typescript.svg" alt="TypeScript Icon" />
		</div>
		<Slider />
	</div>
</footer>

<style>
	.footer {
		display: flex;
		justify-content: space-between;
		margin-bottom: 2vh;
	}

	.links-container {
		display: flex;
		flex-direction: column;
		gap: 0.5vh;
	}

	.links-container h2 {
		font-size: var(--subheading-font-size);
		color: var(--text-primary-colour);
	}

	.links-container h3 {
		font-size: var(--body-font-size);
		color: var(--text-secondary-colour);
	}

	.icon-link {
		display: flex;
		color: var(--text-secondary-colour);
		text-decoration: none;
	}

	.icon-link p {
		font-size: var(--body-font-size);
	}

	.icon-link img {
		height: var(--subheading-font-size);
		width: auto;
	}

	.icons-and-slider-container {
		text-align: center;
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		align-items: center;
	}

	.powered-by-text {
		font-size: var(--body-font-size);
		color: var(--text-secondary-colour);
	}

	.technology-icons-container {
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 0.5rem;
	}

	.technology-icons-container img {
		height: 40px;
		width: auto;
	}

	.large {
		height: 25px !important;
		width: auto !important;
	}

	@media (max-width: 800px) {
		.links-container {
			justify-content: space-around;
		}
		.icons-and-slider-container {
			gap: 2vh;
		}
		.technology-icons-container {
			flex-wrap: wrap;
		}
	}
</style>
