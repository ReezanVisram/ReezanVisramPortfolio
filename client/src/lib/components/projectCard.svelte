<script lang="ts">
	import ToolPill from './toolPill.svelte';

	export let name: string = '';
	export let description: string = '';
	export let repoLink: string = '';
	export let releaseLink: string = '';
	export let imageLink: string = '';
	export let tools: string[] = [];
	export let leftImage = true;
</script>

<div class="project-card-container" class:left-image={leftImage}>
	<div class="content-container" class:left-margin={!leftImage}>
		<div class="project-name-container">
			<h3>{name}</h3>
			<a href={repoLink} target="_blank">
				<img src="/cardIcons/github.svg" width={75} alt="Github icon" />
			</a>
			<a href={releaseLink} target="_blank">
				<img src="/cardIcons/externalLink.svg" width={75} alt="External link icon" />
			</a>
		</div>
		<p>{description}</p>
		<div class="tools-container">
			{#each tools as tool}
				<ToolPill {tool} />
			{/each}
		</div>
	</div>

	<img src={imageLink} />
</div>

<style>
	.project-card-container {
		border: 3px solid var(--text-secondary-colour);
		display: flex;
		justify-content: space-between;
		border-radius: 10px;
	}

	.left-image {
		flex-direction: row-reverse;
	}

	.content-container {
		width: 50%;
		margin-bottom: 1vh;
		margin-top: 1vh;
		text-align: left;
	}

	.content-container h3 {
		font-size: var(--subheading-font-size);
	}

	.left-margin {
		margin-left: 1vw;
	}

	.content-container p {
		font-size: var(--body-font-size);
	}

	.project-name-container {
		display: flex;
	}

	.tools-container {
		display: flex;
		flex-direction: row;
		justify-content: space-around;
		align-items: center;
		margin-top: 1vh;
		flex-wrap: wrap;
		gap: 1vh;
	}
</style>
