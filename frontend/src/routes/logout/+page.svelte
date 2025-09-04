<script lang="ts">
  import { logout } from "$lib/logout";
  import { onMount } from "svelte";
  import { status } from "../../stores/status";

  onMount(async () => {
    $status = await logout();
  });
</script>

<p>...</p>
