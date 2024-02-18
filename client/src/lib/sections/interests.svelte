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
				cardTitle="Software Engineering"
				cardBulletPoints={softwareEngineeringBulletPoints}
				cardSvgPath="/cardIcons/software.svg"
			/>
		</div>
		<div class="interests-card">
			<InterestsCard
				cardTitle="Hardware Engineering"
				cardBulletPoints={hardwareEngineeringBulletPoints}
				cardSvgPath="/cardIcons/hardware.svg"
			/>
		</div>
		<div class="interests-card">
			<InterestsCard
				cardTitle="Cloud Infrastructure"
				cardBulletPoints={cloudInfraBulletPoints}
				cardSvgPath="/cardIcons/cloud.svg"
			/>
		</div>
	</div>
</div>

<style>
	.interests-container {
		text-align: center;
		min-height: 60vh;
	}

	.interests-container h1 {
		color: var(--text-primary-colour);
		font-size: var(--heading-font-size);
	}

	.interests-cards-container {
		margin-top: 2.5vh;
		display: flex;
		justify-content: space-between;
		width: 100%;
	}

	.interests-card {
		min-width: 400px;
		min-height: 400px;
	}
</style>
