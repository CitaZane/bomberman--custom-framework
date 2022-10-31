/* @jsx jsx */

import jsx from "../../framework/vDom/jsx"
import { ws } from "../websocket";

function sendMessage(e) {
    e.preventDefault();

    const formInput = e.target.elements.message;

    const msg = {
        type: "Message",
        body: formInput.value,
    }

    ws.send(JSON.stringify(msg));

    formInput.value = "";
}

export function ChatRoom() {
    return {
        template: (
            <div id="chatroom">
                <header>
                    <h2>Chat room</h2>
                    <p>Connected as: John Doe</p>
                </header>

                <div id="chat"></div>
                <form id="send-message" onSubmit={sendMessage}>
                    <input type="text" name="message" placeholder="Write a message..."></input>
                    <button>Send</button>
                </form>
            </div>
        )
    }
}