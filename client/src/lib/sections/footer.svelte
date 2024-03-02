<script>
	import Slider from '$lib/components/slider.svelte';
</script>

<footer class="footer">
	<div class="links-container">
		<h2>Reezan Visram</h2>
		<h3>Links:</h3>
		<a href="https://github.com/reezanvisram" target="_blank" class="icon-link">
			<img src="/cardIcons/github.svg" alt="Github Icon" />
			<p>Github</p>
		</a>
		<a href="https://linked.com/in/reezan-visram" target="_blank" class="icon-link">
			<img src="/cardIcons/linkedin.svg" alt="LinkedIn Icon" />
			<p>LinkedIn</p>
		</a>
	</div>
	<div class="icons-and-slider-container">
		<div>
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
		</div>
		<Slider />
	</div>
</footer>

<style>
	.footer {
		display: flex;
		justify-content: space-between;
		margin-top: 15vh;
		min-height: 15vh;
	}

	.links-container {
		display: flex;
		flex-direction: column;
		justify-content: space-around;
	}

	.links-container h2 {
		font-size: var(--subheading-font-size);
		color: var(--text-primary-colour);
	}

	.links-container h3 {
		font-size: var(--body-font-size);
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
	}

	.technology-icons-container {
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 0.5rem;
	}

	.technology-icons-container img {
		height: 75px;
		width: auto;
	}

	.large {
		height: 50px !important;
		width: auto !important;
	}
</style>
