import axios from "axios";

export const api = axios.create({
  baseURL: "http://localhost:8080",
});

export function getStoredAuthSession(): AuthSessionResponse | ErrorResponse {
  const raw = localStorage.getItem("auth-session");
  const authSession = raw && JSON.parse(raw);
  return (
    (authSession as AuthSessionResponse) ?? { error: "user not authenticated" }
  );
}
