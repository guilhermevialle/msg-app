"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { z } from "zod";
import useAuthSession from "../../hooks/use-auth-session";
import { authenticateUser } from "../../services/api/user";

const loginSchema = z.object({
  login: z
    .string()
    .min(3, "Login é obrigatório")
    .refine((val) => /\S+@\S+\.\S+/.test(val) || /^[a-zA-Z0-9_]+$/.test(val), {
      message: "Deve ser um e-mail ou nome de usuário válido",
    }),
  password: z.string().min(6, "Senha deve ter pelo menos 6 caracteres"),
});

export type LoginSchema = z.infer<typeof loginSchema>;

export default function LoginPage() {
  const navigate = useNavigate();
  const { updateSession } = useAuthSession();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginSchema>({
    resolver: zodResolver(loginSchema),
  });

  const onSubmit = async (data: LoginSchema) => {
    const response = await authenticateUser(data);

    if ("error" in response) return console.error(response.error);

    updateSession(response);
    navigate("/");
  };

  return (
    <main className="flex h-dvh w-full items-center justify-center bg-black text-white">
      <div className="w-full max-w-md rounded-xl bg-zinc-900 p-8 shadow">
        <h1 className="mb-6 text-center text-2xl font-bold">Login</h1>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div>
            <label className="mb-1 block text-sm">
              Login (e-mail ou usuário)
            </label>
            <input
              defaultValue={"guivialle"}
              type="text"
              {...register("login")}
              className="w-full rounded border border-zinc-700 bg-zinc-800 px-3 py-2"
            />
            {errors.login && (
              <p className="mt-1 text-sm text-red-500">
                {errors.login.message}
              </p>
            )}
          </div>

          <div>
            <label className="mb-1 block text-sm">Senha</label>
            <input
              defaultValue={"123456"}
              type="password"
              {...register("password")}
              className="w-full rounded border border-zinc-700 bg-zinc-800 px-3 py-2"
            />
            {errors.password && (
              <p className="mt-1 text-sm text-red-500">
                {errors.password.message}
              </p>
            )}
          </div>

          <button
            type="submit"
            className="w-full rounded bg-white py-2 text-black transition hover:bg-gray-200"
          >
            Entrar
          </button>
        </form>
      </div>
    </main>
  );
}
