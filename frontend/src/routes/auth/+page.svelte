<script lang="ts">
  import { login } from "$lib/login";
  import type { Res } from "$lib/types"

  let username = "";
  let password = "";
  let res:Res; //TODO: Use this w some sort of status component

  async function handleSubmit(e: Event) {
    e.preventDefault();
    res = await login(username, password);
    console.log(res)
  }
</script>

<p>register</p>

<form on:submit|preventDefault={handleSubmit} action="submit">
  <input type="text" bind:value={username} placeholder="username" />
  <input type="text" bind:value={password} placeholder="password" />
  <button>submit</button>
</form>
