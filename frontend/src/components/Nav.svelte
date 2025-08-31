<script>
  import { page } from "$app/state";
  import NavLink from "./NavLink.svelte";
</script>

{#if page.url.pathname !== "/"}
  <div>
    <NavLink href="/logout">auth</NavLink>
    <NavLink href="/files">files</NavLink>
    <NavLink href="/upload">upload</NavLink>
  </div>
{/if}

<style>
  div {
    position: absolute;
    display: flex;
    align-items: center;
    gap: 1em;
  }
</style>
