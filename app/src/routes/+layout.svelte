<script>
  import "../app.css";

  import Header from "../components/Header.svelte";
</script>

<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
  <Header />
  <slot />
</div>
