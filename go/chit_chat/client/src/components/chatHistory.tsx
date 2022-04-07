import React from "react";
import { Map } from "../App";
import { Message } from "./message";

const styles: Map = {
  div: {
    margin: 0,
    padding: "20px",
    overflow: "auto",
    height: "calc(100vh - 180px)",
  },
};

interface IProps {
  chatHistory: {
    sender?: string;
    data: string;
  }[];
}

export const ChatHistory: React.FC<IProps> = ({ chatHistory }) => {
  const messages = chatHistory.map((message, index) => {
    return <Message key={index} message={message.data} />;
  });

  return (
    <>
      <div style={styles.div}>{messages}</div>
    </>
  );
};
