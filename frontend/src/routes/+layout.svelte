<div>
  <a href="/">home</a>
  <a href="/login">login</a>
  <a href="/register">register</a>
  <a href="/logout">logout</a>
  <a href="/upload">upload</a>
  <a href="/files">files</a>
</div>
<slot />

<style>
  div {
    display: flex;  
    align-items: center;
    gap: 1em;
  }
</style>
