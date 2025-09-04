<script lang="ts">
  import Check from "../icons/Check.svelte";
  import Warn from "../icons/Warn.svelte";
  import { status } from "../stores/status";

  const SHOW_TIME = 3000;
  $: Icon = $status.error ? Warn : Check;

  let timeout = setTimeout(() => {});
  let visible: boolean = false;
  let desc: string = "";

  $: if ($status != null) {
    visible = true;
    clearTimeout(timeout);
    desc = $status.message || $status.error || ""

    timeout = setTimeout(() => {
      visible = false;
    }, SHOW_TIME);
  }
</script>

{#if visible}
  <div>
    <Icon />
    <p>{desc}</p>
  </div>
{/if}
