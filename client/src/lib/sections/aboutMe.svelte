<div class="about-me-container">
	<div class="intro-text-container">
		<h1>Hi! I'm Reezan</h1>
		<h2>I'm a Computer Engineering Student at the University of Waterloo</h2>
		<p>
			I'm passionate about both software and hardware, and I'm interested in the boundaries at which
			they meet!
		</p>
	</div>
	<div class="headshot-container">
		<enhanced:img class="headshot" src="/static/me.webp" alt="A picture of me" />
	</div>
</div>

<style>
	.about-me-container {
		min-height: 100vh;
		display: flex;
		align-items: center;
	}

	.intro-text-container {
		width: 42%;
	}

	.headshot-container {
		width: 33%;
		display: inline-block;
		position: relative;
		display: flex;
		justify-content: space-around;
		margin-right: auto;
		margin-left: auto;
		animation: fade-in 1s forwards 1.5s;
		opacity: 0;
	}

	.headshot-container::before {
		position: absolute;
		top: 3.5%;
		left: 7.5%;
		width: 100%;
		height: 100%;
		border: 4px solid var(--text-primary-colour);
		z-index: -1;
		content: '';
		border-radius: 10px;
		transition: 0.3s linear;
		box-shadow: rgba(99, 99, 99, 0.2) 0px 2px 8px 0px;
	}

	.headshot-container:hover.headshot-container::before {
		top: 1.5%;
		left: 5.5%;
	}

	.headshot {
		width: 100%;
		height: auto;
		border-radius: 10px;
	}

	h1 {
		font-size: var(--heading-font-size);
		color: var(--text-primary-colour);
		position: relative;
		opacity: 0;
		animation: fade-in 1s forwards;
	}

	h2 {
		font-size: var(--subheading-font-size);
		color: var(--text-secondary-colour);
		position: relative;
		opacity: 0;
		animation: fade-in 1s forwards 0.5s;
	}

	p {
		font-size: var(--body-font-size);
		color: var(--text-secondary-colour);
		position: relative;
		opacity: 0;
		animation: fade-in 1s forwards 1s;
	}

	@keyframes fade-in {
		0% {
			bottom: -0.2em;
		}

		50% {
			bottom: 0.2em;
		}

		100% {
			bottom: 0;
			opacity: 1;
		}
	}
</style>
