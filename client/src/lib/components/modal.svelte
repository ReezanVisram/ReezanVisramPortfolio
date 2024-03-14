<script lang="ts">
	import { theme } from '../../stores';
	import AnimatedX from './animatedX.svelte';
	import Checkmark from './checkmark.svelte';

	export let showModal: boolean;
	export let isCheckmark: boolean;

	let dialog: HTMLDialogElement;

	$: if (showModal && dialog) dialog.showModal();

	let themeValue: string;

	theme.subscribe((value) => {
		themeValue = value;
	});
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog
	class="dialog"
	bind:this={dialog}
	on:close={() => (showModal = false)}
	on:click|self={() => dialog.close()}
>
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div class="dialog-container" on:click|stopPropagation>
		<div class="dialog-heading-container">
			<slot name="dialog-heading" />
			<button class="dialog-close-button" on:click={() => dialog.close()}>
				<img src={`/${themeValue}/x.svg`} width={32} alt="X icon" />
			</button>
		</div>

		<slot name="dialog-content" />

		<div class="modal-image-container">
			{#if isCheckmark}
				<Checkmark />
			{:else}
				<AnimatedX />
			{/if}
		</div>
	</div>
</dialog>

<style>
	.dialog {
		max-width: 50vw;
		border-radius: 10px;
		border: none;
		padding: 0;
		margin: auto;
		overflow: hidden;
		background-color: var(--background-colour);
	}

	.dialog::backdrop {
		background: var(--backdrop-colour);
	}

	.dialog[open] {
		animation: zoom-in cubic-bezier(0.34, 1.56, 0.64, 1) 0.5s;
	}

	@keyframes zoom-in {
		from {
			transform: scale(0.3);
		}
		to {
			transform: scale(1);
		}
	}

	.dialog[open]::backdrop {
		animation: fade ease-in-out 0.5s;
	}

	@keyframes fade {
		from {
			opacity: 0;
		}

		to {
			opacity: 1;
		}
	}

	.dialog-container {
		padding: 1em;
	}

	.dialog-heading-container {
		display: flex;
		justify-content: space-between;
	}

	.dialog-close-button {
		border: none;
		outline: none;
		background: none;
	}

	.dialog-close-button:hover {
		cursor: pointer;
	}

	.modal-image-container {
		width: 100%;
		display: flex;
		justify-content: space-around;
		margin: 2rem 0;
	}
</style>
