<script>
  import NavLink from "./NavLink.svelte";
</script>

<div>
  <NavLink href="/">home</NavLink>
  <NavLink href="/upload">upload</NavLink>
  <NavLink href="/files">files</NavLink>
  <NavLink href="/logout">logout</NavLink>
</div>
<slot />

<style>
  div {
    position: absolute;
    display: flex;
    align-items: center;
    gap: 1em;
  }
</style>
