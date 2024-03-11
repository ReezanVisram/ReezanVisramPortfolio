<script lang="ts">
	let isLightMode = true;

	const handleToggle = () => {
		isLightMode = !isLightMode;
	};
</script>

<div class="slider">
	<img src="/cardIcons/sun.svg" alt="Sun Icon" width={75} class:active={isLightMode} />
	<input type="checkbox" id="darkmode-toggle" on:click={handleToggle} />
	<label for="darkmode-toggle" class="label"></label>
	<img src="/cardIcons/moon.svg" alt="Moon Icon" width={75} class:active={!isLightMode} />
</div>

<style>
	.label {
		width: 300px;
		height: 50px;
		position: relative;
		display: block;
		background-color: var(--background-colour);
		border-radius: 200px;
		border: 3px solid var(--text-secondary-colour);
		cursor: pointer;
		transition: 0.3s;
		box-sizing: border-box;
		transition: 0.3s;
	}

	.label:after {
		content: '';
		width: 65px;
		height: 50px;
		position: absolute;
		top: -3px;
		left: -3px;
		background: var(--text-primary-colour);
		border-radius: 200px;
		border: 3px solid var(--text-secondary-colour);
		transition: 0.3s;
		box-sizing: border-box;
	}

	input {
		width: 0;
		height: 0;
		visibility: hidden;
	}

	input:checked + .label:after {
		left: 297px;
		transform: translateX(-100%);
	}

	.slider {
		display: flex;
		align-items: center;
		gap: 0.3vw;
	}

	img {
		transition: 0.3s;
	}

	.active {
		filter: var(--svg-filter);
	}
</style>
