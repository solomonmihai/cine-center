<script>
  import FilmBox from "$lib/components/FilmBox.svelte";
  export let films;
</script>

<div>
  {#if films.length > 0}
    {#each films as film}
      <FilmBox {film} />
    {/each}
  {:else}
    <div class="italic">nu sunt filme</div>
  {/if}
</div>
