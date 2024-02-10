<script>
  import Header from "$lib/header.svelte";
</script>

<Header />

<style>
  @font-face {
    src: url("./Outfit.ttf");
    font-family: Outfit;
  }

  :global(:root) {
    --background-colour: #FFFFFF;
    --text-primary-colour: #0065FF;
    --text-secondary-colour: #000000;
    --button-text-colour: #FFFFFF;
  }

  :global(*) {
    margin: 0;
    padding: 0;
    font-family: Outfit, 'sans-serif';
  }

  :global(body) {
    margin-left: 12.5vw;
    margin-right: 12.5vw;
  }
</style>