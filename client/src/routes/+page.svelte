<script>
  import Container from "$lib/components/Container.svelte";
  import FilmList from "$lib/components/FilmList.svelte";

  export let data;
</script>

<Container>
  <h1 class="block text-2xl italic font-bold">Cine Centru</h1>
  <div class="font-bold">cinemauri</div>
  <div class="w-full">
    <div class="flex flex-col items-start w-full">
      {#each data.cinemas as cinema}
        <a class="block italic hover:underline" href={`/cinema/${encodeURIComponent(cinema)}`}
          >{cinema}</a
        >
      {/each}
    </div>
  </div>

  <div class="font-bold">toate</div>
  <FilmList films={data.films} />
</Container>
