<script lang="ts">
	export let cardTitle = 'Default Card Title';
	export let cardBulletPoints: string[] = [];
	export let cardSvgPath = '/cardIcons/software.svg';
</script>

<div class="interestsCardContainer">
	<h3>{cardTitle}</h3>
	<ul class="bulletPointsContainer">
		{#each cardBulletPoints as bulletPoint}
			<li>{bulletPoint}</li>
		{/each}
	</ul>
	<img src={cardSvgPath} alt="Card Type Icon" width={150} />
</div>

<style>
	.interestsCardContainer {
		position: relative;
		border: 2px solid #000000;
		border-radius: 10px;
		width: 100%;
		height: 100%;
	}

	.interestsCardContainer h3 {
		font-size: 3rem;
		padding-left: 1rem;
		padding-right: 1rem;
		padding-top: 0.5rem;
	}

	.bulletPointsContainer {
		text-align: left;
		font-size: 1.5rem;
		list-style-position: inside;
		padding-top: 2vh;
		margin-left: 2vw;
	}

	.interestsCardContainer img {
		position: absolute;
		bottom: 0;
		left: 0;
		margin: 10px;
		color: var(--text-primary-colour);
	}
</style>
