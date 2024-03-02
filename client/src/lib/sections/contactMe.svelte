<script lang="ts">
	import Checkmark from '$lib/components/checkmark.svelte';
	import Modal from '$lib/components/modal.svelte';
	import AnimatedX from '$lib/components/animatedX.svelte';
	import { PUBLIC_BASE_URL } from '$env/static/public';

	let showModal = false;

	let modalHeading = 'Message Sent Successfully!';
	let modalBody =
		"Thank you for your message. I'll try my best to respond within the next couple of days!";
	let isCheckmark: boolean = false;

	const handleSubmit = async (e: SubmitEvent) => {
		const token = await window.grecaptcha.execute('6LcXPIEpAAAAAG9e43MfkzzIHoxA6C0PehTNnI-w', {
			action: 'submit'
		});

		const formData = new FormData(e.target as HTMLFormElement);

		const data: any = {};
		formData.forEach((value, key) => (data[key] = value));
		data.token = token;

		const body = JSON.stringify(data);

		const res = await fetch(`https://api.${PUBLIC_BASE_URL}/message`, {
			method: 'POST',
			body: body
		});

		if (res.status === 200 || res.status === 204) {
			modalHeading = 'Message Sent Successfully!';
			modalBody =
				"Thank you for your message. I'll try my best to respond within the next couple of days!";
			isCheckmark = true;
			showModal = true;
			const target = e.target as HTMLFormElement;
			target.reset();
		} else {
			modalHeading = 'Unable to Send Message';
			modalBody = 'Sorry, your message was not successfully sent. Please try again later.';
			isCheckmark = false;
			showModal = true;
			const target = e.target as HTMLFormElement;
			target.reset();
		}
	};
</script>

<Modal bind:showModal bind:isCheckmark>
	<h3 class="modal-heading" slot="dialog-heading">{modalHeading}</h3>

	<p class="modal-body" slot="dialog-content">
		{modalBody}
	</p>
</Modal>

<div class="contact-me-container">
	<h1>Contact Me</h1>
	<h3>Want to get in touch? Send me a message!</h3>

	<form class="contact-form-container" on:submit|preventDefault={handleSubmit}>
		<div class="input-container">
			<input type="text" name="name" required autocomplete="off" class="input" placeholder="" />
			<label class="placeholder-text" for="name" id="placeholder-name">Name</label>
		</div>

		<div class="input-container">
			<input type="text" name="email" required autocomplete="off" class="input" placeholder="" />
			<label class="placeholder-text" for="email" id="placeholder-email">Email</label>
		</div>

		<div class="input-container">
			<input type="text" name="subject" required autocomplete="off" class="input" placeholder="" />
			<label class="placeholder-text" for="subject" id="placeholder-email">Subject</label>
		</div>
		<div class="input-container">
			<textarea
				rows="20"
				cols="30"
				name="message"
				required
				autocomplete="off"
				class="input"
				placeholder=""
			/>
			<label class="placeholder-text" for="message" id="placeholder-email">Message</label>
		</div>

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

	.input-container {
		position: relative;
		margin-bottom: 0.25vh;
	}

	.input {
		width: 100%;
		height: 100%;
		font-size: var(--body-font-size);
		border: 3px solid var(--text-secondary-colour);
		border-radius: 10px;
		outline: none;
		background: none;
		z-index: 1;
		padding: 1rem;
		box-sizing: border-box;
		width: 100%;
	}

	.placeholder-text {
		position: absolute;
		left: 1rem;
		top: 1rem;
		padding: 0 0.25rem;
		background-color: var(--background-colour);
		color: var(--text-secondary-colour);
		font-size: var(--body-font-size);
		transition: 0.3s;
		pointer-events: none;
	}

	.input:focus + .placeholder-text {
		top: -0.7rem;
		left: 0.8rem;
		color: var(--text-primary-colour);
		font-size: calc(var(--body-font-size) - 0.3rem);
		z-index: 10;
	}

	.input:not(:placeholder-shown).input:not(:focus) + .placeholder-text {
		top: -0.7rem;
		left: 0.8rem;
		font-size: calc(var(--body-font-size) - 0.3rem);
		z-index: 10;
	}

	.input:focus {
		border: 3px solid var(--text-primary-colour);
	}

	.submit-button {
		outline: none;
		border-radius: 10px;
		background-color: var(--text-primary-colour);
		color: var(--text-tertiary-colour);
		border: none;
		font-size: var(--body-font-size);
		padding: 1rem;
	}

	.submit-button:hover {
		cursor: pointer;
		background-color: #0045ad;
	}

	.modal-heading {
		font-size: var(--subheading-font-size);
	}

	.modal-body {
		font-size: var(--body-font-size);
	}
</style>
