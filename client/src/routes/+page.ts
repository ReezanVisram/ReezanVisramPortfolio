import type { PageLoad } from './$types';
import { PUBLIC_BASE_URL } from '$env/static/public';

export const load: PageLoad = async ({ fetch }) => {
	const experienceResponse = await fetch(`https://api.${PUBLIC_BASE_URL}/experience`);
	const experience = await experienceResponse.json();

	const projectsResponse = await fetch(`https://api.${PUBLIC_BASE_URL}/projects`);
	const projects = await projectsResponse.json();
	return {
		experience: experience,
		projects: projects
	};
};
