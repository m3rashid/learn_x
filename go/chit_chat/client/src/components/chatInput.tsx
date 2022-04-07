import React from "react";
import { Map } from "../App";

const styles: Map = {
  input: {
    padding: "10px",
    margin: 0,
    fontSize: "16px",
    borderRadius: "5px",
    border: "1px solid rgba(0, 0, 0, 0.1)",
    boxShadow: "0 5px 15px -5px rgba(0, 0, 0, 0.1)",
    outline: "none",
    width: "100%",
  },
  div: {
    display: "block",
    width: "100%",
  },
};

interface IProps {
  value: string;
  onChange: React.ChangeEventHandler<HTMLInputElement> | undefined;
}

export const ChatInput: React.FC<IProps> = ({ value, onChange }) => {
  return (
    <div style={styles.div}>
      <input value={value} onChange={onChange} style={styles.input} />
    </div>
  );
};
