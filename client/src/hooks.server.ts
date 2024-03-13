import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	const theme = event.cookies.get('siteTheme');
	const response = await resolve(event, {
		transformPageChunk: ({ html }) => {
			return html.replace('data-theme=""', `data-theme="${theme}"`);
		}
	});
	return response;
};
