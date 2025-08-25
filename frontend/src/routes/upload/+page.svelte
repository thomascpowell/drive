<script lang="ts">
  import { upload } from "$lib/upload";
  import type { Res } from "$lib/types";

  let file: File | null = null;
  let files: FileList | null = null;
  $: file = files?.[0] ?? null;

  async function handleSubmit(e: Event) {
    e.preventDefault();
    let res: Res
    if (file == null) {
      return
    }
    res = await upload(file as File);
    console.log(res)
  }
</script>

<form on:submit={handleSubmit}>
  <input type="file" bind:files={files} />
  <button type="submit" disabled={!file}>Upload</button>
</form>
