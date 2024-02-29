<div class="checkmark-container">
	<div class="checkmark-wrapper">
		<span class="checkmark"></span>
	</div>
</div>

<style>
	:root {
		--check-size: 150px;
		--check-border-width: 5px;
		--checkmark-width: calc(var(--check-size) / 2);
		--checkmark-height: calc(var(--checkmark-width) / 2);
		--checkmark-left: calc(var(--checkmark-width) / 2);
		--checkmark-top: 50%;
		--checkmark-colour: #04c230;
	}

	.checkmark-wrapper {
		width: var(--check-size);
		height: var(--check-size);
		position: relative;
		margin-left: auto;
		margin-right: auto;
		margin-top: 1.5vh;
		margin-bottom: 1.5vh;
	}

	.checkmark-wrapper::after {
		content: '';
		position: absolute;
		inset: 0;
		border: var(--check-border-width) solid #f8f8f8;
		width: 100%;
		height: 100%;
		border-radius: 50%;
		display: block;
		z-index: 0;
	}

	.checkmark-wrapper::before {
		content: '';
		position: absolute;
		inset: 0;
		border: var(--check-border-width) solid transparent;
		border-left-color: var(--checkmark-colour);
		width: 100%;
		height: 100%;
		border-radius: 50%;
		display: block;
		z-index: 1;
		animation: circle linear forwards 0.75s;
	}

	.checkmark {
		height: var(--checkmark-height);
		width: var(--checkmark-width);
		position: absolute;
		opacity: 0;
		left: var(--checkmark-left);
		top: var(--checkmark-top);
		display: block;
		border-left: var(--check-border-width) solid var(--checkmark-colour);
		border-bottom: var(--check-border-width) solid var(--checkmark-colour);
		transform-origin: left top;
		transform: rotate(-45deg);
		animation: checkmark linear both 1s;
	}

	@keyframes circle {
		0% {
			border-color: transparent;
			border-left-color: var(--checkmark-colour);
		}
		90% {
			transform: rotate(-360deg);
			border-color: transparent;
			border-left-color: var(--checkmark-colour);
		}
		100% {
			transform: rotate(-360deg);
			border-color: var(--checkmark-colour);
		}
	}

	@keyframes checkmark {
		0% {
			height: 0;
			width: 0;
			opacity: 0;
		}
		80% {
			height: 0;
			width: 0;
			opacity: 0;
		}
		90% {
			height: var(--checkmark-height);
			width: 0;
			opacity: 1;
		}
		100% {
			height: var(--checkmark-height);
			width: var(--checkmark-width);
			opacity: 1;
		}
	}
</style>
