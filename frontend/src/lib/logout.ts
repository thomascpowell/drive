import { goto } from '$app/navigation'
import { API_URL } from './config'

export async function logout() {
  const endpoint = `${API_URL}/logout`
  await fetch(endpoint, {
    method: 'post',
    headers: {
      'content-type': 'application/json'
    },
  })
  goto("/")
}
