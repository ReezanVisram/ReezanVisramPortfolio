<script lang="ts">
	const handleSubmit = async () => {
		const token = await window.grecaptcha.execute('6LcXPIEpAAAAAG9e43MfkzzIHoxA6C0PehTNnI-w', {
			action: 'submit'
		});

		console.log(token);
	};
</script>

<div class="contact-me-container">
	<h1>Contact Me</h1>
	<h3>Want to get in touch? Send me a message!</h3>

	<form class="contact-form-container" on:submit|preventDefault={handleSubmit}>
		<input placeholder="Name" type="text" name="name" />
		<input placeholder="Email" type="email" name="email" />
		<input placeholder="Subject" type="text" name="subject" />
		<textarea placeholder="Message" rows="20" cols="30" name="message" />
		<button class="g-recaptcha submit-button" type="submit">Submit</button>
		<small
			>This site is protected by reCAPTCHA and the Google
			<a href="https://policies.google.com/privacy">Privacy Policy</a> and
			<a href="https://policies.google.com/terms">Terms of Service</a> apply.
		</small>
	</form>
</div>

<style>
	.contact-me-container {
		text-align: center;
		min-height: 45vh;
		margin-top: 5vh;
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.contact-me-container h1 {
		color: var(--text-primary-colour);
		font-size: var(--heading-font-size);
	}

	.contact-me-container h3 {
		color: var(--text-secondary-colour);
		font-size: var(--subheading-font-size);
	}

	.contact-form-container {
		display: flex;
		flex-direction: column;
		width: 67%;
		justify-content: space-between;
		gap: 1vh;
	}

	.contact-form-container input {
		font-size: var(--body-font-size);
		border-radius: 10px;
		border: 3px solid var(--text-secondary-colour);
		padding: 0.25em;
	}

	.contact-form-container textarea {
		resize: none;
		font-size: var(--body-font-size);
		border-radius: 10px;
		border: 3px solid var(--text-secondary-colour);
		padding: 0.25em;
	}

	.submit-button {
		outline: none;
		border-radius: 10px;
		background-color: var(--text-primary-colour);
		color: var(--text-tertiary-colour);
		border: none;
		font-size: var(--body-font-size);
		padding: 0.5em;
	}

	.submit-button:hover {
		cursor: pointer;
		background-color: #0045ad;
	}
</style>
