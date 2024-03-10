<script lang="ts">
	export let cardTitle = 'Default Card Title';
	export let cardBulletPoints: string[] = [];
	export let cardSvgPath = '/cardIcons/software.svg';
</script>

<div class="interests-card-container">
	<h3>{cardTitle}</h3>
	<ul class="bullet-points-container">
		{#each cardBulletPoints as bulletPoint}
			<li>{bulletPoint}</li>
		{/each}
	</ul>
	<img src={cardSvgPath} alt="Card Type Icon" width={150} />
</div>

<style>
	.interests-card-container {
		position: relative;
		border: 2px solid var(--text-secondary-colour);
		border-radius: 10px;
		width: 100%;
		height: 100%;
	}

	.interests-card-container h3 {
		font-size: var(--subheading-font-size);
		color: var(--text-secondary-colour);
		padding-left: 1rem;
		padding-right: 1rem;
		padding-top: 0.5rem;
	}

	.bullet-points-container {
		text-align: left;
		font-size: var(--body-font-size);
		color: var(--text-secondary-colour);
		list-style-position: inside;
		padding-top: 2vh;
		margin-left: 0.5vw;
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
