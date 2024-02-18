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

<div class="interestsContainer">
	<h1>My Interests</h1>
	<div class="interestsCardsContainer">
		<div class="interestsCard">
			<InterestsCard
				cardTitle="Software Engineering"
				cardBulletPoints={softwareEngineeringBulletPoints}
				cardSvgPath="/cardIcons/software.svg"
			/>
		</div>
		<div class="interestsCard">
			<InterestsCard
				cardTitle="Hardware Engineering"
				cardBulletPoints={hardwareEngineeringBulletPoints}
				cardSvgPath="/cardIcons/hardware.svg"
			/>
		</div>
		<div class="interestsCard">
			<InterestsCard
				cardTitle="Cloud Infrastructure"
				cardBulletPoints={cloudInfraBulletPoints}
				cardSvgPath="/cardIcons/cloud.svg"
			/>
		</div>
	</div>
</div>

<style>
	.interestsContainer {
		text-align: center;
		min-height: 60vh;
	}

	.interestsContainer h1 {
		color: var(--text-primary-colour);
		font-size: 4rem;
	}

	.interestsCardsContainer {
		margin-top: 2.5vh;
		display: flex;
		justify-content: space-around;
		width: 100%;
	}

	.interestsCard {
		min-width: 400px;
		min-height: 400px;
	}
</style>
