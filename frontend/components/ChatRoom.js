/* @jsx jsx */

import jsx from "../../framework/vDom/jsx"
import { store } from "../app";
import { ws } from "../websocket";

function sendMessage(e) {
    e.preventDefault();

    const formInput = e.target.elements.message;

    const msg = {
        type: "TEXT_MESSAGE",
        body: formInput.value,
    }

    ws.send(JSON.stringify(msg));

    formInput.value = "";
}

export function ChatRoom() {
    const messages = store.state.messages
    console.log("messages:",messages)
    
    return {
        template: (
            <div id="chatroom">
                <header>
                    <h2>Chat room</h2>
                    <p>Connected as: John Doe</p>
                </header>

                <div id="chat">
                    <p>{messages}</p>
                </div>
                <form id="send-message" onSubmit={sendMessage}>
                    <input type="text" name="message" placeholder="Write a message..."></input>
                    <button>Send</button>
                </form>
            </div>
        )
    }
}