<div class="aboutMeContainer">
	<div class="introTextContainer">
		<h1>Hi! I'm Reezan</h1>
		<h2>I'm a Computer Engineering Student at the University of Waterloo</h2>
		<p>
			I am passionate about both software and hardware, and I'm interested in the boundaries at
			which they meet
		</p>
	</div>
	<div class="headshotContainer">
		<enhanced:img src="/static/me.webp" alt="A picture of me" />
	</div>
</div>

<style>
	.aboutMeContainer {
		height: 60vh;
		display: flex;
		align-items: center;
		justify-content: space-around;
	}

	.introTextContainer {
		width: 42%;
	}

	.headshotContainer {
		width: 42%;
		text-align: right;
	}

	.headshotContainer img {
		width: 400px;
		height: auto;
	}

	h1 {
		font-size: 5rem;
		color: var(--text-primary-colour);
	}

	h2 {
		font-size: 3rem;
	}

	p {
		font-size: 1.5rem;
	}
</style>
