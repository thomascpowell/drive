import { API_URL } from '$lib/utils/config'
import type { Res } from '$lib/utils/types'

export async function deleteFile(fileID: number): Promise<Res> {
  const endpoint = `${API_URL}/files/${fileID}`
  const res = await fetch(endpoint, {
    method: 'delete',
    headers: {
      'content-type': 'application/json'
    },
  })
  const data = await res.json();
  return data
}
