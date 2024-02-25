<script lang="ts">
	export let tool = '';
</script>

<div class="tool-pill-container">
	<p>{tool}</p>
</div>

<style>
	.tool-pill-container {
		background-color: var(--text-primary-colour);
		color: var(--text-tertiary-colour);
		font-size: var(--body-font-size);
		display: flex;
		justify-content: space-around;
		align-items: center;
		border-radius: 30px;
		padding: 0.5em 1em;
	}
</style>
