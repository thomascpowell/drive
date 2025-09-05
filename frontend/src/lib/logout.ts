import { goto } from '$app/navigation'
import { API_URL } from './utils/config'
import type { Res } from './utils/types'

export async function logout(): Promise<Res> {
  const endpoint = `${API_URL}/logout`
  const res = await fetch(endpoint, {
    method: 'post',
    headers: {
      'content-type': 'application/json'
    },
  })
  const data = await res.json()
  goto("/")
  return data
}
