export function has_cookie(name: string = "jwt"): boolean {
  return document.cookie
    .split("; ")
    .some((row) => row.startsWith(`${name}=`));
}
