import Feed from "./components/feed";
import Navbar from "./components/navbar";
import NewPostForm from "./components/new-post-form";

export default function App() {
  return (
    <main className="min-h-[100dvh] w-full bg-black">
      <div className="mx-auto max-w-xl">
        <Navbar />
        <NewPostForm />
        <Feed />
      </div>
    </main>
  );
}
