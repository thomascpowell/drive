import { API_URL } from '$lib/config'

export async function checkHealth() {
  try {
    const res = await fetch(`${API_URL}/health`, {
      method: "GET",
    });
    const data = await res.json()
    return data
  } catch (err) {
    return { status: "error" }
  }
}
