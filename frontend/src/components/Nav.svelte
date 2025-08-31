<script>
  import NavLink from "./NavLink.svelte";
</script>

<div>
  <NavLink href="/files">files</NavLink>
  <NavLink href="/upload">upload</NavLink>
  <NavLink href="/logout">logout</NavLink>
</div>

<style>
  div {
    position: absolute;
    display: flex;
    align-items: center;
    gap: 1em;
  }
</style>
