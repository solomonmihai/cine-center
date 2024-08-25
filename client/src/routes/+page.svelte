<script>
  import Container from "$lib/components/Container.svelte";
  import FilmList from "$lib/components/FilmList.svelte";

  export let data;
  console.log(data);
</script>

<Container>
  <h1 class="block text-2xl italic font-bold">Cine Center</h1>
  <div class="w-full">
    <div class="font-bold">vezi per cinema</div>
    <div class="flex flex-col items-start w-full">
      {#each data.cinemas as cinema}
        <a class="block italic hover:underline" href={`/cinema/${encodeURIComponent(cinema)}`}
          >{cinema}</a
        >
      {/each}
    </div>
  </div>

  <FilmList films={data.films} />
</Container>
