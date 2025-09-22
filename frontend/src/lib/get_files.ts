import { API_URL } from '$lib/utils/config'
import type { Res, FileRec } from '$lib/utils/types'

export async function get_files(): Promise<FileRec[]> {
  const ENDPOINT = API_URL + "/files"
  const res = await fetch(ENDPOINT, {
    method: "GET",
    credentials: "include",
  });
  const data = await res.json();
  if (data.error) {
    console.error(data.error);
  }
  const files: FileRec[] = data.message
  return files
}
