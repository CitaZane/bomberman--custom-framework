/* @jsx jsx */

import jsx from "../../framework/vDom/jsx"

function createWebSocketConn(e) {
    console.log("Creating webSocket conn")

    e.preventDefault();

    const inputElem = e.target.elements["name"];
    const ws = new WebSocket(`ws://localhost:8080/ws?username=${inputElem.value}`);
    ws.onopen = () => {
        console.log("Connection initiated")
    }

    ws.onclose = () => {
        console.log("Connection closed")
    }


    window.location.href = window.location.origin + "/#/game";
}


export function HomeView() {
    return {
        template: (
            <form onSubmit={createWebSocketConn} id="create-username">
                <label for="name">Enter your username: </label>
                <input type="text" id="name"></input>
                <button>Start game</button>
            </form>
        )
    }
}