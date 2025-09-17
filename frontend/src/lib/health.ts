import { API_URL } from '$lib/utils/config'
import type { Res } from '$lib/utils/types'

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
