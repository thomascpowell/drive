<script lang="ts">
  import type { Res } from "$lib/types";
  import { upload } from "$lib/upload";
  import ArrowRight from "../icons/ArrowRight.svelte";
  import Upload from "../icons/Upload.svelte";
  import FileIcon from "../icons/FileIcon.svelte";

  let fileInput: HTMLInputElement;
  $: filename = file?.name ?? "no file selected";
  let file: File | null = null;

  function handleChange(e: Event) {
    const target = e.target as HTMLInputElement;
    let selected: File | null = target.files?.[0] ?? null;
    console.log(target.files?.[0] ?? null);
    file = selected;
  }

  async function handleSubmit(e: Event) {
    e.preventDefault();
    let res: Res;
    if (file == null) {
      return;
    }
    res = await upload(file as File);
    console.log(res);
  }

  async function handleClick() {
    fileInput.click();
  }
</script>

<form on:submit={handleSubmit}>
  <div class="file">
    <input type="file" on:change={handleChange} bind:this={fileInput} />
    <button on:click={handleClick} type="button">
      {#if file == null}
        <Upload />
      {:else}
        <FileIcon />
      {/if}
    </button>
    <span>{filename}</span>
  </div>
  <button type="submit" disabled={!file}>upload <ArrowRight /></button>
</form>

<style>
  .file {
    border-radius: 0.25em;
    background-color: var(--bg3);
    border: 0.1em solid var(--border);
  }
  .file button {
    width: 100%;
    height: 10em;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    color: var(--text);
  }
  .file > input {
    width: 100%;
    display: none;
  }
  .file > span {
    display: flex;
    align-items: center;
    padding: 0.25em 0.5em;
    color: var(--text2);
    border-top: 0.1em solid var(--border);
    width: 100%;
    overflow: hidden;
  }

  form {
    width: 20em;
    background-color: var(--bg2);
    border: 0.1em solid var(--border);
    border-radius: 0.5em;
    padding: 1em;
    display: flex;
    flex-direction: column;
    gap: 1em;
  }

  form > button {
    color: var(--text);
    padding: 0.25em 0.5em;
    border-radius: 0.25em;
    background-color: var(--bg3);
    border: 0.1em solid var(--border);
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
</style>
