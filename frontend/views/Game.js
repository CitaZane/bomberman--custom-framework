/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import { store } from "../app";

// function for getting the user count from backend?

export function GameView() {
    const userQueueCount = store.state.userQueueCount;

    return {
        template: (
            <div>
                <h1>Game room</h1>
                <h3>Users joined: {userQueueCount}</h3>
            </div>
        )
    }
}
