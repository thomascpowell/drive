<script lang="ts">
  import { deleteFile } from "$lib/delete";
  import Dialog from "./Dialog.svelte";
  import { get_share_link } from "$lib/get_share_link";
  import { API_URL } from "$lib/utils/config";
  import type { File, Share } from "$lib/utils/types";
  import Copy from "../icons/Copy.svelte";
  import FileIcon from "../icons/FileIcon.svelte";
  import Trash from "../icons/Trash.svelte";
  import { files } from "../stores/files";
  import { status } from "../stores/status";

  export let file_list: File[];

  async function del(fileID: number) {
    let res = await deleteFile(fileID);
    $status = res;
    if (res.message) {
      files.update((current) => current.filter((f: File) => f.ID !== fileID));
    }
  }
  async function copy_link(fileID: number, TTL: number = 60) {
    let res = await get_share_link({
      FileID: fileID,
      TTL: TTL,
    });
    if (res.message) {
      navigator.clipboard.writeText(res.message);
      res.message = "temp share link copied";
    }
    $status = res;
  }
</script>

<!-- TODO: refactor.  -->

<Dialog open={true}>
  <p>test element</p>
</Dialog>

<div class="wrapper">
  <div class="header">Files</div>

  <div class="files">
    {#each file_list as file}
      <div class="line">
        <div>
          <FileIcon />
          <a href={API_URL + "/files/" + file.ID} download> {file.Filename} </a>
        </div>
        <div>
          <p>{file.UploadedAt.substring(5, 10)}</p>
        </div>
        <div>
          <button on:click={() => del(file.ID)}><Trash /></button>
          <button on:click={() => copy_link(file.ID)}><Copy /></button>
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
    padding: 1em 0.5em;
    background-color: var(--bg2);
    display: grid;
    grid-template-columns: 5fr 1fr 1fr;
  }
  .line :first-child {
    min-width: 0;
  }
  .line > :nth-last-child(2) {
    justify-content: flex-start;
  }
  .line > :last-child {
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
    direction: rtl;
    display: flex;
    align-items: center;
    justify-content: center;
    text-decoration: none;
    gap: 1em;
    white-space: nowrap;
    overflow: scroll;
  }
</style>
