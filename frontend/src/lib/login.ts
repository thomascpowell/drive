import { API_URL } from '$lib/utils/config'
import type { Res } from '$lib/utils/types'

export async function login(username: string, password: string): Promise<Res> {
  const creds = JSON.stringify({ username, password })
  const endpoint = `${API_URL}/login`

  const res = await fetch(endpoint, {
    method: 'post',
    headers: {
      'content-type': 'application/json'
    },
    body: creds,
  })

 
  const data = await res.json();
  return data
}
