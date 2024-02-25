<script lang="ts">
	import ProjectCard from '$lib/components/projectCard.svelte';

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
		<div class="heading-container">
			<button on:click={handleSoftwareClick}>
				<h3 class="project-type-header" class:active={isSoftwareActive}>Software</h3>
			</button>
			<p class="project-number">{softwareProjects.length}</p>
		</div>

		<div class="heading-container">
			<button on:click={handleHardwareClick}>
				<h3 class="project-type-header" class:active={!isSoftwareActive}>Hardware</h3>
			</button>
			<p class="project-number">{hardwareProjects.length}</p>
		</div>
	</div>

	<div class="software-projects-cards-container">
		{#each softwareProjects as project, index}
			<ProjectCard
				name={project.name}
				description={project.description}
				repoLink={project.repo_link}
				releaseLink={project.release_link}
				imageLink={project.image_link}
				tools={project.technologies}
				leftImage={index % 2 === 0}
			/>
		{/each}
	</div>
</div>

<style>
	.projects-container {
		text-align: center;
		min-height: 45vh;
		margin-top: 2.5vh;
		display: flex;
		flex-direction: column;
	}

	.projects-container h1 {
		color: var(--text-primary-colour);
		font-size: var(--heading-font-size);
	}

	.headings-container {
		display: flex;
		justify-content: space-around;
		width: 50%;
	}

	.project-type-header {
		font-size: var(--subheading-font-size);
		font-weight: 400;
	}

	.active {
		font-weight: bold;
		text-decoration: underline;
	}

	.headings-container button {
		outline: none;
		border: none;
		background-color: var(--background-colour);
	}

	.headings-container button:hover {
		cursor: pointer;
	}

	.heading-container {
		display: flex;
		align-items: center;
		gap: 1vw;
		align-items: center;
		width: 50%;
	}

	.project-number {
		font-size: var(--body-font-size);
		color: #aeaeae;
	}

	.software-projects-cards-container {
		display: flex;
		flex-direction: column;
		gap: 1vh;
	}
</style>
