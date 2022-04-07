import React from "react";

import { Header } from "./components/header";
import { ChatHistory } from "./components/chatHistory";

import { connect, sendMsg } from "./api";
import { ChatInput } from "./components/chatInput";

export interface Map {
  [key: string]: React.CSSProperties | undefined;
}

const styles: Map = {
  app: {
    width: "100%",
    maxWidth: "500px",
    display: "flex",
    flexDirection: "column",
    overflow: "hidden",
    justifyContent: "space-between",
    height: "100vh",
    backgroundColor: "#f7f7f7",
  },
  form: {
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    gap: "40px",
    padding: "20px",
    paddingBottom: "15px",
  },
  button: {
    backgroundColor: "#15223b",
    border: 0,
    color: "white",
    borderRadius: "5px",
    boxShadow: "0 5px 15px -5px rgba(0, 0, 0, 0.2)",
    padding: "10px 20px",
    fontWeight: "bold",
  },
};

interface IHistory {
  sender?: string;
  data: string;
}

const App = () => {
  const [message, setMessage] = React.useState<string>("");
  const [history, setHistory] = React.useState<IHistory[]>([]);
  React.useEffect(() => {
    connect((msg: IHistory) => {
      console.log("New Message");
      setHistory((prev) => [...prev, msg]);
    });
  }, []);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (message.length > 0) {
      sendMsg(message);
      setMessage("");
    }
  };

  return (
    <div style={styles.app}>
      <div>
        <Header />
        <ChatHistory chatHistory={history} />
      </div>
      <form style={styles.form} onSubmit={handleSubmit}>
        <ChatInput
          value={message}
          onChange={(event: any) => setMessage(event.target.value)}
        />
        <button style={styles.button}>Talk</button>
      </form>
    </div>
  );
};

export default App;
