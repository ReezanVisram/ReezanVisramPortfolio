<div class="education-container">
	<div class="text-container">
		<h1>Education</h1>
		<h2>University of Waterloo</h2>
		<h3>Candidate for Bachelor of Applied Science in Honours Computer Engineering Co-op</h3>
		<h3>September 2021 - April 2026</h3>
		<p>90% Cumulative Average</p>
		<p>Term Dean's Honours List</p>
		<p>Recepient of Richard and Elizabeth Madter Scholarship for Outstanding Students</p>
		<p>Recepient of President's Scholarship of Distinction for admission average of 99%</p>
	</div>
	<enhanced:img
		class="uwaterloo"
		src="/static/uwaterloo.webp"
		alt="A picture of the University of Waterlo Campus, taken by me from the E5 Fourth Floor Balcony"
	/>
</div>

<style>
	.education-container {
		display: flex;
		flex-direction: row-reverse;
		justify-content: space-between;
		min-height: 50vh;
	}

	.text-container {
		width: 33%;
		text-align: right;
	}

	.text-container h1 {
		font-size: var(--heading-font-size);
		color: var(--text-primary-colour);
	}

	.text-container h2 {
		font-size: var(--subheading-font-size);
		color: var(--text-secondary-colour);
	}

	.text-container h3 {
		font-size: var(--body-font-size);
		color: var(--text-secondary-colour);
	}

	.text-container p {
		font-size: var(--body-font-size);
		color: var(--text-secondary-colour);
	}

	.uwaterloo {
		width: 640px;
		height: auto;
		border-radius: 10px;
		box-shadow: rgba(99, 99, 99, 0.2) 0px 2px 8px 0px;
	}

	@media (max-width: 1600px) {
		.uwaterloo {
			width: 480px;
		}
	}

	@media (max-width: 1200px) {
		.education-container {
			flex-direction: column;
			justify-content: space-around;
			align-items: center;
		}

		.text-container {
			width: 100%;
			text-align: center;
		}

		.uwaterloo {
			width: 100%;
		}
	}
</style>
