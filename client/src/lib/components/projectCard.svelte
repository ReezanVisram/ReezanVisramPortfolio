<script lang="ts">
	import { onMount } from 'svelte';
	import ToolPill from './toolPill.svelte';
	import { fade } from 'svelte/transition';

	export let name: string = '';
	export let description: string = '';
	export let repoLink: string = '';
	export let releaseLink: string = '';
	export let imageLink: string = '';
	export let tools: string[] = [];
	export let leftImage = true;

	export let isVisible = false;
	let self: HTMLDivElement;
</script>

<div
	class="project-card-container"
	class:left-image={leftImage}
	class:is-visible={isVisible}
	bind:this={self}
>
	<div class="content-container" class:left-margin={!leftImage}>
		<div class="project-name-container">
			<h3>{name}</h3>
			<div class="icon-links-container">
				<a href={repoLink} target="_blank">
					<img src="/cardIcons/github.svg" width={50} alt="Github icon" />
				</a>
				<a href={releaseLink} target="_blank">
					<img src="/cardIcons/externalLink.svg" width={50} alt="External link icon" />
				</a>
			</div>
		</div>
		<p>{description}</p>
		<div class="tools-container">
			{#each tools as tool}
				<ToolPill {tool} />
			{/each}
		</div>
	</div>

	<div class="project-screenshot-container">
		<img src={imageLink} alt={`Screenshot of ${name}`} class="project-screenshot" />
	</div>
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
		flex-grow: 1;
		display: flex;
		flex-direction: column;
		justify-content: space-around;
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
		justify-content: flex-start;
		align-items: center;
		margin-top: 1vh;
		flex-wrap: wrap;
		gap: 1vh;
	}

	.project-screenshot-container {
		width: 40%;
		padding: 2em;
	}

	.project-screenshot {
		border-radius: 10px;
		width: 100%;
		height: auto;
	}

	.icon-links-container {
		margin-left: 1vw;
	}
</style>
