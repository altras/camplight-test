import { useState } from "react";
import { useQuery } from "react-query";
import { fetchUsers, deleteUser, searchUsers } from "../services/api";
import ErrorMessage from "./ErrorMessage";

export default function UserList({ onUserDeleted }) {
  const [page, setPage] = useState(1);
  const [searchQuery, setSearchQuery] = useState("");
  const limit = 10;

  const { data, isLoading, error } = useQuery(
    ["users", page, searchQuery],
    () =>
      searchQuery
        ? searchUsers(searchQuery, page, limit)
        : fetchUsers(page, limit),
    { keepPreviousData: true }
  );

  // Handle loading state
  if (isLoading) return <div>Loading...</div>;

  // Handle error state
  if (error) return <ErrorMessage message={error.message} />;

  return (
    <div>
      <input
        type="text"
        placeholder="Search users..."
        value={searchQuery}
        onChange={(e) => {
          setSearchQuery(e.target.value);
          setPage(1);
        }}
      />
      <ul>
        {data.users.map((user) => (
          <li key={user.id}>
            {user.name} - {user.email} - {user.phone}
            <button
              onClick={() =>
                deleteUser(user.id)
                  .then(() => onUserDeleted(user.id))
                  .catch((err) =>
                    alert(`Failed to delete user: ${err.message}`)
                  )
              }
            >
              Delete
            </button>
          </li>
        ))}
      </ul>
      <div>
        <button
          onClick={() => setPage((old) => Math.max(old - 1, 1))}
          disabled={page === 1}
        >
          Previous Page
        </button>
        <span>Page {page}</span>
        <button
          onClick={() => setPage((old) => old + 1)}
          disabled={!data.hasMore}
        >
          Next Page
        </button>
      </div>
    </div>
  );
}
