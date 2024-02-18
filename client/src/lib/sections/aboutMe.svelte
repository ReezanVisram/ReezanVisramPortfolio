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
		<enhanced:img src="/static/me.webp" alt="A picture of me" />
	</div>
</div>

<style>
	.about-me-container {
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.intro-text-container {
		width: 42%;
	}

	.headshot-container {
		width: 42%;
		text-align: right;
	}

	.headshot-container img {
		width: 400px;
		height: auto;
	}

	h1 {
		font-size: var(--heading-font-size);
		color: var(--text-primary-colour);
	}

	h2 {
		font-size: var(--subheading-font-size);
	}

	p {
		font-size: var(--body-font-size);
	}
</style>
