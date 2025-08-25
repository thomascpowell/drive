import { API_URL } from '$lib/config'
import type { Res, File } from '$lib/types'

export async function get_files(): Promise<Res> {
  const ENDPOINT = API_URL + "/upload"
  await fetch(ENDPOINT, {
    method: "GET",
    credentials: "include",
  });
}
