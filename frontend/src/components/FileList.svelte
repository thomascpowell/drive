<script lang="ts">
  import { API_URL } from "$lib/config";
  import type { File } from "$lib/types";
  import Download from "../icons/Download.svelte";
  import FileIcon from "../icons/FileIcon.svelte";
  import Trash from "../icons/Trash.svelte";

  export let files: File[];
</script>

<div class="wrapper">
  <div class="line">
    <div><p>name</p></div>
    <div><p>date</p></div>
    <div><p>del</p></div>
  </div>

  {#each files as file}
    <div class="line">
      <div>
        <FileIcon style="transform: scale(0.6)" />
        <a href={API_URL + "/files/" + file.ID} download> {file.Filename} </a>
      </div>
      <div>
        <p>{file.UploadedAt.substring(5, 10)}</p>
      </div>
      <div>
        <a href="/todo"><Trash style="transform: scale(0.6)" /></a>
      </div>
    </div>
  {/each}
</div>

<style>
  .line {
    display: flex;
    gap: 2em;
    border-bottom: 0.1em solid var(--border);
    justify-content: space-between;
    padding: 0.5em;
    background-color: var(--bg2);
    display: grid;
    grid-template-columns: 9fr 2fr 1fr;
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
  p {
    direction: rtl;
    display: flex;
    align-items: center;
    text-decoration: none;
    overflow: hidden!important;
    max-width: 10em;
    white-space: nowrap;
  }

  .wrapper {
    border: 0.1em solid var(--border);
    border-radius: 0.5em;
  }
  .wrapper > :first-child {
    border-radius: 0.4em 0.4em 0 0;
    background-color: var(--bg3);
    padding: 1em;
  }
  .wrapper :nth-last-child(1) {
    border-radius: 0 0 0.4em 0.4em !important;
    border: none !important;
  }
</style>
