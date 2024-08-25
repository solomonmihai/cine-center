<script>
  import Container from "$lib/components/Container.svelte";
  import FilmList from "$lib/components/FilmList.svelte";

  export let data;
</script>

<Container>
  <h1 class="block text-2xl italic font-bold">{data.name}</h1>
  <a class="block italic hover:underline" href="/">{'<'} inapoi</a>
  <FilmList films={data.films} />
</Container>
