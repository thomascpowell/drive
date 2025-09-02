<script lang="ts">
  import ArrowRight from "../icons/ArrowRight.svelte";
  import { login } from "$lib/login";
  import { register } from "$lib/register";
  import { status } from "../stores/status";
  import { goto } from "$app/navigation";

  // Terrible, but enums are not supported in svelte
  const LOGIN = 0;
  const REGISTER = 1;

  let mode: number;
  $: mode = LOGIN;

  let username: string;
  $: username = "";

  let password: string;
  $: password = "";

  async function handleSubmit(e: SubmitEvent) {
    e.preventDefault();
    if (mode == REGISTER) {
      $status = await register(username, password);
      return;
    }

    $status = await login(username, password);
    if ($status.message) {
      goto("/files")
    }
    return;
  }
</script>

<div class="wrapper">
  <div class="mode_buttons">
    <button
      onclick={() => (mode = LOGIN)}
      class={mode == LOGIN ? "selected" : ""}>login</button
    >
    <button
      onclick={() => (mode = REGISTER)}
      class={mode == REGISTER ? "selected" : ""}>register</button
    >
  </div>

  <form onsubmit={handleSubmit} action="submit">
    <input type="text" bind:value={username} placeholder="username" />
    <input type="text" bind:value={password} placeholder="password" />
    <button>{mode == LOGIN ? "login" : "register"} <ArrowRight /></button>
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
    width: 20em;
    border-radius: 0.5em;
  }

  form {
    display: flex;
    max-width: 100%;
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
