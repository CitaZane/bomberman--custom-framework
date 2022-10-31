import { store } from "./app";

export let ws;

export function defineWebSocket(name) {
    const ws1 = new WebSocket(`ws://localhost:8080/ws?username=${name}`);

    ws1.onopen = () => {
        console.log("Connection initiated");
    };

    ws1.onclose = () => {
        console.log("Connection closed");
    };

    ws1.onmessage = (e) => {
        const data = JSON.parse(e.data);
        switch (data["type"]) {
            case "NEW_USER":
            case "USER_LEFT":
                store.commit("updateUserQueueCount", data.body);

            case "INIT_ROOM":
                store.commit("updateUserQueueCount", data.body);
                console.log("Navigating to game room")
                window.location.href = window.location.origin + "/#/game";
        }
    };

    ws = ws1
}
