import type { LoginSchema } from "../../lib/pages/login";
import type { RegisterSchema } from "../../lib/pages/register";
import { api, getStoredAuthSession } from "./client";

export async function authenticateUser({
  login,
  password,
}: LoginSchema): Promise<AuthSessionResponse | ErrorResponse> {
  try {
    const { data } = await api.post<AuthSessionResponse>("/auth/login", {
      login,
      password,
    });
    return data;
  } catch (error: unknown) {
    console.error(error);
    return {
      // @ts-expect-error no types for this error
      error: error.data.message || "Err at trying to authenticate user",
    };
  }
}

export async function registerUser({
  email,
  password,
  username,
}: Omit<RegisterSchema, "confirmPassword">): Promise<
  BasicResponse | ErrorResponse
> {
  try {
    const { data } = await api.post<BasicResponse>("/auth/register", {
      email,
      password,
      username,
    });
    return data;
  } catch (error: unknown) {
    console.error(error);
    return {
      // @ts-expect-error no types for this error
      error: error.data.message || "Err at trying to register user",
    };
  }
}

export async function getUserProfileData(): Promise<
  UserProfileDataResponse | ErrorResponse
> {
  const result = getStoredAuthSession();

  if ("error" in result) return result;

  const token = result.auth.token;

  try {
    const { data } = await api.get<UserProfileDataResponse>("/me/profile", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return data;
  } catch (error: unknown) {
    console.error(error);
    return {
      // @ts-expect-error no types for this error
      error: error.data.message || "Err at trying to get user profile data",
    };
  }
}
