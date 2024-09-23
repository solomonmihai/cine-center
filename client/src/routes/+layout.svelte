<script>
  import "../app.css";

  function goToTop() {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
  }

  let visible = true;
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
