<script lang="ts">
    import { fade } from "svelte/transition";
  import Check from "../icons/Check.svelte";
  import Warn from "../icons/Warn.svelte";
  import { status } from "../stores/status";

  const SHOW_TIME = 2000;
  $: Icon = $status.error ? Warn : Check;
  let timeout = setTimeout(() => {});
  let visible: boolean = false;
  let desc: string = "";

  $: if ($status != null) {
    visible = true;
    clearTimeout(timeout);
    desc = $status.message || $status.error || "";
    timeout = setTimeout(() => {
      visible = false;
    }, SHOW_TIME);
  }
</script>

{#if visible && desc}
  <div transition:fade={{ duration: 150 }}>
    <Icon />
    <p>{desc}</p>
  </div>
{/if}

<style>
  div {
    background-color: var(--bg2);
    padding: 1em;
    margin: 1em;
    border-radius: 0.5em;
    border: 0.1em solid var(--border);
    display: flex;
    align-items: center;
    gap: 1em;
    position: absolute;
    bottom: 0;
    right: 0;

    transition: fade;
  }
</style>
