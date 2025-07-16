"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import useAuthSession from "../hooks/use-auth-session";
import { submitPost } from "../services/api/post";
import { queryClient } from "./query-provider";

const postSchema = z.object({
  content: z
    .string()
    .min(1, "Content is required")
    .max(400, "Max 400 characters"),
});

type PostSchema = z.infer<typeof postSchema>;

export default function NewPostForm() {
  const { isAuthenticated, session } = useAuthSession();
  const {
    register,
    handleSubmit,
    watch,
    reset,
    formState: { errors, isSubmitting },
  } = useForm<PostSchema>({
    resolver: zodResolver(postSchema),
  });

  const content = watch("content", "");

  const onSubmit = async ({ content }: PostSchema) => {
    const response = await submitPost(content);

    if ("error" in response) return console.error(response.error);

    queryClient.invalidateQueries({
      queryKey: ["posts"],
    });

    reset();
  };

  const show = isAuthenticated && session;

  return (
    show && (
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="relative mx-auto mt-12 max-w-xl space-y-2 rounded-lg border border-zinc-700 bg-zinc-900 p-4"
      >
        <textarea
          {...register("content")}
          rows={5}
          maxLength={400}
          className="w-full resize-none rounded-md bg-zinc-800 p-3 text-sm text-white placeholder-zinc-400 outline-none"
          placeholder="What's on your mind?"
        />
        <div className="flex items-center justify-between text-xs text-zinc-400">
          <span>{errors.content?.message}</span>
          <span>{content.length} / 400</span>
        </div>
        <div className="flex justify-end">
          <button
            type="submit"
            disabled={isSubmitting}
            className="rounded-md bg-blue-600 px-4 py-2 text-sm font-semibold text-white hover:bg-blue-700 disabled:opacity-50"
          >
            Post
          </button>
        </div>
      </form>
    )
  );
}
