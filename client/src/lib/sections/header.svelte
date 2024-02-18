<nav class="header">
	<h2 class="header-logo">Reezan Visram</h2>

	<div class="header-links-container">
		<p class="header-link">Education</p>
		<p class="header-link">Experience</p>
		<p class="header-link">Projects</p>
		<p class="header-link">Contact Me</p>
		<p class="header-button">Resume</p>
	</div>
</nav>

<style>
	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		width: 100%;
		margin-top: 0.5vw;
	}

	.header-logo {
		font-size: 2em;
		color: var(--text-primary-colour);
		margin-right: 25%;
	}

	.header-link {
		font-size: 1.5em;
		color: var(--text-primary-colour);
	}

	.header-links-container {
		display: flex;
		flex-grow: 1;
		justify-content: space-between;
		align-items: center;
	}

	.header-button {
		border-radius: 10px;
		outline: none;
		border: none;
		font-size: 1.5em;
		color: var(--button-text-colour);
		background-color: var(--text-primary-colour);
		padding: 0.5em 2.5em;
		text-align: center;
	}

	.header-button:hover {
		cursor: pointer;
		background-color: #0045ad;
	}
</style>
