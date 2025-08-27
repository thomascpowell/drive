<script lang="ts">
  import ArrowRight from "../icons/ArrowRight.svelte";

  // Terrible, but enums are not supported in svelte
  const LOGIN = 0;
  const REGISTER = 1;

  let mode: number;
  $: mode = LOGIN;

  let username: string;
  $: username = "";

  let password: string;
  $: password = "";

  function handleSubmit(e: SubmitEvent) {
    e.preventDefault();
    if (mode == LOGIN) {
      //TODO
    }

    if (mode == LOGIN) {
      //TODO
    }
  }
</script>

<!-- going to do some sort of tabbed auth thing -->

<div class="wrapper">
  <div class="mode_buttons">
    <button
      onclick={() => (mode = LOGIN)}
      class={mode == LOGIN ? "selected" : ""}>Login</button
    >
    <button
      onclick={() => (mode = REGISTER)}
      class={mode == REGISTER ? "selected" : ""}>Register</button
    >
  </div>

  <form onsubmit={handleSubmit} action="submit">
    <input type="text" bind:value={username} placeholder="Username" />
    <input type="text" bind:value={password} placeholder="Password" />
    <button>{mode == LOGIN ? "Login" : "Register"} <ArrowRight /></button>
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
    border-radius: 0.5em;
  }

  form {
    display: flex;
    flex-direction: column;
    gap: 1em;
  }
  form > input {
    color: var(--text);
    padding: 0.25em 0.5em;
    border-radius: 0.25em;
    background-color: var(--bg3);
    border: 0.1em solid var(--border);
  }
  form > button {
    color: var(--text);
    padding: 0.25em 0.5em;
    border-radius: 0.25em;
    background-color: var(--bg3);
    border: 0.1em solid var(--border);
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .mode_buttons {
    display: flex;
  }
  .mode_buttons > button {
    background-color: var(--bg2);
    border: 0.1em solid var(--border);
    color: var(--text);
    padding: 0.25em 0.5em;
    border-radius: 0.5em 0 0 0.5em;
  }
  .mode_buttons :last-child {
    border-radius: 0 0.5em 0.5em 0;
    border-left: none;
  }
  .mode_buttons .selected {
    background-color: var(--bg3);
  }
</style>
