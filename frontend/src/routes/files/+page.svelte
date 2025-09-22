<script lang="ts">
  import { onMount } from "svelte";
  import { get_files } from "$lib/get_files";
  import { files } from "../../stores/files";
  import FileList from "../../components/FileList.svelte";
  import Drawer from "../../components/Drawer.svelte";

  onMount(async () => {
    $files = await get_files();
  });
</script>

<Drawer open={false}>hello</Drawer>

<div>
  {#if $files.length == 0}
    <p>no files</p>
  {:else}
    <FileList file_list={$files} />
  {/if}
</div>

<style>
  div {
    display: flex;
    gap: 1em;
    flex-direction: column;
  }
</style>
