<script>
  import dayjs from "dayjs";

  export let film;

  const ts = dayjs.unix(film.date);
  const date = ts.format("DD MMM YYYY");
  const time = ts.format("HH:mm");
</script>

<a class="flex flex-row max-w-lg py-2 gap-4" href={film.link} target="_blank">
  <img class="max-w-32 rounded-md" alt="poster" src={film.img_url} />
  <div class="flex flex-col gap-1">
    <div class="font-bold">{film.title}</div>
    <div>🕰️ {time} // {date}</div>
    <div>📍 {film.location}</div>
    <div>💵 {film.price}</div>
    <!-- <a href={film.link} class="hover:underline"> -->
    <!-- 🎫 <span class="italic ">cumpara bilet</span> -->
    <!-- </a> -->
  </div>
</a>
