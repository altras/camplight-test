import { useState } from 'react'
import { useQuery } from 'react-query'
import { fetchUsers, deleteUser } from '../services/api'

export default function UserList({ onUserDeleted }) {
  const [page, setPage] = useState(1)
  const limit = 10

  const { data, isLoading, error } = useQuery(['users', page], () => fetchUsers(page, limit), {
    keepPreviousData: true,
  })

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>Error: {error.message}</div>

  return (
    <div>
      <ul>
        {data.users.map(user => (
          <li key={user.id}>
            {user.name} - {user.email} - {user.phone}
            <button onClick={() => deleteUser(user.id).then(() => onUserDeleted(user.id))}>Delete</button>
          </li>
        ))}
      </ul>
      <div>
        <button onClick={() => setPage(old => Math.max(old - 1, 1))} disabled={page === 1}>
          Previous Page
        </button>
        <span>Page {page}</span>
        <button onClick={() => setPage(old => old + 1)} disabled={!data.hasMore}>
          Next Page
        </button>
      </div>
    </div>
  )
}