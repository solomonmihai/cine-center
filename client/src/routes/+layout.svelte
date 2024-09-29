<script>
  import { onMount } from "svelte";
  import "../app.css";

  function goToTop() {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
  }

  let visible = false;

  onMount(() => {
    const onScroll = () => {
      visible = window.scrollY > 0;
    };

    document.addEventListener("scroll", onScroll);

    return () => {
      document.removeEventListener("scroll", onScroll);
    };
  });
</script>

<div>
  <div
    class={`fixed top-0 px-3 py-0 text-white bg-black cursor-pointer right-4 ${
      !visible && "hidden"
    }`}
    on:click={goToTop}
  >
    ^sus
  </div>
  <slot />
</div>
