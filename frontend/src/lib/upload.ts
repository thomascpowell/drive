import { API_URL } from "./config";
import type { Res } from '$lib/types'

export async function upload(file: File): Promise<Res> {
  const ENDPOINT = API_URL + "/upload"
  const formData = new FormData();
  formData.append("file", file)

  const res = await fetch(ENDPOINT, {
    method: "POST",
    body: formData,
    credentials: "include",
  });

  const data = await res.json();
  return data
}
