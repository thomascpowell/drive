<script lang="ts">
  import { onMount } from "svelte";
  import { checkHealth } from "$lib/health";
  import type { Res } from "$lib/types";

  let health = "unknown";

  onMount(async () => {
    const result: Res = await checkHealth();
    health = result.message || "invalid res";
  });
</script>

<p>health: {health}</p>
