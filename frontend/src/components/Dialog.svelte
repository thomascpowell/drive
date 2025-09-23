<script lang="ts">
  import Close from "../icons/Close.svelte";
  export let title: string;
  export let open: boolean;
</script>

<div class="wrapper" role="none" on:click={() => (open = false)}>
  <div class="inner" on:click|stopPropagation role="none">
    <div class="header">
      <h2>{title}</h2>
      <button on:click={() => (open = false)}>
        <Close />
      </button>
    </div>
    <div class="content">
      <slot />
    </div>
  </div>
</div>

<style>
  .wrapper {
    backdrop-filter: brightness(90%) blur(2px);
    position: absolute;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1em;
    width: 100vw;
    height: 100vh;
    z-index: 5;
  }
  .inner {
    width: 100%;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    height: min-content;
    max-height: 20em;
    min-height: 10em;
    max-width: 25em;
    overflow: scroll;
    background-color: var(--bg2);
    border-radius: 1em;
    border: 0.1em solid var(--border);
  }

  .header {
    padding: 1.5em 1.5em 0em 1.5em;
    display: flex;
    gap: 2em;
    align-items: center;
    width: 100%;
    justify-content: space-between;
  }
  .header > h2 {
    display: flex;
    align-items: center;
    white-space: nowrap;
    overflow: scroll;
  }
  .header button {
    display: flex;
    align-items: center;
  }
  .content {
    justify-content: baseline;
    display: flex;
    gap: 1em;
    flex-direction: column;
    padding: 1.5em;
    height: 100%;
    width: 100%;
  }
</style>
