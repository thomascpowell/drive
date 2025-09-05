<script lang="ts">
  import "../app.css";
  import Nav from "../components/Nav.svelte";
  import Toast from "../components/Toast.svelte";
</script>

<Nav />
<Toast />
<div class="inner">
  <slot />
</div>
