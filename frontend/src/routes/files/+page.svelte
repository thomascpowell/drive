<script lang="ts">
  import { onMount } from "svelte";
  import { get_files } from "$lib/get_files";
  import type { File } from "$lib/types";
import { API_URL } from '$lib/config'

  let files: File[];
  onMount(async () => {
    files = await get_files();
    console.log(files);
  });
</script>

<div>
  {#each files as file}
    <a href={API_URL + "/files/" + file.ID} download>{file.Filename}</a>
  {/each}
</div>

<style>
  div {
    display: flex;
    gap: 1em;
    flex-direction: column;
  }
</style>
