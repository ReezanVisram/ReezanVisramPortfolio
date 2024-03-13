<script lang="ts">
	import { theme } from '../../stores';
	import { onMount } from 'svelte';

	let isLightMode = true;

	const determineTheme = () => {
		isLightMode = document.documentElement.dataset.theme == 'light' ? true : false;
		theme.update((theme) => {
			if (isLightMode) {
				theme = 'light';
			} else {
				theme = 'dark';
			}

			return theme;
		});
	};

	onMount(determineTheme);

	const handleToggle = (themeToToggle: string) => {
		isLightMode = !isLightMode;
		theme.update((theme) => {
			if (isLightMode) {
				theme = 'light';
			} else {
				theme = 'dark';
			}

			return theme;
		});

		document.documentElement.dataset.theme = themeToToggle;
		document.cookie = `siteTheme=${themeToToggle};max-age=31536000;path="/"`;
	};

	let themeValue: string;
	theme.subscribe((value) => {
		themeValue = value;
	});
</script>

<div class="slider">
	<img src={`/${themeValue}/sun.svg`} alt="Sun Icon" width={50} class:active={isLightMode} />
	<input
		type="checkbox"
		id="darkmode-toggle"
		on:click={() => handleToggle(isLightMode ? 'dark' : 'light')}
		checked={!isLightMode}
	/>
	<label for="darkmode-toggle" class="label"></label>
	<img src={`/${themeValue}/moon.svg`} alt="Moon Icon" width={50} class:active={!isLightMode} />
</div>

<style>
	.label {
		width: 100px;
		height: 45px;
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
		width: 50px;
		height: 45px;
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
		left: 97px;
		transform: translateX(-100%);
	}

	.slider {
		display: flex;
		align-items: center;
		gap: 0.3vw;
	}

	.active {
		filter: var(--svg-filter);
	}
</style>
