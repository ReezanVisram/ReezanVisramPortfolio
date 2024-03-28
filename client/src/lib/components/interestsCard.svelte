<script lang="ts">
	export let frontCardTitle = 'Default Card Title';
	export let frontCardBulletPoints: string[] = [];
	export let frontCardSvgPath = '/cardIcons/software.svg';

	export let backCardTitle = 'Default Card Title';
	export let backCardText: string = '';
	export let backCardSvgPath = '/cardIcons/software.svg';

	let isFrontShown = true;

	const handleClick = () => {
		isFrontShown = !isFrontShown;
	};
</script>

<button class="interests-card-container" on:click={handleClick} class:show-back={!isFrontShown}>
	<div class="interests-card-inner">
		<div class="interests-card-front">
			<h3>{frontCardTitle}</h3>
			<ul class="bullet-points-container">
				{#each frontCardBulletPoints as bulletPoint}
					<li>{bulletPoint}</li>
				{/each}
			</ul>
			<img src={frontCardSvgPath} alt="Card Type Icon" width={65} />
		</div>
		<div class="interests-card-back">
			<h3>{backCardTitle}</h3>
			<p>{backCardText}</p>
			<img src={backCardSvgPath} alt="Card Type Icon" width={65} />
		</div>
	</div>
</button>

<style>
	.interests-card-container {
		background-color: var(--background-colour);
		width: 100%;
		height: 100%;
		perspective: 1000px;
		border: none;
		outline: none;
	}

	.interests-card-container:hover {
		cursor: pointer;
	}

	.interests-card-inner {
		position: relative;
		width: 100%;
		height: 100%;
		transition: transform 0.6s;
		transform-style: preserve-3d;
		border: 3px solid var(--text-secondary-colour);
		border-radius: 10px;
	}

	.show-back .interests-card-inner {
		transform: rotateY(180deg);
	}

	.interests-card-front,
	.interests-card-back {
		position: absolute;
		width: 100%;
		height: 100%;
		-webkit-backface-visibility: hidden;
		backface-visibility: hidden;
		transform: rotateX(0deg);
	}

	.interests-card-back {
		transform: rotateY(180deg);
	}

	.interests-card-container h3 {
		font-size: var(--subheading-font-size);
		color: var(--text-secondary-colour);
		padding-left: 1rem;
		padding-right: 1rem;
		padding-top: 0.5rem;
	}

	.interests-card-container p {
		text-align: left;
		font-size: var(--body-font-size);
		color: var(--text-secondary-colour);
		padding-top: 2vh;
		margin: 0 0.5vw;
	}

	.bullet-points-container {
		text-align: left;
		font-size: var(--body-font-size);
		color: var(--text-secondary-colour);
		padding-top: 2vh;
		margin-left: 0.5vw;
	}

	.bullet-points-container li {
		margin-left: 1em;
	}

	.interests-card-container img {
		position: absolute;
		bottom: 0;
		left: 0;
		margin: 10px;
		color: var(--text-primary-colour);
		color: var(--text-secondary-colour);
		filter: var(--svg-filter);
	}
</style>
