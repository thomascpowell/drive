import { API_URL } from "./config";

export async function upload(file: File) {
  const endpoint = API_URL + "/upload"
  const formData = new FormData();
  formData.append("file", file)

  await fetch(endpoint, {
    method: "POST",
    body: formData,
    credentials: "include",
  });
}
