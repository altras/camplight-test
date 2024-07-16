import React from "react";

const ErrorMessage = ({ message }) => (
  <div
    style={{
      color: "red",
      padding: "10px",
      border: "1px solid red",
      borderRadius: "5px",
      marginBottom: "10px",
    }}
  >
    Error: {message}
  </div>
);

export default ErrorMessage;
