<script lang="ts">
	export let jobTitle = '';
	export let company = '';
	export let startAndEndDate = '';
</script>

<div class="experience-card">
	<div class="experience-information">
		<h3>{`${jobTitle} at ${company}`}</h3>
		<h3>{startAndEndDate}</h3>
	</div>
	<img src="/cardIcons/plus.svg" width={50} alt="Plus Icon" />
</div>

<style>
	.experience-card {
		position: relative;
		width: 100%;
		background-color: var(--text-secondary-colour);
		border-radius: 10px;
		text-align: left;
	}

	.experience-information {
		color: var(--text-tertiary-colour);
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
</style>
