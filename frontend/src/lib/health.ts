import { API_URL } from '$lib/config'
import type { Res } from '$lib/types'

export async function checkHealth(): Promise<Res> {
  try {
    const res = await fetch(`${API_URL}/health`, {
      method: "GET",
    });
    const data = await res.json()
    return data
  } catch (err) {
    return { message: "error" }
  }
}
