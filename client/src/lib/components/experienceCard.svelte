<script lang="ts">
	import ToolPill from './toolPill.svelte';

	export let jobTitle = '';
	export let company = '';
	export let startAndEndDate = '';
	export let bulletPoints: string[] = [];
	export let tools: string[] = [];

	let isActive = false;
	const handleOnClick = () => {
		isActive = !isActive;
	};
</script>

<div class="experience-card" class:active={isActive}>
	<button class="experience-card-button" class:button-active={isActive} on:click={handleOnClick}>
		<div class="experience-information">
			<h3>{`${jobTitle} at ${company}`}</h3>
			<h3>{startAndEndDate}</h3>
		</div>
		{#if !isActive}
			<img src="/cardIcons/plus.svg" width={75} alt="Plus Icon" />
		{:else}
			<img src="/cardIcons/minus.svg" width={75} alt="Minus Icon" class="minus" />
		{/if}
	</button>
	{#if isActive}
		<div class="bullet-points-container">
			<ul>
				{#each bulletPoints as bulletPoint}
					<li>{bulletPoint}</li>
				{/each}
			</ul>
			<div class="tools-container">
				{#each tools as tool}
					<ToolPill {tool} />
				{/each}
			</div>
		</div>
	{/if}
</div>

<style>
	.experience-card {
		width: 100%;
		background-color: var(--background-colour);
		border: 3px solid var(--text-secondary-colour);
		border-radius: 10px;
	}

	.experience-card-button {
		position: relative;
		width: 100%;
		background-color: var(--background-colour);
		border-radius: 10px;
		text-align: left;
		border: none;
		outline: none;
	}

	.active {
		border: 3px solid var(--text-primary-colour);
	}

	.button-active {
		border-bottom: 3px solid var(--text-primary-colour);
	}

	.experience-card-button:hover {
		cursor: pointer;
	}

	.experience-information {
		color: var(--text-secondary-colour);
		font-size: var(--body-font-size);
		padding: 0.5em;
	}

	.experience-information h3 {
		font-weight: 400;
	}

	.experience-card img {
		position: absolute;
		right: 0;
		top: 0;
		bottom: 0;
		margin-top: auto;
		margin-bottom: auto;
		margin-right: 10px;
	}

	.bullet-points-container {
		text-align: left;
		height: 100%;
		list-style-position: inside;
		font-size: var(--body-font-size);
		padding: 1vw;
	}

	.minus {
		color: var(--text-primary-colour);
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
