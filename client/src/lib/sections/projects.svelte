<script lang="ts">
	import ProjectCard from '$lib/components/projectCard.svelte';
	import { fade, fly } from 'svelte/transition';

	export let projects: any[] = [];

	const softwareProjects: any[] = [];
	const hardwareProjects: any[] = [];

	projects.forEach((project: any) => {
		if (project.is_hardware) {
			hardwareProjects.push(project);
		} else {
			softwareProjects.push(project);
		}
	});

	let isSoftwareActive = true;

	const handleSoftwareClick = () => {
		isSoftwareActive = true;
	};

	const handleHardwareClick = () => {
		isSoftwareActive = false;
	};
</script>

<div class="projects-container">
	<h1>Projects</h1>

	<div class="headings-container">
		<div class="heading-container" class:active={isSoftwareActive}>
			<button on:click={handleSoftwareClick}>
				<h3 class="project-type-header">Software</h3>
			</button>
		</div>

		<div class="heading-container hardware" class:active={!isSoftwareActive}>
			<button on:click={handleHardwareClick}>
				<h3 class="project-type-header">Hardware</h3>
			</button>
		</div>
	</div>

	<div class="project-cards-container">
		{#if isSoftwareActive}
			{#each softwareProjects as project, index (project)}
				<div in:fade|global={{ delay: 400 + index * 100, duration: 200 }} out:fade|global>
					<ProjectCard
						name={project.name}
						description={project.description}
						repoLink={project.repo_link}
						releaseLink={project.release_link}
						imageLink={project.image_link}
						tools={project.technologies}
						leftImage={index % 2 === 0}
					/>
				</div>
			{/each}
		{:else}
			{#each hardwareProjects as project, index (project)}
				<div in:fade|global={{ delay: 400 + index * 100, duration: 200 }} out:fade|global>
					<ProjectCard
						name={project.name}
						description={project.description}
						repoLink={project.repo_link}
						releaseLink={project.release_link}
						imageLink={project.image_link}
						tools={project.technologies}
						leftImage={index % 2 === 0}
					/>
				</div>
			{/each}
		{/if}
	</div>
</div>

<style>
	.projects-container {
		text-align: center;
		min-height: 45vh;
		margin-top: 2.5vh;
		display: flex;
		flex-direction: column;
		position: relative;
	}

	.projects-container h1 {
		color: var(--text-primary-colour);
		font-size: var(--heading-font-size);
	}

	.headings-container {
		display: flex;
		justify-content: space-between;
		width: 100%;
	}

	.project-type-header {
		font-size: var(--subheading-font-size);
		font-weight: 400;
		width: 100%;
	}

	.headings-container button {
		outline: none;
		border: none;
		background-color: var(--background-colour);
	}

	.headings-container button:hover {
		cursor: pointer;
	}

	.heading-container h3 {
		color: var(--text-secondary-colour);
	}

	.heading-container h3 {
		transition: 0.2s ease-out;
	}

	.active h3 {
		font-weight: bold;
		color: var(--text-primary-colour);
	}

	.project-cards-container {
		display: flex;
		flex-direction: column;
		gap: 1vh;
		margin-top: 2.5vh;
	}
</style>
