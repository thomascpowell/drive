<script lang="ts">
  import Close from "../icons/Close.svelte";
  export let title: string;
  export let open: boolean;
</script>

<div class="wrapper">
  <div class="inner">
    <div class="header">
      <p>{title}</p>
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
    max-width: 20em;
    height: min-content;
    max-height: 20em;
    min-height: 10em;
    overflow: scroll;
    background-color: var(--bg2);
    border-radius: 1em;
    border: 0.1em solid var(--border);
  }

  .header {
    padding: 1em;
    display: flex;
    gap: 2em;
    align-items: center;
    width: 100%;
    justify-content: space-between;
  }
  .header > p {
    display: flex;
    align-items: center;
    white-space: nowrap;
    overflow: scroll;
  }
  .content {
    justify-content: baseline;
    padding: 1em;
    height: 100%;
    width: 100%;
  }
</style>
