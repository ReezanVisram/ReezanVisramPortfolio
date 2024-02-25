<script lang="ts">
	import ExperienceCard from '$lib/components/experienceCard.svelte';

	export let experience: any = [];
</script>

<div class="experience-container">
	<h1>Experience</h1>

	<div class="experience-cards-container">
		{#each experience as exp}
			<ExperienceCard
				jobTitle={exp.job_title}
				company={exp.name}
				startAndEndDate={exp.start_and_end_date}
				bulletPoints={exp.bullet_points}
				tools={exp.tools}
			/>
		{/each}
	</div>
</div>

<style>
	.experience-container {
		text-align: center;
		min-height: 45vh;
		margin-bottom: 15vh;
		display: flex;
		flex-direction: column;
	}

	.experience-container h1 {
		color: var(--text-primary-colour);
		font-size: var(--heading-font-size);
	}

	.experience-cards-container {
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		gap: 1vh;
		margin-top: 2vh;
	}
</style>
