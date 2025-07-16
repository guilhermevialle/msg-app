import { useAtom } from "jotai/react";
import { useEffect, useState } from "react";
import { authSessionAtom } from "../store/auth-session";

export default function useAuthSession() {
  const [session, setSession] = useAtom(authSessionAtom);
  const [isAuthenticated, setAuthenticated] = useState(!!session);

  const updateSession = (session: AuthSessionResponse | null) => {
    setSession(session);
    setAuthenticated(!!session);
  };

  const logout = () => {
    updateSession(null);
  };

  useEffect(() => {
    setAuthenticated(!!session);
  }, [session]);

  return { session, isAuthenticated, logout, updateSession };
}
