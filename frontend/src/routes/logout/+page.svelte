<script lang="ts">
  import { logout } from "$lib/logout";
  import { onMount } from "svelte";
  onMount(async () => {
    logout()
  });
</script>

<p>...</p>
