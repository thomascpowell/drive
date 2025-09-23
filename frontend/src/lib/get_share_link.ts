import { API_URL } from '$lib/utils/config'
import type { Res, FileRec, Share } from '$lib/utils/types'

export async function get_share_link(req: Share): Promise<Res> {
  const ENDPOINT = API_URL + "/share/"
  const res = await fetch(ENDPOINT, {
    method: "POST",
    body: JSON.stringify(req),
    credentials: "include",
  });
  const data = await res.json();
  if (data.error) {
    console.error(data.error);
  }
  // looks right
  console.log(req.FileID, req.TTL);
  return data
}
