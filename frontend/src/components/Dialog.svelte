<script lang="ts">
  export let open: boolean;
</script>

{#if open}
  <div class="wrapper">
    <div class="inner">
      <div class="close">
        <button on:click={() => open = false}>X</button>
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
    padding: 1em;

    justify-content: center;
    align-items: end;
    width: 100%;
    height: 100%;
    bottom: 0;
    z-index: 5;
  }
  .inner {
    position: relative;
    width: 100%;
    height: 20em;
    max-width: 30em;
    background-color: var(--bg2);
    border-radius: 1em;
    border: 0.1em solid var(--border);
  }

  .close {
    position: absolute;
    top: 0;
    right: 0;
    padding: 1em;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .content {
    padding: 1em;
    display: flex;
    flex-direction: column;
    gap: 1em;
  }
</style>
