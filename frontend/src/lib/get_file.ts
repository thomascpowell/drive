import { API_URL } from '$lib/config'
import type { File } from '$lib/types'

export async function get_file(id: string): Promise<File> {
  const ENDPOINT = API_URL + "/files/" + id
  const res = await fetch(ENDPOINT, {
    method: "GET",
    credentials: "include",
  });
  const data = await res.json();
  if (data.error) {
    console.error(data.error);
  }
  const file: File = data.message
  return file
}
