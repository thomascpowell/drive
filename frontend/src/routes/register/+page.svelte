<script lang="ts">
  import { register } from "$lib/register";
  import type { Res } from "$lib/types"

  let username = "";
  let password = "";
  let res:Res; //TODO: Use this w some sort of status component

  async function handleSubmit(e: Event) {
    e.preventDefault();
    res = await register(username, password);
    console.log(res)
  }
</script>

<p>register</p>

<form on:submit|preventDefault={handleSubmit} action="submit">
  <input type="text" bind:value={username} placeholder="username" />
  <input type="text" bind:value={password} placeholder="password" />
  <button>submit</button>
</form>
