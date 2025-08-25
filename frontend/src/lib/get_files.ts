import { API_URL } from '$lib/config'
import type { Res, File } from '$lib/types'

export async function get_files(): Promise<File[]> {
  const ENDPOINT = API_URL + "/files"
  const res = await fetch(ENDPOINT, {
    method: "GET",
    credentials: "include",
  });
  const data = await res.json();
  if (data.error) {
    console.error(data.error);
  }
  const files: File[] = data.message
  return files
}
