"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { z } from "zod";
import { registerUser } from "../../services/api/user";

const registerSchema = z
  .object({
    username: z
      .string()
      .min(3, "Username deve ter no mínimo 3 caracteres")
      .max(32, "Máximo 32 caracteres"),
    email: z.string().email("Email inválido"),
    password: z.string().min(6, "Senha deve ter pelo menos 6 caracteres"),
    confirmPassword: z.string(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    path: ["confirmPassword"],
    message: "As senhas não coincidem",
  });

export type RegisterSchema = z.infer<typeof registerSchema>;

export default function RegisterPage() {
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterSchema>({
    resolver: zodResolver(registerSchema),
  });

  const onSubmit = async (data: RegisterSchema) => {
    const response = await registerUser(data);

    if ("error" in response) return console.error(response.error);

    navigate("/login");
  };

  return (
    <main className="flex h-dvh w-full items-center justify-center bg-black text-white">
      <div className="w-full max-w-md rounded-xl bg-zinc-900 p-8 shadow">
        <h1 className="mb-6 text-center text-2xl font-bold">Criar Conta</h1>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div>
            <label className="mb-1 block text-sm">Username</label>
            <input
              defaultValue={"guivialle"}
              type="text"
              {...register("username")}
              className="w-full rounded border border-zinc-700 bg-zinc-800 px-3 py-2"
            />
            {errors.username && (
              <p className="mt-1 text-sm text-red-500">
                {errors.username.message}
              </p>
            )}
          </div>

          <div>
            <label className="mb-1 block text-sm">Email</label>
            <input
              defaultValue={"guivialle@gmail.com"}
              type="email"
              {...register("email")}
              className="w-full rounded border border-zinc-700 bg-zinc-800 px-3 py-2"
            />
            {errors.email && (
              <p className="mt-1 text-sm text-red-500">
                {errors.email.message}
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

          <div>
            <label className="mb-1 block text-sm">Confirmar Senha</label>
            <input
              defaultValue={"123456"}
              type="password"
              {...register("confirmPassword")}
              className="w-full rounded border border-zinc-700 bg-zinc-800 px-3 py-2"
            />
            {errors.confirmPassword && (
              <p className="mt-1 text-sm text-red-500">
                {errors.confirmPassword.message}
              </p>
            )}
          </div>

          <button
            type="submit"
            className="w-full rounded bg-white py-2 text-black transition hover:bg-gray-200"
          >
            Registrar
          </button>
        </form>
      </div>
    </main>
  );
}
