import axios from 'axios'

const API_URL = 'http://localhost:8080/api'

const api = axios.create({
  baseURL: API_URL,
})

api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

export const login = (email, password) =>
  api.post('/login', { email, password }).then((res) => {
    localStorage.setItem('token', res.data.token)
    return res.data
  })

export const logout = () => {
  localStorage.removeItem('token')
}

export const fetchUsers = (page, limit) => 
  api.get(`${API_URL}/users?page=${page}&limit=${limit}`)
    .then(res => ({
      users: res.data,
      hasMore: res.data.length === limit
    }))

export const createUser = (user) => api.post(`${API_URL}/users`, user).then(res => res.data)
export const deleteUser = (id) => api.delete(`${API_URL}/users/${id}`)

export const searchUsers = (query, page, limit) => 
  api.get(`${API_URL}/users/search?q=${query}&page=${page}&limit=${limit}`)
    .then(res => ({
      users: res.data,
      hasMore: res.data.length === limit
    }))
