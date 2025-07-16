import { api, getStoredAuthSession } from "./client";

export async function getPosts(): Promise<PostResponse[] | ErrorResponse> {
  try {
    const { data } = await api.get<PostResponse[]>("/posts");
    return data;
  } catch (error: unknown) {
    console.error(error);
    return {
      // @ts-expect-error no types for this error
      error: error.data.message || "Err at trying to get posts",
    };
  }
}

export async function submitPost(
  content: string,
): Promise<BasicResponse | ErrorResponse> {
  const result = getStoredAuthSession();

  if ("error" in result) return result;

  const token = result.auth.token;

  try {
    const { data } = await api.post<BasicResponse>(
      "/post/new",
      { content },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );
    return data;
  } catch (error: unknown) {
    console.error(error);
    return {
      // @ts-expect-error no types for this error
      error: error.data.message || "Err at trying to submit post",
    };
  }
}
