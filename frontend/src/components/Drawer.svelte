<script lang="ts">
  export let open: boolean;
</script>

{#if open}
  <div class="wrapper">
    <div class="inner">
      <slot/>
    </div>
  </div>
{/if}

<style>
  .wrapper {
    position: absolute;
    display: flex;
    padding: 1em;
    width: 100%;
    bottom: 0;
    z-index: 5;
  }
  .inner {
    display: flex;
    flex-direction: column;
    gap: 1em;
    padding: 1em;
    width: 100%;
    height: 20em;
    max-width: 30em;
    background-color: var(--bg2);
  }

</style>
