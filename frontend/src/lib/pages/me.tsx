import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import Navbar from "../../components/navbar";
import useAuthSession from "../../hooks/use-auth-session";
import { getUserProfileData } from "../../services/api/user";

export default function MePage() {
  const { session, isAuthenticated } = useAuthSession();
  const navigate = useNavigate();
  const { data } = useQuery<UserProfileDataResponse | ErrorResponse>({
    queryFn: getUserProfileData,
    queryKey: ["user-profile-data"],
  });

  if (!session && !isAuthenticated) navigate("/login");

  return (
    <main className="h-dvh w-full bg-black">
      <div className="mx-auto max-w-xl">
        <Navbar />

        <pre className="text-zinc-200">{JSON.stringify(data, null, 2)}</pre>
      </div>
    </main>
  );
}
