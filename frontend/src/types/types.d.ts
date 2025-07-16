interface AuthSessionResponse {
  user: {
    id: string;
    username: string;
    email: string;
  };
  auth: {
    token: string;
  };
}

interface ErrorResponse {
  error: string;
}

interface BasicResponse {
  message: string;
}

interface UserProfileDataResponse {
  bio: string;
  avatar_url: string;
}

interface PostResponse {
  id: string;
  content: string;
  createdAt: string;
  likes: number;
  comments: unknown[];
  author: {
    id: string;
    username: string;
  };
}
