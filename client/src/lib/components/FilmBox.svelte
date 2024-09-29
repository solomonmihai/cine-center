<script>
  import dayjs from "dayjs";

  export let film;

  // todo: maybe hardcoding for this time zone is not the best thing to do
  const ts = dayjs.unix(film.date);
  const date = ts.format("DD MMM YYYY");
  const time = ts.format("HH:mm");
</script>

<a
  class="group flex flex-row max-w-lg p-2 gap-4 hover:border-black border-white border-[1px]"
  href={film.link}
  target="_blank"
>
  <img class="max-w-32 rounded-md" alt="poster" src={film.img_url} />
  <div class="flex flex-col gap-1">
    <div class="font-bold group-hover:underline">{film.title}</div>
    <div><span class="w-1">🕰️</span> {time} // {date}</div>
    <div><span class="w-1">📍</span> {film.location}</div>
    <div><span class="w-1">💵</span> {film.price}</div>
    <!-- <a href={film.link} class="hover:underline"> -->
    <!-- 🎫 <span class="italic ">cumpara bilet</span> -->
    <!-- </a> -->
  </div>
</a>
