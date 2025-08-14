export async function checkHealth() {
  try {
    const res = await fetch("http://localhost/api/health", {
      method: "GET",
    });
    const data = await res.json();
    return data;
  } catch (err) {
    return { status: "error" };
  }
}
