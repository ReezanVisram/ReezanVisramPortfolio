<script>
	import Pdfseite from './pdfseite.svelte';
	import { onMount, tick } from 'svelte';
	import * as pdfjs from 'pdfjs-dist';
	import { GlobalWorkerOptions } from 'pdfjs-dist';

	GlobalWorkerOptions.workerSrc = `//cdnjs.cloudflare.com/ajax/libs/pdf.js/${pdfjs.version}/pdf.worker.js`;

	let geladen = false;

	export let url;
	export let data = '';
	let password = '';
	let pdfdok = null;
	let seinum = 1;
	let wartenummer = null;
	let seiteamladen = false;

	let dokument = [];
	let scale = 2;
	let viewport;

	let canvas;

	const loadPDF = async () => {
		console.log('Lädt: ' + url);

		let pdfladen = pdfjs.getDocument({
			...(url && { url }),
			...(data && { data })
		});

		pdfladen.promise
			.then(async function (pdf_) {
				console.log('Dokument geladen');

				pdfdok = pdf_;
				console.log('pdfdok def');

				console.log('Seiten: ' + pdfdok.numPages);

				await tick();
				geladen = true;
			})
			.catch(function (error) {
				console.log('PDF - Fehler: ' + error);
			});
	};

	$: if (geladen) andiearbeit(seinum);

	const andiearbeit = (num) => {
		if (seiteamladen) {
			wartenummer = num;
		} else {
			seitlad(num);
			console.log('An die Arbeit!');
		}
	};

	const seitlad = (num) => {
		console.log('Seitlad gestartet');
		seiteamladen = true;

		pdfdok.getPage(num).then(function (page) {
			viewport = page.getViewport({ scale: scale });
			console.log('Viewport: ' + viewport);

			let weitergeben = [page, viewport];

			dokument = [...dokument, weitergeben];
			seinum++;
			seiteamladen = false;
		});
	};

	onMount(loadPDF);
</script>

<div id="dokwrap">
	{#key dokument}
		{#each dokument as seite}
			<Pdfseite canvasseite={seite} />
		{/each}
	{/key}
</div>

<style>
	#dokwrap {
		border: 1pt solid #000;
		border-radius: 8px;
		height: 80vh;
	}
	canvas {
		width: 100%;
		height: 100%;
	}
</style>
