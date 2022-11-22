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
        {/* <p>Connected as: {currentPlayername}</p> */}
        {/* {messages.map((message) => {
              return (
                <p>
                  {message.creator}: {message.body}
                </p>
              );
            })} */}

        <header>
          <h2 id="chat-header">Chatroom</h2>
        </header>

        <ul id="chat">
          <li>
            <p class="chat-username">Player 1</p>
            <p>Hello there</p>
          </li>

          <li>
            <p class="chat-username">Player 2</p>
            <p>Whats up</p>
          </li>

          <li>
            <p class="chat-username">Player 3</p>
            <p>Lets play boys</p>
          </li>
          <li>
            <p class="chat-username">Player 3</p>
            <p>Lets play boys</p>
          </li>
          <li>
            <p class="chat-username">Player 3</p>
            <p>Lets play boys</p>
          </li>
          <li>
            <p class="chat-username">Player 3</p>
            <p>Lets play boys</p>
          </li>
          <li>
            <p class="chat-username">Player 3</p>
            <p>Lets play boys</p>
          </li>
          <li>
            <p class="chat-username">Player 3</p>
            <p>Lets play boys</p>
          </li>
          <li>
            <p class="chat-username">Player 3</p>
            <p>Lets play boys</p>
          </li>
          <li>
            <p class="chat-username">Player 3</p>
            <p>Lets play boys</p>
          </li>
        </ul>

        <form id="send-message" onSubmit={sendMessage}>
          <input type="text" name="message" placeholder="Send message"></input>
          <button>&gt;</button>
        </form>
      </div>
    ),
  };
}
