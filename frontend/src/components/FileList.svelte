<script lang="ts">
  import { deleteFile } from "$lib/delete";
  import { get_share_link } from "$lib/get_share_link";
  import { API_URL } from "$lib/utils/config";
  import type { File, Share } from "$lib/utils/types";
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
  async function copy_link(fileID: number) {
    // TODO: change to somthing either longer or user supplied
    let TTL = 30;
    let req: Share = {
      FileID: fileID,
      TTL: TTL,
    };
    let res = await get_share_link(req);
    $status = res;
    if (res.message) {
      navigator.clipboard.writeText(res.message);
    }
  }
</script>


<!-- TODO: refactor.  -->

<div class="wrapper">
  <div class="header">
    Files
  </div>

  <div class="files">
    {#each file_list as file}
      <div class="line">
        <div>
          <FileIcon style="transform: scale(0.6)" />
          <a href={API_URL + "/files/" + file.ID} download> {file.Filename} </a>
        </div>
        <div>
          <p>{file.UploadedAt.substring(5, 10)}</p>
        </div>
        <div>
          <button on:click={() => del(file.ID)}
            ><Trash style="transform: scale(0.6)" /></button
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
  }
  .header {
    border-bottom: 0.1em solid var(--border);
    border-radius: 0.4em 0.4em 0 0;
    background-color: var(--bg3);
    border-bottom: 0.1e;
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
    width: 100%;
    display: flex;
    gap: 2em;
    border-bottom: 0.1em solid var(--border);
    justify-content: space-between;
    padding: 0.5em;
    background-color: var(--bg2);
    display: grid;
    grid-template-columns: 5fr 1fr 1fr;
  }
  .line > * {
    display: flex;
    align-items: center;
    text-align: center;
    min-width: 0;
    white-space: nowrap;
  }
  .line > :last-child {
    justify-content: flex-end;
  }
  .line > :nth-last-child(2) {
    justify-content: flex-start;
  }

  a,
  p,
  button {
    direction: rtl;
    display: flex;
    align-items: center;
    text-decoration: none;
    max-width: 10em;
    white-space: nowrap;
  }
</style>
