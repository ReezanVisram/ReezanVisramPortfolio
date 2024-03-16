<script>
	import InterestsCard from '$lib/components/interestsCard.svelte';

	let softwareEngineeringBulletPoints = [
		'Frontend and Backend Development',
		'Cybersecurity',
		'Systems Programming',
		'Computer Graphics'
	];

	let hardwareEngineeringBulletPoints = [
		'Embedded Systems',
		'FPGA Design',
		'Circuit Design and Analysis',
		'ASIC Design and Verification'
	];

	let cloudInfraBulletPoints = [
		'CI/CD Pipelines',
		'Cloud System Design',
		'Infrastructure as Code',
		'Distributed Systems'
	];
</script>

<div class="interests-container">
	<h1>My Interests</h1>
	<div class="interests-cards-container">
		<div class="interests-card">
			<InterestsCard
				frontCardTitle="Software Engineering"
				frontCardBulletPoints={softwareEngineeringBulletPoints}
				frontCardSvgPath="/cardIcons/software.svg"
				backCardTitle="Sports"
				backCardText="I love both watching and playing sports. My favourite sport is hockey, and my favourite team is the Toronto Maple Leafs. Go Leafs Go!"
				backCardSvgPath="/cardIcons/hockey.svg"
			/>
		</div>
		<div class="interests-card">
			<InterestsCard
				frontCardTitle="Hardware Engineering"
				frontCardBulletPoints={hardwareEngineeringBulletPoints}
				frontCardSvgPath="/cardIcons/hardware.svg"
				backCardTitle="Video Games"
				backCardText="I like to relax by playing video games. Some of my favourite games are Hades, Sid Meier's Civilization and Hogwarts: Legacy."
				backCardSvgPath="/cardIcons/video-games.svg"
			/>
		</div>
		<div class="interests-card">
			<InterestsCard
				frontCardTitle="Cloud Infrastructure"
				frontCardBulletPoints={cloudInfraBulletPoints}
				frontCardSvgPath="/cardIcons/cloud.svg"
				backCardTitle="Music"
				backCardText="You can often find me with headphones in listening to music. I enjoy basically all genres of music, from Hip-Hop and Afro to Rock and EDM."
				backCardSvgPath="/cardIcons/music.svg"
			/>
		</div>
	</div>
</div>

<style>
	.interests-container {
		text-align: center;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: space-around;
	}

	.interests-container h1 {
		color: var(--text-primary-colour);
		font-size: var(--heading-font-size);
	}

	.interests-cards-container {
		margin-top: 2.5vh;
		display: flex;
		gap: 0.5vw;
	}

	.interests-card {
		width: 400px;
		height: 400px;
	}

	/* @media (max-width: 1366px) {
		.interests-cards-container {
			flex-direction: column;
		}
	} */

	@media (max-width: 1280px) {
		.interests-cards-container {
			gap: 1vw;
		}

		.interests-card {
			width: 350px;
			height: 350px;
		}
	}

	@media (max-width: 1100px) {
		.interests-cards-container {
			flex-direction: column;
			gap: 1vh;
		}
	}

	@media (max-width: 800px) {
		.interests-card {
			width: 300px;
			height: 300px;
		}
	}
</style>
