<script lang="ts">
  import { get_file } from "$lib/get_file";
  import type { File } from "$lib/types";
  import { onMount } from "svelte";

  let file: File;
  export let data;

  onMount(async () => {
    const id = data.id;
    file = await get_file(id);
    console.log(file);
  });
</script>
