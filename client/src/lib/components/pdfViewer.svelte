<script lang="ts">
	import * as pdfjs from 'pdfjs-dist';
	import * as pdfjsViewer from 'pdfjs-dist/web/pdf_viewer.js';
	import { GlobalWorkerOptions } from 'pdfjs-dist';
	import { onMount } from 'svelte';
	import { PUBLIC_BASE_URL } from '$env/static/public';

	GlobalWorkerOptions.workerSrc = `//cdnjs.cloudflare.com/ajax/libs/pdf.js/${pdfjs.version}/pdf.worker.js`;

	const defaultUrl = `https://api.${PUBLIC_BASE_URL}/resume`;

	const enableXfa = true;
	const searchFor = '';

	const sandboxBundleSrc = `//cdnjs.cloudflare.com/ajax/libs/pdf.js/${pdfjs.version}/pdf.sandbox.mjs`;

	let pdfViewer: pdfjsViewer.PDFViewer;

	onMount(async () => {
		const container = document.getElementById('viewerContainer') as HTMLDivElement;
		console.log(container);

		const eventBus = new pdfjsViewer.EventBus();

		const pdfLinkService = new pdfjsViewer.PDFLinkService({
			eventBus
		});

		const pdfFindController = new pdfjsViewer.PDFFindController({
			eventBus,
			linkService: pdfLinkService
		});

		const pdfScriptingManager = new pdfjsViewer.PDFScriptingManager({
			eventBus,
			sandboxBundleSrc: sandboxBundleSrc
		});

		pdfViewer = new pdfjsViewer.PDFViewer({
			container,
			eventBus,
			linkService: pdfLinkService,
			findController: pdfFindController,
			scriptingManager: pdfScriptingManager
		});

		pdfLinkService.setViewer(pdfViewer);
		pdfScriptingManager.setViewer(pdfViewer);

		eventBus.on('pagesinit', () => {
			pdfViewer.currentScaleValue = 'auto';
			if (searchFor) {
				eventBus.dispatch('find', { type: '', query: searchFor });
			}
		});

		const loadingTask = pdfjs.getDocument({
			url: defaultUrl,
			enableXfa: enableXfa
		});

		const pdfDocument = await loadingTask.promise;
		console.log(pdfDocument);
		pdfViewer.setDocument(pdfDocument);
		pdfLinkService.setDocument(pdfDocument, null);
	});

	const resizePDFViewer = () => {
		const currentScaleValue = pdfViewer.currentScaleValue;
		if (
			currentScaleValue == 'auto' ||
			currentScaleValue == 'page-fit' ||
			currentScaleValue == 'page-width'
		) {
			pdfViewer.currentScaleValue = currentScaleValue;
		}

		pdfViewer.update();
	};

	onMount(() => {
		window.addEventListener('resize', resizePDFViewer);
	});
</script>

<div id="viewerContainer">
	<div id="viewer" class="pdfViewer">
		<div class="page"></div>
	</div>
</div>

<style>
	#viewerContainer {
		overflow: auto;
		width: 50%;
		left: 0;
		right: 0;
		margin-left: auto;
		margin-right: auto;
		position: absolute;
		border-radius: 5px;
		background-color: #eee;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
		max-height: 90vh;
	}

	@media (max-width: 1600px) {
		#viewerContainer {
			width: 75%;
		}
	}

	@media (max-width: 1024px) {
		#viewerContainer {
			width: 90%;
		}
	}
</style>
