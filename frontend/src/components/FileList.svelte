<script lang="ts">
  import { deleteFile } from "$lib/delete";
  import Dialog from "./Dialog.svelte";
  import { get_share_link } from "$lib/get_share_link";
  import { API_URL } from "$lib/utils/config";
  import type { FileRec } from "$lib/utils/types";
  import FileIcon from "../icons/FileIcon.svelte";
  import Trash from "../icons/Trash.svelte";
  import { files } from "../stores/files";
  import { status } from "../stores/status";
  import LinkIcon from "../icons/LinkIcon.svelte";
  import DotsIcon from "../icons/DotsIcon.svelte";
  import Download from "../icons/Download.svelte";

  export let file_list: FileRec[];

  let current_file: FileRec | null = null;
  let dialog_open: boolean = false;
  let dialog_ttl: number = 60;

  async function del(fileID: number) {
    let res = await deleteFile(fileID);
    $status = res;
    if (res.message) {
      files.update((current) =>
        current.filter((f: FileRec) => f.ID !== fileID),
      );
    }
  }
  async function copy_link(fileID: number, TTL: number) {
    let res = await get_share_link({
      FileID: fileID,
      TTL: Number(TTL),
    });
    if (res.message) {
      navigator.clipboard.writeText(res.message);
      res.message = "temp share link copied";
    }
    $status = res;
  }
</script>

{#if dialog_open && current_file}
  <Dialog bind:open={dialog_open} title={current_file!.Filename}>
    <a href={API_URL + "/files/" + current_file!.ID} download>
      <Download /> download
    </a>
    <button on:click={() => del(current_file!.ID)}>
      <Trash /> delete
    </button>
    <div class="ttl_form">
      <button on:click={() => copy_link(current_file!.ID, dialog_ttl)}>
        <LinkIcon />get link
      </button>
       w/ TTL
      <input bind:value={dialog_ttl} type="text">
      (s)
    </div>
  </Dialog>
{/if}

<div class="wrapper">
  <div class="header">
    <h1>Files</h1>
  </div>
  <div class="files">
    {#each file_list as file}
      <div class="line">
        <div class="name">
          <FileIcon />
          <p>{file.Filename}</p>
        </div>

        <p class="date">{file.UploadedAt.substring(5, 10)}</p>

        <div class="actions">
          <button on:click={() => ((current_file = file), (dialog_open = true))}
            ><DotsIcon /></button
          >
        </div>
      </div>
    {/each}
  </div>
</div>

<style>
  .wrapper {
    border: 0.1em solid var(--border);
    border-radius: 0.5em;
    max-width: 30em;
    width: 100%;
  }
  .header {
    border-bottom: 0.1em solid var(--border);
    border-radius: 0.4em 0.4em 0 0;
    background-color: var(--bg3);
    padding: 1em;
  }
  .wrapper :last-child {
    border-radius: 0 0 0.4em 0.4em !important;
    border: none !important;
  }

  .files {
    max-height: 20em;
    overflow: scroll;
  }

  .line {
    flex-shrink: 0;
    width: 100%;
    display: flex;
    gap: 1em;
    border-bottom: 0.1em solid var(--border);
    justify-content: space-between;
    padding: 1em 1em;
    background-color: var(--bg2);
    display: grid;
    grid-template-columns: 1fr auto auto;
  }
  .name {
    min-width: 0;
    white-space: nowrap;
    overflow: scroll;
  }
  .date {
    flex-shrink: 0;
    width: min-content;
    justify-content: flex-start;
  }
  .actions {
    justify-content: flex-end;
  }
  .line > div {
    display: flex;
    align-items: center;
    text-align: center;
    white-space: nowrap;
    gap: 0.5em;
  }

  a,
  p,
  button {
    width: min-content;
    white-space: nowrap;
    display: flex;
    align-items: center;
    justify-content: center;
    text-decoration: none;
    gap: 1em;
  }

  .ttl_form {
    display: flex;
    align-items: center;
    gap: 0.5em;
  }
  .ttl_form input {
    width: 2ch;
    max-width: 3em;
    border-bottom: 0.025em solid var(--text);
  }
</style>
