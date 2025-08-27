<script lang="ts">
  import { page } from "$app/state";
  export let href: string;
</script>

<a {href} class:active={page.url.pathname === href}>
  <slot />
</a>

<style>
  a {
    text-decoration: none;
  }

  .active {
    text-decoration: underline;
  }
</style>
