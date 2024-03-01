<div class="wrapper">
	<svg class="checkmark" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 52 52"
		><circle class="checkmark_circle" cx="26" cy="26" r="25" fill="none" /><path
			class="checkmark_check"
			fill="none"
			d="M14.1 14.1l23.8 23.8 m0,-23.8 l-23.8,23.8"
		/></svg
	>
</div>

<style>
	.wrapper {
		transform: scale(2);
	}
	.checkmark {
		width: 56px;
		height: 56px;
		border-radius: 50%;
		display: block;
		stroke-width: 2;
		stroke: #ffffff;
		stroke-miterlimit: 10;
		box-shadow: inset 0px 0px 0px #ff5630;
		animation:
			fill 0.4s ease-in-out 0.4s forwards,
			scale 0.3s ease-in-out 0.9s both;
	}

	.checkmark_circle {
		stroke-dasharray: 166;
		stroke-dashoffset: 166;
		stroke-width: 2;
		stroke-miterlimit: 10;
		stroke: #ff5630;
		fill: none;
		animation: stroke 0.6s cubic-bezier(0.65, 0, 0.45, 1) forwards;
	}

	.checkmark_check {
		transform-origin: 50% 50%;
		stroke-dasharray: 48;
		stroke-dashoffset: 48;
		animation: stroke 0.3s cubic-bezier(0.65, 0, 0.45, 1) 0.8s forwards;
	}

	@keyframes stroke {
		100% {
			stroke-dashoffset: 0;
		}
	}
	@keyframes scale {
		0%,
		100% {
			transform: none;
		}
		50% {
			transform: scale3d(1.1, 1.1, 1);
		}
	}
	@keyframes fill {
		100% {
			box-shadow: inset 0px 0px 0px 30px #ff5630;
		}
	}
</style>
