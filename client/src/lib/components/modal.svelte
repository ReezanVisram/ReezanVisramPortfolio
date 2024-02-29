<script lang="ts">
	export let showModal: boolean = false;

	let dialog: HTMLDialogElement;

	$: if (showModal && dialog) dialog.showModal();
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
				<img src="/cardIcons/x.svg" width={32} alt="X icon" />
			</button>
		</div>

		<slot name="dialog-content" />

		<slot name="dialog-image" />
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
	}

	.dialog::backdrop {
		background: rgba(0, 0, 0, 0.3);
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
</style>
