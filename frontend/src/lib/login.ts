import { API_URL } from '$lib/config'
import type { Res } from '$lib/types'

export async function login(username: string, password: string): Promise<Res> {
  const creds = JSON.stringify({ username, password })
  const endpoint = `${API_URL}/auth`

  const res = await fetch(endpoint, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: creds,
  })

  const data = await res.json();
  return data
}
