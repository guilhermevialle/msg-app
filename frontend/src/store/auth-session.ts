import { atomWithStorage } from "jotai/utils";

export const authSessionAtom = atomWithStorage<AuthSessionResponse | null>(
  "auth-session",
  null,
);
