/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";
import { store } from "../app";
import { ChatRoom } from "../components/ChatRoom";

export function GameView() {
    let userQueueCount = store.state.userQueueCount;

    return {
        template: (
            <div>
                <h1>Game room</h1>
                <h3>Users joined: {String(userQueueCount)}</h3>
                <ChatRoom />

            </div>
        ),
    };
}
