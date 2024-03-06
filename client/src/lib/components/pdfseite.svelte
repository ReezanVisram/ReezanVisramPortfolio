<script>
	import { onMount, tick } from 'svelte';

	export let canvasseite;
	let seite = canvasseite[0];
	let viewport = canvasseite[1];
	let canvas;
	let geladen = false;

	function anzeigen() {
		console.log('anzeigen gestartet');

		const canvasContext = canvas.getContext('2d');
		console.log('kontext gesetzt');

		canvas.height = viewport.height;
		canvas.width = viewport.width;
		console.log('Canvas dimensioniert: ' + canvas.height + '/' + canvas.width);

		let renderContext = {
			canvasContext,
			viewport
		};
		console.log('Renderkontext');

		let renderTask = seite.render(renderContext);
		console.log('Rendertask');
		renderTask.promise.then(function () {
			console.log('Canvasseite gerendert');
			geladen = true;
		});
	}

	onMount(anzeigen);
</script>

<div class="wrupper">
	<canvas bind:this={canvas} width="500px" height="500px" />
</div>

<style>
	canvas {
		width: 100%;
	}
	.wrupper {
		padding: 8px;
	}
</style>
