/* @jsx jsx */

import jsx from "../framework/vDom/jsx";
import { store } from "../app";
import { ws } from "../websocket";

function sendMessage(e) {
  e.preventDefault();

  const formInput = e.target.elements.message;

  const msg = {
    type: "TEXT_MESSAGE",
    creator: store.state.currentPlayerName,
    body: formInput.value,
    // body: {
    //     username: store.state.currentPlayerName,
    //     message: formInput.value
    // }
  };

  ws.send(JSON.stringify(msg));

  formInput.value = "";
}

export function ChatRoom() {
  const messages = store.state.messages;
  const currentPlayername = store.state.currentPlayerName;
  console.log("messages:", messages);

  return {
    template: (
      <div id="chatroom">
        <header>
          <h2>Chat room</h2>
          <p>Connected as: {currentPlayername}</p>
        </header>

        <div id="chat">
          <ul>
            {messages.map((message) => {
              return (
                <p>
                  {message.creator}: {message.body}
                </p>
              );
            })}
          </ul>
        </div>
        <form id="send-message" onSubmit={sendMessage}>
          <input
            type="text"
            name="message"
            placeholder="Write a message..."
          ></input>
          <button>Send</button>
        </form>
      </div>
    ),
  };
}
