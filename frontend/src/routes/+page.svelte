<script lang="ts">
  import { onMount } from "svelte";
  import { checkHealth } from "$lib/health";

  let health = "unknown";

  onMount(async () => {
    const result = await checkHealth();
    health = result.status;
  });
</script>

<p>health: {health}</p>
