<script lang="ts">
  import { has_cookie } from "$lib/has_cookie";

  // Terrible, but enums are not supported in svelte
  const LOGIN = 0;
  const REGISTER = 1;

  let mode: number;
  $: mode = LOGIN;

  let username: string;
  $: username = "";

  let password: string;
  $: password = "";

  function handleSubmit() {
    if (mode == LOGIN) {
      //TODO
    }

    if (mode == LOGIN) {
      //TODO
    }
  }
</script>

<!-- going to do some sort of tabbed auth thing -->
<!-- move this to a component after dev -->

<div class="wrapper">
  <div class="buttons">
    <button onclick={() => mode = LOGIN} class={mode == LOGIN ? 'selected' : ''}>login</button>
    <button onclick={() => mode = REGISTER} class={mode == REGISTER ? 'selected' : ''}>register</button>
  </div>

  <form onsubmit={handleSubmit} action="submit">
    <input type="text" bind:value={username} placeholder="username" />
    <input type="text" bind:value={password} placeholder="password" />
    <button>submit</button>
  </form>
</div>

<style>
  .wrapper {
    display: flex;
    flex-direction: column;
    gap: 1em;
    padding: 1em;
    width: min-content;
    background-color: var(--bg2);
    border: 0.1em solid var(--border);
    border-radius: 1em;
  }

  form {
    display: flex;
    flex-direction: column;
    gap: 0.5em;
  }
  form > input {
    border-bottom: 0.1em solid var(--text);
  }

  form > button {
    color: var(--text);
    padding: 0.25em 0.5em;
    border-radius: 0.25em;
    margin-top: 0.5em;
    background-color: var(--bg3);
    border: 0.1em solid var(--border);
  }

  .buttons {
    display: flex;
  }
  .buttons > button {
    background-color: var(--bg);
    border: 0.1em solid var(--border);
    color: var(--text);
    padding: 0.25em 0.5em;
    border-radius: 0.5em 0 0 0.5em;
  }
  .buttons :last-child {
    border-radius: 0 0.5em 0.5em 0;
    border-left: none;
  }
  .buttons .selected {
    background-color: var(--bg3);
  }
</style>
