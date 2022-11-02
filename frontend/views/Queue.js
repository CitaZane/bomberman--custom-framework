/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";
import { store } from "../app";
import { ChatRoom } from "../components/ChatRoom";
import { ws } from "../websocket";
function startGame() {
    ws.send(JSON.stringify({ type: "START_GAME" }));
}


export function QueueView() {
    let userQueueCount = store.state.userQueueCount;

    return {
        template: (
            <div>
                <h1>Queue</h1>
                <h3>Users joined: {String(userQueueCount)}</h3>
                <button onClick={startGame}>Start Game</button>
                <ChatRoom />

            </div>
        ),
    };
}
