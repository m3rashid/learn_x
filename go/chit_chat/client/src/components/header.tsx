import React from "react";
import { Map } from "../App";

const styles: Map = {
  div: {
    backgroundColor: "#15223b",
    width: "100%",
    margin: 0,
    padding: "20px",
    color: "white",
  },
  h2: {
    margin: 0,
    padding: 0,
  },
};

interface IProps {}

export const Header: React.FC<IProps> = () => {
  return (
    <div style={styles.div}>
      <h2 style={styles.h2}>Chit Chat</h2>
    </div>
  );
};
