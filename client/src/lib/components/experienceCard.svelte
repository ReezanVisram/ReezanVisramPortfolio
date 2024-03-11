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

<div class="experience-card-wrapper" class:experience-card-wrapper-active={isActive}>
	<button
		class="experience-card-button"
		class:experience-card-button-active={isActive}
		on:click={handleOnClick}
	>
		<div class="experience-information">
			<h3>{`${jobTitle} at ${company}`}</h3>
			<h3>{startAndEndDate}</h3>
		</div>
		<div class="plus-minus-container">
			<span></span>
			<span></span>
		</div>
	</button>
	<div class="bullet-points-container" class:bullet-points-container-active={isActive}>
		<div class="bullet-points-content">
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
	</div>
</div>

<style>
	.experience-card-wrapper {
		border-radius: 10px;
		border: 3px solid var(--text-secondary-colour);
	}

	.experience-card-wrapper:hover {
		border: 3px solid var(--text-primary-colour);
	}

	.experience-card-wrapper-active {
		border: 3px solid var(--text-primary-colour);
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

	.experience-card-button:hover {
		cursor: pointer;
	}

	.experience-card-wrapper:hover .plus-minus-container span {
		background-color: var(--text-primary-colour);
	}

	.experience-card-wrapper-active .plus-minus-container span {
		background-color: var(--text-primary-colour);
	}

	.plus-minus-container {
		position: absolute;
		right: 0;
		width: 100px;
		height: 100px;
		top: 0;
		bottom: 0;
		margin-top: auto;
		margin-bottom: auto;
		margin-right: 10px;
		transition: transform 0.4s ease-out;
	}

	.plus-minus-container span {
		background-color: var(--text-secondary-colour);
	}

	.plus-minus-container span:first-child {
		position: absolute;
		height: 50px;
		inset: 0;
		margin-top: auto;
		margin-left: auto;
		margin-right: auto;
		margin-bottom: auto;
		width: 5px;
		border-radius: 10px;
	}

	.plus-minus-container span:last-child {
		position: absolute;
		height: 50px;
		inset: 0;
		margin-top: auto;
		margin-left: auto;
		margin-right: auto;
		margin-bottom: auto;
		width: 5px;
		border-radius: 10px;
		transform: rotate(90deg);
		transition: height 0.4s ease-out;
	}

	.experience-card-wrapper-active .plus-minus-container {
		transform: rotate(90deg);
	}

	.experience-card-wrapper-active .plus-minus-container span:last-child {
		height: 0px;
	}

	.experience-information {
		color: var(--text-secondary-colour);
		font-size: var(--body-font-size);
		padding: 0.5em;
	}

	.experience-information h3 {
		font-weight: 400;
	}

	.bullet-points-container {
		text-align: left;
		list-style-position: inside;
		font-size: var(--body-font-size);
		overflow: hidden;
		transition: max-height 0.4s ease-out;
		border: none;
		max-height: 0;
	}

	.bullet-points-container-active {
		max-height: 100vh;
	}

	.bullet-points-content {
		padding: 1vh 1vw;
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
