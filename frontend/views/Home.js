/* @jsx jsx */

import jsx from "../../framework/vDom/jsx"

function createWebSocketConn(e) {
    console.log("Creating webSocket conn")
    e.preventDefault();

}


export function HomeView() {
    return {
        template: (
            <form onSubmit={createWebSocketConn}>
                <label for="name">Enter your username: </label>
                <input type="text" id="name"></input>
                <button>Start game</button>
            </form>
        )
    }
}