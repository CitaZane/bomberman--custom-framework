/* @jsx jsx */

import jsx from "../../framework/vDom/jsx"

export function ChatRoom() {
    return {
        template: (
            <div id="chatroom">
                <header>
                    <h2>Chat room</h2>
                    <p>Connected as: John Doe</p>
                </header>

                <div id="chat"></div>
                <form id="send-message">
                    <input type="text" placeholder="Write a message..."></input>
                    <button>Send</button>
                </form>
            </div>
        )
    }
}