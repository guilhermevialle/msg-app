import { useQuery } from "@tanstack/react-query";
import { getPosts } from "../services/api/post";

export default function Feed() {
  const { data } = useQuery({
    queryFn: getPosts,
    queryKey: ["posts"],
  });

  return (
    <div className="flex gap-3 text-zinc-200">
      {data && <pre>{JSON.stringify(data, null, 2)}</pre>}
    </div>
  );
}
