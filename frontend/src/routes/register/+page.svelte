<script lang="ts">
  import { register } from "$lib/register";

  let username = "";
  let password = "";
  let res; //TODO: Use this w some sort of status component

  async function handleSubmit(e: Event) {
    e.preventDefault();
    res = register(username, password);
  }
</script>

<p>register</p>

<form on:submit|preventDefault={handleSubmit} action="submit">
  <input type="text" placeholder="username" />
  <input type="text" placeholder="password" />
  <button>submit</button>
</form>
