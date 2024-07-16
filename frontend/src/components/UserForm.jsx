import { useState } from "react";
import { useMutation } from "react-query";
import { createUser } from "../services/api";
import ErrorMessage from "./ErrorMessage";

export default function UserForm({ onUserAdded }) {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [error, setError] = useState(null);

  const mutation = useMutation(createUser, {
    onSuccess: (data) => {
      onUserAdded(data);
      setName("");
      setEmail("");
      setPhone("");
      setError(null);
    },
    onError: (error) => {
      setError(error.message);
    },
  });

  const handleSubmit = (e) => {
    e.preventDefault();
    setError(null);
    mutation.mutate({ name, email, phone });
  };

  return (
    <form onSubmit={handleSubmit}>
      {error && <ErrorMessage message={error} />}
      <input
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="Name"
        required
      />
      <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Email"
        required
      />
      <input
        type="tel"
        value={phone}
        onChange={(e) => setPhone(e.target.value)}
        placeholder="Phone"
        required
      />
      <button type="submit" disabled={mutation.isLoading}>
        {mutation.isLoading ? "Adding..." : "Add User"}
      </button>
    </form>
  );
}
