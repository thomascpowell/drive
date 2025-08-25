<script lang="ts">
  import { onMount } from "svelte";
  import { get_files } from "$lib/get_files"
  import type { File } from "$lib/types"

  let files: File[]
  onMount(async () => {
    files = await get_files();
    console.log(files)
  });
</script>

{#each files as file}
  <p>{file.Filename}</p>
{/each}
