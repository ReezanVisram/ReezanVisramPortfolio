import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	const experienceResponse = await fetch('http://localhost:3000/experience');
	const experience = await experienceResponse.json();
	return {
		experience: experience
	};
};
