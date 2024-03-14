<script lang="ts">
	import { expandSidebar } from '../../stores';

	let buttonActive = false;

	const handleClick = () => {
		buttonActive = !buttonActive;
		expandSidebar.update((value) => (value = buttonActive));
	};

	let expandSidebarValue: boolean;
	expandSidebar.subscribe((value) => {
		expandSidebarValue = value;
	});

	$: if (!expandSidebarValue) {
		buttonActive = false;
	}
</script>

<button on:click={handleClick} class:active={buttonActive}>
	<span class="top"></span>
	<span class="middle"></span>
	<span class="bottom"></span>
</button>

<style>
	:root {
		--bar-width: 40px;
		--bar-height: 4px;
		--bar-gap: 8px;
		--button-height: calc(var(--bar-height) * 3 + var(--bar-gap) * 2);
	}

	button {
		background-color: var(--background-colour);
		width: max-content;
		position: relative;
		border: none;
		outline: none;
		display: flex;
		flex-direction: column;
		gap: var(--bar-gap);
		z-index: 3;
	}

	button:hover {
		cursor: pointer;
	}

	.active .top {
		transform: rotate(45deg) translateY(calc(var(--bar-height) / -2));
		width: calc(var(--button-height) * 1.41421356237);
	}

	.active .middle {
		width: 0;
		opacity: 0;
	}

	.active .bottom {
		transform: rotate(-45deg) translateY(calc(var(--bar-height) / 2));
		width: calc(var(--button-height) * 1.41421356237);
	}

	span {
		width: var(--bar-width);
		height: var(--bar-height);
		background-color: var(--text-primary-colour);
		transition: 0.3s;
		transform-origin: left center;
	}

	.top {
		top: 0%;
	}

	.middle {
		top: 40%;
		opacity: 1;
	}

	.bottom {
		top: 80%;
	}
</style>
