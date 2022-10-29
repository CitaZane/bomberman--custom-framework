/* @jsx jsx */

import jsx from "../../framework/vDom/jsx"
import { store } from "../app";

function createWebSocketConn(e) {
    e.preventDefault();

    const inputElem = e.target.elements["name"];

    defineWebSocket(inputElem.value)

    window.location.href = window.location.origin + "/#/game";

}

function defineWebSocket(name) {
    const ws = new WebSocket(`ws://localhost:8080/ws?username=${name}`);

    ws.onopen = () => {
        console.log("Connection initiated")
    }

    ws.onclose = () => {
        console.log("Connection closed")
    }

    ws.onmessage = (e) => {
        const data = JSON.parse(e.data);
        switch (data["type"]) {
            case "NEW_USER": {
                store.commit("updateUserQueueCount", data.body)
            }
            case "USER_LEFT": {
                store.commit("updateUserQueueCount", data.body)
            }
        }

    }
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