<script lang="ts">
  import { onMount } from "svelte";
  import { get_files } from "$lib/get_files";
  import type { File } from "$lib/utils/types";
  import FileList from "../../components/FileList.svelte";

  let files: File[] = [];

  onMount(async () => {
    files = await get_files();
    console.log(files);
  });
</script>

<div>
  {#if files.length == 0}
    <p>no files</p>
  {:else}
    <FileList {files} />
  {/if}
</div>

<style>
  div {
    display: flex;
    gap: 1em;
    flex-direction: column;
  }
</style>
