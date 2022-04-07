import React from "react";
import { Map } from "../App";

const styles: Map = {
  me: {
    color: "white",
    float: "right",
    backgroundColor: "#328ec4",
  },
  div: {
    display: "block",
    backgroundColor: "white",
    margin: "10px auto",
    boxShadow: "0 5px 15px -5px rgba(0, 0, 0, 0.2)",
    padding: "10px 20px",
    borderRadius: "5px",
    clear: "both",
  },
};

interface IProps {
  message: any;
}

export const Message: React.FC<IProps> = ({ message }) => {
  const temp = JSON.parse(message);

  return <div style={styles.div}>{temp.body}</div>;
};
