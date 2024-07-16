import React, { useState } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import Login from "./components/Login";
import UserList from "./components/UserList";
import UserForm from "./components/UserForm";
import { logout } from "./services/api";

const queryClient = new QueryClient();

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(!!localStorage.getItem("token"));

  const handleLogin = () => {
    setIsLoggedIn(true);
  };

  const handleLogout = () => {
    logout();
    setIsLoggedIn(false);
  };

  return (
    <QueryClientProvider client={queryClient}>
      <div className="App">
        {isLoggedIn ? (
          <>
            <button onClick={handleLogout}>Logout</button>
            <h1>User Management</h1>
            <UserForm
              onUserAdded={() => queryClient.invalidateQueries("users")}
            />
            <UserList />
          </>
        ) : (
          <Login onLogin={handleLogin} />
        )}
      </div>
    </QueryClientProvider>
  );
}

export default App;
