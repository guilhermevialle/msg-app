import useAuthSession from "../hooks/use-auth-session";

export default function Navbar() {
  const { session, isAuthenticated, logout } = useAuthSession();

  return (
    <nav className="flex h-16 w-full items-center justify-between border-b border-b-neutral-800">
      <a href="/" className="font-bold text-zinc-100">
        My app
      </a>

      {isAuthenticated && session ? (
        <div className="flex items-center gap-3">
          <a
            href="/me"
            className="cursor-pointer font-medium text-zinc-200 underline-offset-4 hover:underline"
          >
            {session.user.username}
          </a>

          <button
            onClick={logout}
            className="flex h-8 items-center justify-center rounded-lg bg-neutral-800 px-4 font-medium text-white transition-colors hover:bg-neutral-700"
          >
            Log out
          </button>
        </div>
      ) : (
        <div className="flex items-center gap-3">
          <a
            href="/login"
            className="flex h-8 items-center justify-center rounded-lg bg-neutral-800 px-4 font-medium text-white transition-colors hover:bg-neutral-700"
          >
            Login
          </a>

          <a
            href="/register"
            className="flex h-8 items-center justify-center rounded-lg bg-blue-500 px-4 font-medium text-white transition-colors hover:bg-blue-600"
          >
            Sign in
          </a>
        </div>
      )}
    </nav>
  );
}
