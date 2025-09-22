<script lang="ts">
    import Close from "../icons/Close.svelte";

  export let open: boolean;
  export let title: string;
</script>

{#if open}
  <div class="wrapper">
    <div class="inner">

      <div class="header">
        <p>{title}</p>
        <button on:click={() => (open = false)}>
          <Close/>
        </button>
      </div>

     <div class="content">
        <slot />
      </div>

    </div>
  </div>
{/if}

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
    justify-content: space-between;
    height: 20em;
    max-width: 30em;
    background-color: var(--bg2);
    border-radius: 1em;
    border: 0.1em solid var(--border);
  }

  .header {
    padding: 1em;
    display: flex;
    align-items: center;
    width: 100%;
    justify-content: space-between;
  }
  .content {
    justify-content: baseline;
    padding: 1em;
    height: 100%;
    width: 100%;

  }
</style>
