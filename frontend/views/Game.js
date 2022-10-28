/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import { store } from "../app";

function updateCount() {
    // console.log("CLICKED")
    // console.log(++store.state.count)
    // console.log(store.state.userQueueCount)
}
export function GameView() {
    const userQueueCount = store.state.userQueueCount;
    console.log(userQueueCount)
    return {
        template: (
            <div>
                <h1 onClick={updateCount}>Game room</h1>
                <h3>Users joined: {userQueueCount}</h3>
            </div>
        )
    }
}
